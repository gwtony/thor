package server

import (
	//"time"
	"github.com/gwtony/thor/gapi/hserver"
	"github.com/gwtony/thor/gapi/config"
	"github.com/gwtony/thor/gapi/errors"
	"github.com/gwtony/thor/gapi/log"
)

// Server is A HTTP server
type Server struct {
	haddr   string

	hsch    chan int

	hs      *hserver.HttpServer

	log     log.Log
}

// InitServer inits server
func InitServer(conf *config.Config, log log.Log) (*Server, error) {
	s := &Server{}

	s.log = log

	s.hsch = make(chan int, 1)

	if conf.HttpAddr != "" {
		s.haddr = conf.HttpAddr
		hs, err := hserver.InitHttpServer(conf.HttpAddr, s.log)
		if err != nil {
			s.log.Error("Init http server failed")
				return nil, err
		}
		s.hs = hs
	}

	if s.hs == nil {
		s.log.Error("No server inited")
		return nil, errors.InitServerError
	}

	s.log.Debug("Init server done")

	//modules.InitModules(conf, hs, log)

	return s, nil
}

// Run starts server
func (s *Server) Run() error {
	if s.hs != nil {
		go s.hs.Run(s.hsch)
	}

	//TODO: monitor or something
	select {
		case <-s.hsch:
			s.log.Error("http server run failed")
			break
	}

	return nil
}

func (s *Server) GetHttpServer() (*hserver.HttpServer) {
	return s.hs
}
