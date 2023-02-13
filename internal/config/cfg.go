package config

var (
	App     = Config{}
	Checked = Check{
		Delims:  []string{"", ""},
		WorkDir: ".",
	}
)

type Config struct {
	Data      []string
	DataRaw   string
	Parse     []string
	Delims    string
	Output    string
	Silience  bool
	List      bool
	DisableAt bool
}

type Check struct {
	Delims  []string
	WorkDir string
}
