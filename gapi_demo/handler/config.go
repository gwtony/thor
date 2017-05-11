package handler

import (
	"github.com/gwtony/thor/gapi/config"
	"github.com/gwtony/thor/gapi/errors"
)

type DemoConfig struct {
	data int
}

// ParseConfig parses config
func (conf *DemoConfig) ParseConfig(cf *config.Config) error {
	if cf.C == nil {
		return errors.BadConfigError
	}
	data := cf.Get("demo.data")
	conf.data = int(data.(int64))

	return nil
}
