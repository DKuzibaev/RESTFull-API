package main

import (
	"fmt"
	"net/http"
	"spm2/internal/handler/auth"
)

func main() {
	//conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port :8081")

	server.ListenAndServe()
}
