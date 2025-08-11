package req

import (
	"net/http"
	"restfull_api/pkg/resp"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		resp.Json(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValidate(body)
	if err != nil {
		resp.Json(*w, err.Error(), 402)
		return nil, err
	}

	return &body, nil
}
