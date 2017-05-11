package main

import (
	"fmt"
	"github.com/gwtony/thor/gapi/api"
	"github.com/gwtony/thor/gapi_demo/handler"
)

func main() {
	err := api.Init("gapi_demo.conf")
	if err != nil {
		fmt.Println("[Error] Init api failed:", err)
		return
	}
	config := api.GetConfig()
	log := api.GetLog()

	err = handler.InitContext(config, log)
	if err != nil {
		fmt.Println("[Error] Init demo failed:", err)
		return
	}

	api.Run()
}
