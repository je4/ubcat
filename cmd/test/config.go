package main

import (
	"emperror.dev/errors"
	"github.com/BurntSushi/toml"
	"github.com/je4/ubcat/v2/pkg/config"
	"io/fs"
	"os"
)

func LoadTestConfig(fSys fs.FS, fp string, conf *TestConfig) error {
	if _, err := fs.Stat(fSys, fp); err != nil {
		path, err := os.Getwd()
		if err != nil {
			return errors.Wrap(err, "cannot get current working directory")
		}
		fSys = os.DirFS(path)
		fp = "test.toml"
	}
	data, err := fs.ReadFile(fSys, fp)
	if err != nil {
		return errors.Wrapf(err, "cannot read file [%v] %s", fSys, fp)
	}
	_, err = toml.Decode(string(data), conf)
	if err != nil {
		return errors.Wrapf(err, "error loading config file %v", fp)
	}
	return nil
}

type TestConfig struct {
	LogFile       string                     `toml:"logfile"`
	LogLevel      string                     `toml:"loglevel"`
	Template      string                     `toml:"template"`
	Index         string                     `toml:"index"`
	ElasticSearch config.ElasticSearchConfig `toml:"elasticsearch"`
}
