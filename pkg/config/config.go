package config

import "github.com/je4/utils/v2/pkg/config"

type ElasticSearchConfig struct {
	Endpoint []string         `toml:"endpoint"`
	Index    string           `toml:"index"`
	ApiKey   config.EnvString `toml:"apikey"`
	Debug    bool             `toml:"debug"`
}
