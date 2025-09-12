package args

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/rakunlabs/into"
	"github.com/rakunlabs/logi"
	"github.com/rytsh/liz/file"
	"github.com/rytsh/liz/mapx"
	"github.com/rytsh/liz/shutdown"
	"github.com/spf13/cobra"

	_ "github.com/rytsh/mugo/fstore/registry"

	"github.com/rytsh/mugo/fstore"
	"github.com/rytsh/mugo/fstore/registry/random"
	"github.com/rytsh/mugo/internal/banner"
	"github.com/rytsh/mugo/internal/config"
	"github.com/rytsh/mugo/internal/request"
	"github.com/rytsh/mugo/templatex"
)

type AppInfo struct {
	Version     string
	BuildCommit string
	BuildDate   string
}

var rootCmd = &cobra.Command{
	Use:           "mugo <template>",
	Short:         "go template executor",
	Long:          banner.Logo + "execute go template and export it to stdout or file",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if config.App.Silience {
			if err := logi.SetLogLevel("ERROR"); err != nil {
				slog.Error("failed to set log level", slog.String("error", err.Error()))
			}
		} else {
			if err := logi.SetLogLevel(config.App.LogLevel); err != nil {
				slog.Error("failed to set log level", slog.String("error", err.Error()))
			}
		}

		slog.Info("MUGO [" + cmd.Version + "]")
	},

	Example: "mugo -d @data.yaml template.tpl" + "\n" +
		`mugo -d '{"Name": "mugo"}' -o output.txt template.tpl` + "\n" +
		`mugo -d '{"Name": "mugo"}' -o output.txt - < template.tpl` + "\n" +
		`mugo -d '{"Name": "mugo"}' - <<< "{{.Name}}"` + "\n" +
		`mugo -d '{"Name": "mugo"}' -t @template.tpl` + "\n" +
		`mugo -t '{{.Name}}' data.yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		if config.App.List {
			slog.Info("print function list")

			tpl := templatex.New(templatex.WithAddFuncMapWithOpts(FStore))

			for _, v := range tpl.ListFuncs() {
				fmt.Println(v.Description)
			}

			return nil
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
				httpReq, err := request.New()
				if err != nil {
					return err
				}

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
			if !config.App.NoStdin && !(len(config.App.Data) > 0 && config.App.Template != "") {
				slog.Info("read input from stdin")
				// read from stdin
				body, err := io.ReadAll(inputReader)
				if err != nil {
					return err
				}

				inputData = body
			}
		}

		if config.Checked.WorkDir == "" {
			if isFile {
				workDir, err := filepath.Abs(filepath.Clean(args[0]))
				if err != nil {
					return fmt.Errorf("failed to get absolute path: %w", err)
				}

				config.Checked.WorkDir = filepath.Dir(workDir)
			} else {
				config.Checked.WorkDir = "."
			}
		}

		return mugo(ctx, inputData, info)
	},
}

// Execute is the entry point for the application.
func Execute(ctx context.Context, appInfo AppInfo) error {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel != "" {
		config.App.LogLevel = logLevel
	}

	rootCmd.Version = appInfo.Version
	rootCmd.Long += "\nversion: " + appInfo.Version + " commit: " + appInfo.BuildCommit + " buildDate:" + appInfo.BuildDate

	return rootCmd.ExecuteContext(ctx)
}

func init() {
	rootCmd.Flags().BoolVar(&config.App.HtmlTemplate, "html", config.App.HtmlTemplate, "use html/template instead")
	rootCmd.Flags().StringVar(&config.App.Delims, "delims", config.App.Delims, "comma or space separated list of delimiters to alternate the default \"{{ }}\"")
	rootCmd.Flags().StringArrayVarP(&config.App.Data, "data", "d", config.App.Data, "input data as json/yaml or file path with @ prefix could be '.yaml','.yml','.json','.toml' extension")
	rootCmd.Flags().BoolVarP(&config.App.DataRaw, "data-raw", "r", config.App.DataRaw, "set input data as raw")
	rootCmd.Flags().BoolVarP(&config.App.DataRawByte, "data-raw-byte", "b", config.App.DataRawByte, "raw data is byte")
	rootCmd.Flags().StringVarP(&config.App.Template, "template", "t", config.App.Template, "input template as raw or file path with @ prefix could be file with any extension")
	rootCmd.Flags().StringArrayVarP(&config.App.Parse, "parse", "p", config.App.Parse, "parse file pattern for define templates 'testdata/**/*.tpl'")
	rootCmd.Flags().StringVarP(&config.App.Output, "output", "o", config.App.Output, "output file, default is stdout")
	rootCmd.Flags().BoolVarP(&config.App.Silience, "silience", "s", config.App.Silience, "silience log")
	rootCmd.Flags().BoolVarP(&config.App.List, "list", "l", config.App.List, "function list")
	rootCmd.Flags().StringArrayVar(&config.App.SpecificGroups, "enable-group", config.App.SpecificGroups, "specific function groups for run template")
	rootCmd.Flags().StringArrayVar(&config.App.SpecificFuncs, "enable-func", config.App.SpecificFuncs, "specific functions for run template")
	rootCmd.Flags().StringArrayVar(&config.App.DisabledGroups, "disable-group", config.App.DisabledGroups, "disabled groups for run template")
	rootCmd.Flags().StringArrayVar(&config.App.DisabledFuncs, "disable-func", config.App.DisabledFuncs, "disabled functions for run template")
	rootCmd.Flags().BoolVarP(&config.App.NoStdin, "no-stdin", "n", config.App.NoStdin, "disable stdin input")
	rootCmd.Flags().BoolVar(&config.App.DisableAt, "no-at", config.App.DisableAt, "disable @ prefix for file path")
	rootCmd.Flags().BoolVar(&config.App.Trust, "trust", config.App.Trust, "trust to execute dangerous functions")
	rootCmd.Flags().BoolVarP(&config.App.SkipVerify, "insecure", "k", config.App.SkipVerify, "skip verify ssl certificate")
	rootCmd.Flags().BoolVar(&config.App.DisableRetry, "no-retry", config.App.DisableRetry, "disable retry on request")
	rootCmd.Flags().StringVar(&config.App.LogLevel, "log-level", config.App.LogLevel, "log level (debug, info, warn, error), default is info")
	rootCmd.Flags().StringVarP(&config.Checked.WorkDir, "work-dir", "w", config.Checked.WorkDir, "work directory for run template")
	rootCmd.Flags().StringVar(&config.App.FolderPerm, "perm-folder", config.App.FolderPerm, "create folder permission, default is 0755")
	rootCmd.Flags().StringVar(&config.App.FilePerm, "perm-file", config.App.FilePerm, "create file permission, default is 0644")
	rootCmd.Flags().Int64Var(&config.App.RandomSeed, "random-seed", config.App.RandomSeed, "seed for random function, default is 0 (random by time)")
}

// mugo is the main function for the application.
//
// input is the content, it could be template or data.
// info is the information about the input, it could be file path or url.
func mugo(ctx context.Context, input []byte, info string) (err error) {
	httpReq, err := request.New()
	if err != nil {
		return err
	}

	into.ShutdownAdd(into.FnWarp(shutdown.Global.Run), "shutdown funcs")

	if config.App.RandomSeed != 0 {
		random.DefaultRandom = rand.New(rand.NewSource(config.App.RandomSeed))
	}

	tpl := templatex.New(templatex.WithAddFuncMapWithOpts(FStore)).SetDelims(config.Checked.Delims[0], config.Checked.Delims[1])

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

	fileAPI := file.New()
	output := os.Stdout

	if config.App.Output != "" {
		output, err = fileAPI.OpenFile(
			config.App.Output,
			file.WithFolderPerm(config.App.FolderPerm),
			file.WithFilePerm(config.App.FilePerm),
		)
		if err != nil {
			return err
		}

		defer output.Close()
	}

	// read input data
	var inputData any

	{ // load data
		var storeData any
		for _, data := range config.App.Data {
			var currentData any
			var err error

			if config.App.DataRaw {
				if !config.App.DisableAt && data[0] == '@' {
					d, err := fileAPI.LoadRaw(data[1:])
					if err != nil {
						return fmt.Errorf("failed to load input data: %w", err)
					}

					if config.App.DataRawByte {
						storeData = d
					} else {
						storeData = string(d)
					}
				} else {
					storeData = data
				}

				continue
			}

			if !config.App.DisableAt && data[0] == '@' {
				err = fileAPI.Load(data[1:], &currentData)
			} else {
				err = fileAPI.LoadContent([]byte(data), &currentData, fileAPI.Codec["YAML"])
			}

			if err != nil {
				return fmt.Errorf("failed to load input data: %w", err)
			}

			storeData = mapx.MergeAny(currentData, storeData)
		}

		inputData = storeData
	}

	var templateData []byte
	if config.App.Template != "" {
		if !config.App.DisableAt && config.App.Template[0] == '@' {
			if d, err := fileAPI.LoadRaw(config.App.Template[1:]); err != nil {
				return fmt.Errorf("failed to load template: %w", err)
			} else {
				input = d
			}
		} else {
			templateData = []byte(config.App.Template)
		}
	}

	isInputTemplate := false
	// templateData is not empty than input is data
	if templateData == nil {
		templateData = input
		isInputTemplate = true
	}

	slog.Info("output: " + output.Name())

	if isInputTemplate {
		if config.App.NoStdin && info == os.Stdin.Name() {
			slog.Info("execute template")
		} else {
			slog.Info("execute template: " + info)
		}
	} else if inputData == nil {
		if config.App.DataRaw {
			if config.App.DataRawByte {
				inputData = input
			} else {
				inputData = string(input)
			}
		} else if input != nil {
			if err := fileAPI.LoadContent(input, &inputData, fileAPI.Codec["YAML"]); err != nil {
				return fmt.Errorf("failed to load input data: %w", err)
			}
		}
	}

	// run template
	if err := tpl.Parse(string(templateData)); err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	if err := tpl.Execute(
		templatex.WithIO(output),
		templatex.WithData(inputData),
		templatex.WithParsed(true),
	); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	slog.Info("render completed")

	return nil
}

func ReadAll(r io.Reader) (string, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	return string(content), nil
}

func FStore(opt templatex.Option) map[string]any {
	return fstore.FuncMap(
		fstore.WithSpecificGroups(config.App.SpecificGroups...),
		fstore.WithSpecificFuncs(config.App.SpecificFuncs...),
		fstore.WithDisableGroups(config.App.DisabledGroups...),
		fstore.WithDisableFuncs(config.App.DisabledFuncs...),

		fstore.WithLog(slog.Default()),
		fstore.WithTrust(config.App.Trust),
		fstore.WithWorkDir(config.Checked.WorkDir),
		fstore.WithExecuteTemplate(opt.T),
	)
}
