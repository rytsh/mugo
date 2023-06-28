module github.com/rytsh/mugo

go 1.20

replace (
	github.com/rytsh/mugo/pkg/fstore => ./pkg/fstore
	github.com/rytsh/mugo/pkg/templatex => ./pkg/templatex
)

require (
	github.com/rs/zerolog v1.29.1
	github.com/rytsh/liz/file v0.1.4
	github.com/rytsh/liz/mapx v0.1.1
	github.com/rytsh/liz/shutdown v0.1.0
	github.com/rytsh/mugo/pkg/fstore v0.0.0-00010101000000-000000000000
	github.com/rytsh/mugo/pkg/templatex v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.6.1
	github.com/worldline-go/logz v0.3.3
	github.com/worldline-go/utility/httpx v0.2.3
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.2.0 // indirect
	github.com/Masterminds/sprig/v3 v3.2.3 // indirect
	github.com/cli/safeexec v1.0.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/gomarkdown/markdown v0.0.0-20230322041520-c84983bdbf2a // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.2 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jaswdr/faker v1.18.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/rytsh/call v0.2.1 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tdewolff/minify/v2 v2.12.7 // indirect
	github.com/tdewolff/parse/v2 v2.6.6 // indirect
	github.com/worldline-go/utility/contextx v0.1.0 // indirect
	golang.org/x/crypto v0.10.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/text v0.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
