package httpx

import (
	"net/http"
	"time"
	"fmt"
)

type Server []struct {
	name string
	srv *http.Server
}

func	(server *Server)CreateAndStartServers(server_config Server_config) error{
	error_ch := make(chan error)


	for i,_ := range server_config {
		server.newServer(server_config, i)
		printServerConfig(server_config, i)
		go func(){error_ch <- (*server)[i].srv.ListenAndServe()}()
	}

	return <- error_ch
}

func	(srv *Server)newServer(server_config Server_config, index int) {
	*srv = append(*srv, struct{
		name string
		srv *http.Server
	}{
		name: server_config[index].name,
		srv: &http.Server{
			Addr:			fmt.Sprint(server_config[index].host, ":", server_config[index].port),
			Handler:		server_config[index].router,
			ReadTimeout:	time.Second * server_config[index].readTimeout,
			IdleTimeout:	time.Second * server_config[index].idleTimeout,
			WriteTimeout:	time.Second * server_config[index].writeTimeout,
		},
	})
}

func	printServerConfig(server_config Server_config, index int){
	fmt.Println(
		"Server config:", server_config[index].name,
		"\n\tHost_addr:", server_config[index].host,
		"\n\tPort:", server_config[index].port,
		"\n\tWriteTimeout", int(server_config[index].writeTimeout),
		"\n\tReadTimeout", int(server_config[index].readTimeout),
		"\n\tIdleTimeout", int(server_config[index].idleTimeout),
		"\n",
	)
}
