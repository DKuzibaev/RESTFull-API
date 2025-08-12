package auth

import (
	"fmt"
	"net/http"
	"restfull_api/configs"
	"restfull_api/pkg/req"
	"restfull_api/pkg/resp"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handle := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handle.Login())
	router.HandleFunc("POST /auth/register", handle.Register())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}

		fmt.Println(body)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		data := LoginResponce{
			Token: "123",
		}
		resp.Json(w, data, 200)
	}
}
