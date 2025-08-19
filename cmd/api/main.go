package main

import (
	"flag"
	"fmt"

	"github.com/re-miranda/go_http_api/internal/v1/httpx"
)

var config = flag.String("config", "Default", "Path to server config file")

func main(){
	flag.Parse()

	done := make(chan string)
	go httpx.Router(*config, done)
	fmt.Println(<-done)

	for {}
}
