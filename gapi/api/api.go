package api

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"github.com/gwtony/thor/gapi/log"
	"github.com/gwtony/thor/gapi/errors"
	"github.com/gwtony/thor/gapi/utils"
	"github.com/gwtony/thor/gapi/server"
	"github.com/gwtony/thor/gapi/hserver"
	"github.com/gwtony/thor/gapi/config"
	"github.com/gwtony/thor/gapi/variable"
)

type apiContext struct {
	config *config.Config
	server *server.Server
	log    log.Log
}

var api apiContext

//type TcpHandler func(net.Conn, log.Log)
//type UdpHandler func([]byte, int, log.Log)

// Run runs program
func Init(file string) error {
	if utils.ParseOption() < 0 {
		return errors.ParseOptionError
	}

	conf := new(config.Config)
	if file != "" {
		conf.SetConf(file)
	}
	err := conf.ReadConf(*utils.ConfigFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] Read config file failed")
		time.Sleep(variable.DEFAULT_QUIT_WAIT_TIME)
		return errors.ReadConfigError
	}
	err = conf.ParseConf()
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] Parse config file failed")
		time.Sleep(variable.DEFAULT_QUIT_WAIT_TIME)
		return errors.ParseConfigError
	}
	api.config = conf

    rlog := log.GetLogger(conf.Log, conf.Level, conf.RotateLine)
	if rlog == nil {
		fmt.Fprintln(os.Stderr, "[Error] Init log failed")
		time.Sleep(variable.DEFAULT_QUIT_WAIT_TIME)
		return errors.InitLogError
	}
	api.log = rlog

    server, err := server.InitServer(conf, rlog)
    if err != nil {
        rlog.Error("[Error] Init server failed")
		time.Sleep(variable.DEFAULT_QUIT_WAIT_TIME)
        return err
    }

	api.server = server
	return nil
}

func GetConfig()(*config.Config) {
	return api.config
}

func GetLog()(log.Log) {
	return api.log
}

func Run() {
	err := api.server.Run()
	if err != nil {
		time.Sleep(variable.DEFAULT_QUIT_WAIT_TIME)
		return
	}
}

func AddHttpHandler(url string, handler http.Handler) {
	api.server.GetHttpServer().AddRouter(url, handler)
}

func ReturnError(r *http.Request, w http.ResponseWriter, msg string, err error, log log.Log) {
	hserver.ReturnError(r, w, msg, err, log)
}

func ReturnResponse(r *http.Request, w http.ResponseWriter, msg string, log log.Log) {
	hserver.ReturnResponse(r, w, msg, log)
}

func SetConfig(file string) {
	api.config.SetConf(file)
}
