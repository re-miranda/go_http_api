package httpx

import "time"

type Server_config []struct {
	name string
	addr string
	readTimeout time.Duration
	writeTimeout time.Duration
	idleTimeout time.Duration
}

func	(cfg *Server_config)LoadFromFile(file_path string) {
	// placeholder test of two servers
	if file_path == "Default" {
		cfg.newServerConfig(file_path, ":8080", 5, 10, 60)
		cfg.newServerConfig(file_path, ":8080", 5, 10, 60)
	}
}

func	(cfg *Server_config)newServerConfig(name, addr string, rd_time, wt_time, idl_time int) {
	*cfg = append(*cfg, struct{
		name string
		addr string
		readTimeout time.Duration
		writeTimeout time.Duration
		idleTimeout time.Duration
	}{
		name: name,
		addr: addr,
		readTimeout: time.Duration(rd_time),
		writeTimeout: time.Duration(wt_time),
		idleTimeout: time.Duration(idl_time),
	})
}
