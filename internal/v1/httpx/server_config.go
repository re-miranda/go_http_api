package httpx

import (
	"encoding/json"
	"os"
	"time"
	"github.com/julienschmidt/httprouter"
)

// JSON reading structs
type Server_config []struct {
	name			string
	host			string
	port			string
	readTimeout		time.Duration
	writeTimeout	time.Duration
	idleTimeout		time.Duration
	router			*httprouter.Router
}

type GlobalConfig struct {
	Name			string
	ReadTimeout		time.Duration
	WriteTimeout	time.Duration
	IdleTimeout		time.Duration
}
type RoutesJSON struct {
	Path	string
	Method	string
	Handler	string
}
type ServerJSON struct {
	Name			string
	Host			string
	Port			string
	ReadTimeout		time.Duration
	WriteTimeout	time.Duration
	IdleTimeout		time.Duration
	Routes			[]RoutesJSON
}
type GlobalJSON struct {
	Global	GlobalConfig
	Servers	[]ServerJSON
}

func	(cfg *Server_config)LoadFromFile(file_path string) error {
	file, err := os.ReadFile(file_path)
	if err != nil {
		return err
	}

	var dat GlobalJSON
	err = json.Unmarshal(file, &dat)
	if err != nil {
		return err
	}

	for _, n := range dat.Servers {
		if n.ReadTimeout == 0 {
			n.ReadTimeout = dat.Global.ReadTimeout
		}
		if n.WriteTimeout == 0 {
			n.WriteTimeout = dat.Global.WriteTimeout
		}
		if n.IdleTimeout == 0 {
			n.IdleTimeout = dat.Global.IdleTimeout
		}
		cfg.newServerConfig(n)
	}

	return nil
}

func	(cfg *Server_config)newServerConfig(config ServerJSON) {
	*cfg = append(*cfg, struct{
		name			string
		host			string
		port			string
		readTimeout		time.Duration
		writeTimeout	time.Duration
		idleTimeout		time.Duration
		router			*httprouter.Router
	}{
		name:			config.Name,
		host:			config.Host,
		port:			config.Port,
		readTimeout:	time.Duration(config.ReadTimeout),
		writeTimeout:	time.Duration(config.WriteTimeout),
		idleTimeout:	time.Duration(config.IdleTimeout),
		router:			newRouter(config.Routes),
	})
}
