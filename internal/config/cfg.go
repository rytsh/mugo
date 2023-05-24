package config

var (
	App = Config{
		LogLevel: "info",
	}
	Checked = Check{
		Delims:  []string{"", ""},
		WorkDir: "",
	}
)

type Config struct {
	LogLevel     string
	Data         []string
	DataRaw      string
	Parse        []string
	Delims       string
	Output       string
	Silience     bool
	List         bool
	DisableAt    bool
	Trust        bool
	SkipVerify   bool
	DisableRetry bool

	FolderPerm string
	FilePerm   string
}

type Check struct {
	Delims  []string
	WorkDir string
}
