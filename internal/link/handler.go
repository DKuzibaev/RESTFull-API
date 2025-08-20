package link

import (
	"fmt"
	"net/http"
	"restfull_api/pkg/req"
	"restfull_api/pkg/resp"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handle := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}

	router.HandleFunc("POST /link", handle.Create())
	router.HandleFunc("GET /{alias}", handle.GoTo())
	router.HandleFunc("PATCH /link/{id}", handle.Update())
	router.HandleFunc("DELETE /link/{id}", handle.Delete())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		createdLink, err := handler.LinkRepository.Create(link)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp.Json(w, createdLink, 201)
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
