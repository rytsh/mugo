package args

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/rytsh/liz/loader/file"
	"github.com/rytsh/liz/utils/mapx"
	"github.com/spf13/cobra"
	"github.com/worldline-go/logz"

	"github.com/rytsh/mugo/internal/config"
	"github.com/rytsh/mugo/pkg/template"
)

type AppInfo struct {
	Version     string
	BuildCommit string
	BuildDate   string
}

var ErrShutdown = errors.New("shutting down signal received")

var rootCmd = &cobra.Command{
	Use:           "mugo template.tpl",
	Short:         "go template executor",
	Long:          "execute go template and export it to stdout or file",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if config.App.Silience {
			err := logz.SetLogLevel("panic")
			if err != nil {
				log.Error().Err(err).Msg("failed to set log level")
			}
		}

		log.Info().Msgf("MUGO [%s]", cmd.Version)
	},

	Example: "mugo -d @data.yaml template.tpl" + "\n" +
		`mugo -d '{"Name": "mugo"}' -o output.txt template.tpl`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.App.List {
			log.Info().Msg("print function list")
			template.New().ListFunctions()

			return nil
		}

		if len(args) == 0 {
			return fmt.Errorf("no input file")
		}

		if config.App.Delims != "" {
			fields := strings.Fields(strings.ReplaceAll(config.App.Delims, ",", " "))
			if len(fields) != 2 {
				return fmt.Errorf("invalid delims: %s", config.App.Delims)
			}

			config.Checked.Delims = fields
		}

		workDir, err := filepath.Abs(filepath.Clean(args[0]))
		if err != nil {
			return fmt.Errorf("failed to get absolute path: %w", err)
		}

		config.Checked.WorkDir = filepath.Dir(workDir)

		return mugo(cmd.Context(), args)
	},
}

// Execute is the entry point for the application.
func Execute(ctx context.Context, appInfo AppInfo) error {
	rootCmd.Version = appInfo.Version
	rootCmd.Long += "\nversion: " + appInfo.Version + " commit: " + appInfo.BuildCommit + " buildDate:" + appInfo.BuildDate

	return rootCmd.ExecuteContext(ctx)
}

func init() {
	rootCmd.Flags().StringVar(&config.App.Delims, "delims", config.App.Delims, "comma or space separated list of delimiters to alternate the default \"{{ }}\"")
	rootCmd.Flags().StringArrayVarP(&config.App.Data, "data", "d", config.App.Data, "input data as json/yaml or file path with @ prefix")
	rootCmd.Flags().StringVarP(&config.App.Output, "output", "o", config.App.Output, "output file, default is stdout")
	rootCmd.Flags().BoolVarP(&config.App.Silience, "silience", "s", config.App.Silience, "silience log")
	rootCmd.Flags().BoolVarP(&config.App.List, "list", "l", config.App.List, "function List")
}

func mugo(ctx context.Context, inputs []string) (err error) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	ctx, ctxCancel := context.WithCancel(ctx)
	defer ctxCancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-sig:
			log.Warn().Msg("received shutdown signal")
			ctxCancel()

			if err != nil {
				err = ErrShutdown
			}
		case <-ctx.Done():
		}
	}()

	tpl := template.New()
	tpl.SetDelims(config.Checked.Delims[0], config.Checked.Delims[1])

	output := os.Stdout
	if config.App.Output != "" {
		output, err = os.OpenFile(config.App.Output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		defer output.Close()
	}

	// read input data
	inputData := make(map[string]interface{})

	fileAPI := file.New()
	for _, data := range config.App.Data {
		var currentData map[string]interface{}
		var err error

		if data[0] == '@' {
			currentData, err = fileAPI.LoadMap(data[1:])
		} else {
			currentData, err = fileAPI.ContentMap(data, fileAPI.Codec["YAML"])
		}

		if err != nil {
			return fmt.Errorf("failed to load input data: %w", err)
		}

		mapx.Merge(currentData, inputData)
	}

	log.Info().Msgf("output: %s", output.Name())
	log.Info().Msgf("execute template: %s", inputs[0])

	// open file
	v, err := os.ReadFile(inputs[0])
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if err := tpl.ExecuteContent(output, inputData, v); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	log.Info().Msg("render completed")

	return
}
