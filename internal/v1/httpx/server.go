package httpx

import (
	"fmt"
	"net/http"
	"time"
	"github.com/julienschmidt/httprouter"
)

type Config struct {
	Addr string;
	ReadTimeout int;
	WriteTimeout int;
	IdleTimeout int;
	Router *httprouter.Router;
}

func	CreateAndStartServer(config Config) error{
	fmt.Println("Server config: ", config)

	srv := &http.Server{
		Addr: config.Addr,
		Handler: config.Router,
		ReadTimeout: time.Second * time.Duration(config.ReadTimeout),
		IdleTimeout: time.Second * time.Duration(config.IdleTimeout),
		WriteTimeout: time.Second * time.Duration(config.WriteTimeout),
	}

	return srv.ListenAndServe()
}