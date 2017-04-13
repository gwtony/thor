package config

import (
	"os"
	"fmt"
	//"time"
	"path/filepath"
	"github.com/pelletier/go-toml"
	"github.com/gwtony/thor/gapi/errors"
	"github.com/gwtony/thor/gapi/variable"
)

// Config of server
type Config struct {
	HttpAddr    string  /* http server bind address */
	UdpAddr     string  /* udp server bind address */
	TcpAddr     string  /* tcp server bind address */
	UsocketAddr string  /* usocket server bind address */

	//UdpNFI     string  /* udp server multicast receive interface */

	Location    string  /* handler location */

	Log         string  /* log file */
	Level       string  /* log level */
	RotateLine  int     /* log rotate line */

	File        string  /* config file */
	C           *toml.TomlTree /* toml config */
}

func (conf *Config) SetConf(file string) {
	conf.File = filepath.Join(variable.DEFAULT_CONFIG_PATH, file)
}
// ReadConf reads conf from file
func (conf *Config) ReadConf(file string) error {
	if file == "" {
		if conf.File == "" {
			file = filepath.Join(variable.DEFAULT_CONFIG_PATH, variable.DEFAULT_CONFIG_FILE)
		} else {
			file = conf.File
		}
	}

	//c, err := goconf.ReadConfigFile(file)
	c, err := toml.LoadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] Read conf file %s failed", file)
		return err
	}
	conf.C = c
	return nil
}

// ParseConf parses config
func (conf *Config) ParseConf() error {
	//var err error

	if conf.C == nil {
		fmt.Fprintln(os.Stderr, "[Error] Must read config first")
		return errors.BadConfigError
	}

	item := conf.C.Get("default.http_addr")
	if item == nil {
		conf.HttpAddr = ""
	} else {
		conf.HttpAddr = item.(string)
		fmt.Fprintln(os.Stderr, "[Info] [Default] listen on http addr:", conf.HttpAddr)
	}

	item = conf.C.Get("default.tcp_addr")
	if item == nil {
		conf.TcpAddr = ""
	} else {
		conf.TcpAddr = item.(string)
		fmt.Fprintln(os.Stderr, "[Info] [Default] listen on tcp addr:", conf.TcpAddr)
	}

	item = conf.C.Get("default.udp_addr")
	if item == nil {
		conf.UdpAddr = ""
	} else {
		conf.UdpAddr = item.(string)
		fmt.Fprintln(os.Stderr, "[Info] [Default] listen on udp addr:", conf.UdpAddr)
	}

	item = conf.C.Get("default.usocket_addr")
	if item == nil {
		conf.UsocketAddr = ""
	} else {
		conf.UsocketAddr = item.(string)
		fmt.Fprintln(os.Stderr, "[Info] [Default] listen on usocket addr:", conf.UsocketAddr)
	}

	item = conf.C.Get("default.log")
	if item == nil {
		fmt.Fprintln(os.Stderr, "[Info] [Default] log not found, use default log file")
		conf.Log = ""
	} else {
		conf.Log = item.(string)
	}

	item = conf.C.Get("default.level")
	if item == nil {
		conf.Level = "error"
		fmt.Fprintln(os.Stderr, "[Info] [Default] level not found, use default log level error")
	} else {
		conf.Level = item.(string)
	}

	item = conf.C.Get("default.rotate_line")
	if item == nil {
		fmt.Fprintln(os.Stderr, "[Info] [Default] rotate_line not found, use default", variable.DEFAULT_ROTATE_LINE)
		conf.RotateLine = int(variable.DEFAULT_ROTATE_LINE)
	} else {
		conf.RotateLine = int(item.(int64))
	}

	return nil
}

func (conf *Config) Get(key string) interface{} {
	return conf.C.Get(key)
}
