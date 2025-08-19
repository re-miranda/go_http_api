package main

import (
	"flag"
	"fmt"
	"github.com/re-miranda/go_http_api/internal/v1/httpx"
)

var config_arg = flag.String("config", "Default", "Path to server config file")

func main(){
	flag.Parse()

	var config httpx.Config = httpx.Config{
		Addr: ":8080",
		ReadTimeout: 5,
		WriteTimeout: 10,
		IdleTimeout: 60,
	}

	router := httpx.Router(*config_arg)

	fmt.Println("Server is starting")
	err := httpx.CreateAndStartServer(config, router)
	if err != nil {
		fmt.Println("Error on server startup: ", err)
	}

	fmt.Println("Server shutting down")
}
