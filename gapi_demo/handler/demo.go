package handler

import (
	"github.com/gwtony/thor/gapi/log"
	"github.com/gwtony/thor/gapi/config"
)

// InitContext inits demo context
func InitContext(conf *config.Config, log log.Log) error {
	cf := &DemoConfig{}
	err := cf.ParseConfig(conf)
	if err != nil {
		log.Error("Demo parse config failed")
		return err
	}
	log.Info("Demo parse config done: data is %d", cf.data)

	return nil
}


