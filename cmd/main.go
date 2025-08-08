package main

import (
	"fmt"
	"net/http"
	"restfull_api/configs"
	"restfull_api/internal/handler/auth"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port :8081")

	server.ListenAndServe()
}
