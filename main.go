package main

import (
	"net/http"

	"github.com/glinboy/gin-oauth2-secured-demo/router"
)

func main() {
	routes := router.NewRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
