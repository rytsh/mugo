package args

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/rytsh/liz/loader/file"
	"github.com/rytsh/liz/utils/mapx"
	"github.com/rytsh/liz/utils/templatex"
	"github.com/rytsh/liz/utils/templatex/functions"
	"github.com/spf13/cobra"
	"github.com/worldline-go/logz"

	"github.com/rytsh/mugo/internal/banner"
	"github.com/rytsh/mugo/internal/config"
	"github.com/rytsh/mugo/internal/render"
	"github.com/rytsh/mugo/internal/request"
	"github.com/rytsh/mugo/internal/shutdown"
)

type AppInfo struct {
	Version     string
	BuildCommit string
	BuildDate   string
}

var ErrShutdown = errors.New("shutting down signal received")

var rootCmd = &cobra.Command{
	Use:           "mugo <template>",
	Short:         "go template executor",
	Long:          banner.Logo + "execute go template and export it to stdout or file",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if config.App.Silience {
			err := logz.SetLogLevel("fatal")
			if err != nil {
				log.Error().Err(err).Msg("failed to set log level")
			}
		} else {
			err := logz.SetLogLevel(config.App.LogLevel)
			if err != nil {
				log.Error().Err(err).Msg("failed to set log level")
			}
		}

		log.Info().Msgf("MUGO [%s]", cmd.Version)
	},

	Example: "mugo -d @data.yaml template.tpl" + "\n" +
		`mugo -d '{"Name": "mugo"}' -o output.txt template.tpl` + "\n" +
		`mugo -d '{"Name": "mugo"}' -o output.txt - < template.tpl`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		if config.App.List {
			log.Info().Msg("print function list")
			tpl := templatex.New(functions.WithAddFuncsTpl(render.FuncMap(config.App.Trust)))
			tpl.ListFunctions()

			return nil
		}

		if len(args) == 0 {
			return fmt.Errorf("missing template file")
		}

		if config.App.Delims != "" {
			fields := strings.Fields(strings.ReplaceAll(config.App.Delims, ",", " "))
			if len(fields) != 2 {
				return fmt.Errorf("invalid delims: %s", config.App.Delims)
			}

			config.Checked.Delims = fields
		}

		var (
			inputReader io.Reader = cmd.InOrStdin()
			isFile                = false
			inputData   []byte
		)

		info := os.Stdin.Name()

		// the argument received looks like a file, we try to open it
		if len(args) > 0 && args[0] != "-" {
			// if p is an http url, we try to download it
			if _, err := url.ParseRequestURI(args[0]); err == nil {
				httpReq := request.New()
				body, err := httpReq.Get(ctx, args[0])
				if err != nil {
					return err
				}

				inputData = body
			} else {
				isFile = true
				body, err := os.ReadFile(args[0])
				if err != nil {
					return err
				}

				inputData = body
			}

			info = args[0]
		} else {
			// read from stdin
			body, err := io.ReadAll(inputReader)
			if err != nil {
				return err
			}

			inputData = body
		}

		if isFile {
			workDir, err := filepath.Abs(filepath.Clean(args[0]))
			if err != nil {
				return fmt.Errorf("failed to get absolute path: %w", err)
			}

			config.Checked.WorkDir = filepath.Dir(workDir)
		}

		return mugo(ctx, inputData, info)
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
	rootCmd.Flags().StringArrayVarP(&config.App.Data, "data", "d", config.App.Data, "input data as json/yaml or file path with @ prefix could be '.yaml','.yml','.json','.toml' extension")
	rootCmd.Flags().StringVarP(&config.App.DataRaw, "data-raw", "r", config.App.DataRaw, "input data as raw or file path with @ prefix could be file with any extension")
	rootCmd.Flags().StringArrayVarP(&config.App.Parse, "parse", "p", config.App.Parse, "parse file pattern for define templates 'testdata/**/*.tpl'")
	rootCmd.Flags().StringVarP(&config.App.Output, "output", "o", config.App.Output, "output file, default is stdout")
	rootCmd.Flags().BoolVarP(&config.App.Silience, "silience", "s", config.App.Silience, "silience log")
	rootCmd.Flags().BoolVarP(&config.App.List, "list", "l", config.App.List, "function list")
	rootCmd.Flags().BoolVar(&config.App.DisableAt, "no-at", config.App.DisableAt, "disable @ prefix for file path")
	rootCmd.Flags().BoolVarP(&config.App.Trust, "trust", "t", config.App.Trust, "trust to execute dangerous functions")
	rootCmd.Flags().BoolVarP(&config.App.SkipVerify, "insecure", "k", config.App.SkipVerify, "skip verify ssl certificate")
	rootCmd.Flags().BoolVar(&config.App.DisableRetry, "no-retry", config.App.DisableRetry, "disable retry")
	rootCmd.Flags().StringVar(&config.App.LogLevel, "log-level", config.App.LogLevel, "log level (debug, info, warn, error, fatal, panic), default is info")
}

func mugo(ctx context.Context, input []byte, info string) (err error) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	ctx, ctxCancel := context.WithCancel(ctx)
	defer ctxCancel()

	wg.Add(1)
	go shutdown.Global.WatchCtx(ctx, wg)

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

	httpReq := request.New()

	tpl := templatex.New(functions.WithAddFuncsTpl(render.FuncMap(config.App.Trust))).SetDelims(config.Checked.Delims[0], config.Checked.Delims[1])
	for _, p := range config.App.Parse {
		// if p is an http url, we try to download it
		if _, err := url.ParseRequestURI(p); err == nil {
			body, err := httpReq.Get(ctx, p)
			if err != nil {
				return err
			}

			if err := tpl.Parse(string(body)); err != nil {
				return fmt.Errorf("failed to parse template: %w", err)
			}

			continue
		}

		if err := tpl.ParseGlob(p); err != nil {
			return fmt.Errorf("failed to parse glob: %w", err)
		}
	}

	output := os.Stdout
	if config.App.Output != "" {
		output, err = os.OpenFile(config.App.Output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		defer output.Close()
	}

	// read input data
	var inputData interface{}

	fileAPI := file.New()
	if config.App.DataRaw != "" {
		if !config.App.DisableAt && config.App.DataRaw[0] == '@' {
			if d, err := fileAPI.LoadRaw(config.App.DataRaw[1:]); err != nil {
				return fmt.Errorf("failed to load input data: %w", err)
			} else {
				inputData = string(d)
			}
		} else {
			inputData = config.App.DataRaw
		}
	} else {
		storeData := make(map[string]interface{})
		for _, data := range config.App.Data {
			var currentData map[string]interface{}
			var err error

			if !config.App.DisableAt && data[0] == '@' {
				err = fileAPI.Load(data[1:], &currentData)
			} else {
				err = fileAPI.LoadContent([]byte(data), &currentData, fileAPI.Codec["YAML"])
			}

			if err != nil {
				return fmt.Errorf("failed to load input data: %w", err)
			}

			mapx.Merge(currentData, storeData)
		}

		inputData = storeData
	}

	log.Info().Msgf("output: %s", output.Name())
	log.Info().Msgf("execute template: %s", info)

	if err := tpl.Parse(string(input)); err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	if err := tpl.Execute(
		templatex.WithIO(output),
		templatex.WithData(inputData),
		templatex.WithParsed(true),
	); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	log.Info().Msg("render completed")

	return
}

func ReadAll(r io.Reader) (string, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	return string(content), nil
}
