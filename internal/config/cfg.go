package config

var (
	App     = Config{}
	Checked = Check{
		Delims: []string{"", ""},
	}
)

type Config struct {
	Data     []string
	Delims   string
	Output   string
	Silience bool
	List     bool
}

type Check struct {
	Delims  []string
	WorkDir string
}
