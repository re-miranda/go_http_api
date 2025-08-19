package main

import (
	"flag"
	"fmt"
	"github.com/re-miranda/go_http_api/internal/v1/httpx"
)

var config = flag.String("config", "Default", "Path to server config file")

func main(){
	flag.Parse()

	router := httpx.Router(*config)

	fmt.Println("Server is starting")

	err := httpx.CreateAndStartServer(*config, router)
	if err != nil {
		fmt.Println("Error on server startup: ", err)
	}

	fmt.Println("Server shutting down")
}
