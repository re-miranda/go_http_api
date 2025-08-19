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
	Config_path string;
}

func	CreateAndStartServer(config Config) error{
	fmt.Println(
		"Server config:",
		"\n\tPort: ", config.Addr,
		"\n\tWriteTimeout", config.WriteTimeout,
		"\n\tReadTimeout", config.ReadTimeout,
		"\n\tIdleTimeout", config.IdleTimeout,
	)

	srv := &http.Server{
		Addr: config.Addr,
		Handler: config.Router,
		ReadTimeout: time.Second * time.Duration(config.ReadTimeout),
		IdleTimeout: time.Second * time.Duration(config.IdleTimeout),
		WriteTimeout: time.Second * time.Duration(config.WriteTimeout),
	}

	return srv.ListenAndServe()
}