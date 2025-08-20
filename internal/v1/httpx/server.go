package httpx

import (
	"fmt"
	"net/http"
	"time"
	"github.com/julienschmidt/httprouter"
)

type Config []struct {
	Name string
	Addr string
	ReadTimeout int
	WriteTimeout int
	IdleTimeout int
	Router *httprouter.Router
	Server *http.Server
}

func	(cfg *Config)LoadFromFile(file_path string) {
	default_routes := "Default"
	if file_path == "Default" {
		cfg.newServerConfig(default_routes, ":8080", default_routes, 5, 10, 60)
		cfg.newServerConfig(default_routes, ":8081", default_routes, 5, 10, 60)
	}
}

func	(cfg *Config)newServerConfig(name, addr, routes string, rd_time, wt_time, idl_time int) {
	*cfg = append(*cfg, struct{
		Name string
		Addr string
		ReadTimeout int
		WriteTimeout int
		IdleTimeout int
		Router *httprouter.Router
		Server *http.Server
	}{
		Name: name,
		Addr: addr,
		ReadTimeout: rd_time,
		WriteTimeout: wt_time,
		IdleTimeout: idl_time,
		Router: Router(routes),
		Server: nil,
	})
}

func	(cfg *Config)CreateAndStartServers() error{
	error_ch := make(chan error)

	for _, n := range *cfg {
		fmt.Println(
			"Server config:", n.Name,
			"\n\tPort: ", n.Addr,
			"\n\tWriteTimeout", n.WriteTimeout,
			"\n\tReadTimeout", n.ReadTimeout,
			"\n\tIdleTimeout", n.IdleTimeout,
			"\n",
		)
		n.Server = &http.Server{
			Addr: n.Addr,
			Handler: n.Router,
			ReadTimeout: time.Second * time.Duration(n.ReadTimeout),
			IdleTimeout: time.Second * time.Duration(n.IdleTimeout),
			WriteTimeout: time.Second * time.Duration(n.WriteTimeout),
		}
		go func(){error_ch <- n.Server.ListenAndServe()}()
	}

	return <- error_ch
}
