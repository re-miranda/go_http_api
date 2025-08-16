package main

import (
	"fmt"
	"flag"
	"github.com/re-miranda/go_http_api/internal/v1/httpx"
)

var config_file = flag.String("config_file", "Default", "Path to server config file")

func main(){
	flag.Parse()
	fmt.Println(*config_file)
	httpx.Router()
}
