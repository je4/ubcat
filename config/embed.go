package configdata

import "embed"

//go:embed test.toml
var ConfigFS embed.FS
