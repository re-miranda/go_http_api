package main

import (
	"flag"
	"fmt"
	"github.com/re-miranda/go_http_api/internal/v1/httpx"
)

var config_arg = flag.String("config", "Default", "Path to server config file")

func main(){
	flag.Parse()

	var config httpx.Config
	config.FromFile(*config_arg)
	config.Router = httpx.Router(config)

	fmt.Println("Server is starting")
	err := httpx.CreateAndStartServer(config)
	if err != nil {
		fmt.Println("Error on server startup: ", err)
	}

	fmt.Println("Server shutting down")
}
