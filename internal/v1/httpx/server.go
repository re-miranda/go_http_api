package httpx

import (
	"fmt"
	"net/http"
	"time"
	"github.com/julienschmidt/httprouter"
)

func	CreateAndStartServer(config string, mux  *httprouter.Router) error{
	fmt.Println(config)

	srv := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	return srv.ListenAndServe()
}