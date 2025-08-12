package link

import (
	"fmt"
	"net/http"
)

type LinkHandlerDeps struct {
}

type LinkHandler struct {
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handle := &LinkHandler{}
	router.HandleFunc("POST /link", handle.Create())
	router.HandleFunc("GET /{alias}", handle.GoTo())
	router.HandleFunc("PATCH /link/{id}", handle.Update())
	router.HandleFunc("DELETE /link/{id}", handle.Delete())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}
