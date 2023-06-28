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
	DisableAt    bool
	Trust        bool
	SkipVerify   bool
	DisableRetry bool

	FolderPerm string
	FilePerm   string

	List           bool
	SpecificGroups []string
	SpecificFuncs  []string
	DisabledGroups []string
	DisabledFuncs  []string

	HtmlTemplate bool
}

type Check struct {
	Delims  []string
	WorkDir string
}
