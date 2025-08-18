package httpx

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/re-miranda/go_http_api/internal/v1/httpx/handlers"
)

func Router(config string) error{
	fmt.Println(config)

	router := httprouter.New()
	router.GET("/healthz", handlers.HealthzHandler)
	router.GET("/v1/ping", handlers.PingHandler)
	router.POST("/v1/reverse", handlers.ReverseHandler)

	fmt.Println("Server is starting")
	return http.ListenAndServe(":8080", router)
}