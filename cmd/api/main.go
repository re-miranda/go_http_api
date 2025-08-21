package main

import (
	"flag"
	"fmt"
	"github.com/re-miranda/go_http_api/internal/v1/httpx"
)

var config_arg = flag.String("config", "Default", "Path to server config file")

func main(){
	flag.Parse()

	var server_config httpx.Server_config
	err := server_config.LoadFromFile(*config_arg)
	if err != nil {
		fmt.Println("Error loading server config: ", err.Error())
		return
	}

	var server httpx.Server
	err = server.CreateAndStartServers(server_config)
	if err != nil {
		fmt.Println("Error on server startup: ", err.Error())
	}

	fmt.Println("Server shutting down")
}
