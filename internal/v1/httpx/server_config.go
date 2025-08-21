package httpx

import (
	"time"
	"github.com/julienschmidt/httprouter"
)

type Server_config []struct {
	name			string
	host			string
	port			string
	logLocation		string
	logLevel		int
	readTimeout		time.Duration
	writeTimeout	time.Duration
	idleTimeout		time.Duration
	router			*httprouter.Router
}

func	(cfg *Server_config)LoadFromFile(file_path string) {
	if file_path == "Default" {
		router := newRouter("router_json")
		cfg.newServerConfig(file_path, "0.0.0.0", "8080", "", 0, 5, 10, 60, router)
	}
}

func	(cfg *Server_config)newServerConfig(name, host, port, logLocation string, logLevel, rd_time, wt_time, idl_time int, router *httprouter.Router) {
	*cfg = append(*cfg, struct{
		name			string
		host			string
		port			string
		logLocation		string
		logLevel		int
		readTimeout		time.Duration
		writeTimeout	time.Duration
		idleTimeout		time.Duration
		router			*httprouter.Router
	}{
		name:			name,
		host:			host,
		port:			port,
		logLocation:	logLocation,
		logLevel:		logLevel,
		readTimeout:	time.Duration(rd_time),
		writeTimeout:	time.Duration(wt_time),
		idleTimeout:	time.Duration(idl_time),
		router:			router,
	})
}
