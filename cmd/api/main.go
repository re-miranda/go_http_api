package main

import (
	"flag"
	"github.com/re-miranda/go_http_api/internal/v1/httpx"
)

var config = flag.String("config", "Default", "Path to server config file")

func main(){
	flag.Parse()
	httpx.Router(*config)
}
