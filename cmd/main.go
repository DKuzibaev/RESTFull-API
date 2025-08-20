package main

import (
	"fmt"
	"net/http"
	"restfull_api/configs"
	"restfull_api/internal/auth"
	"restfull_api/internal/link"
	"restfull_api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	// repo
	linkRepository := link.NewLinkRepository(database)

	//handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port :8081")

	server.ListenAndServe()
}
