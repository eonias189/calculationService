package server

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

func HandleError(w http.ResponseWriter, r *http.Request, err error, status int) {
	render.Status(r, status)
	render.JSON(w, r, ErrorResp{Reason: err.Error()})
}

func ParseBody[T any](r *http.Request) (T, error) {
	defer r.Body.Close()
	var body T

	err := render.DecodeJSON(r.Body, &body)
	if err != nil {
		return body, err
	}

	err = validator.New().Struct(body)
	return body, err
}

func GetIntFromQuery(query url.Values, key string, def int) int {
	resStr := query.Get(key)
	if resStr == "" {
		return def
	}

	res, err := strconv.Atoi(resStr)
	if err != nil {
		return def
	}

	return res
}
