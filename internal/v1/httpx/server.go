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

func	(cfg Config)FromFile(file_path string) {
	if file_path == "Default" {
		cfg.Addr = ":8080"
		cfg.ReadTimeout = 5
		cfg.WriteTimeout = 10
		cfg.IdleTimeout = 60
		cfg.Config_path = file_path
	}
}
