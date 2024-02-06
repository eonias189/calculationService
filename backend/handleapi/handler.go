package handleapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/eonias189/calculationService/backend/config"
)

type Server struct {
	Url string
}

func SetJsonHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func GetBody[B any](r *http.Request) (B, error) {
	var body B

	defer r.Body.Close()
	reqBodyData, err := io.ReadAll(r.Body)
	if err != nil {
		return body, err
	}

	err = json.Unmarshal(reqBodyData, &body)
	return body, err
}

func SendResponse[R any](resp *R, w http.ResponseWriter) error {
	data, err := json.Marshal(*resp)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, string(data))
	return err
}

func ParseRestParams[RP any](r *http.Request, endpoint config.Endpoint) (RP, error) {
	var res RP
	query := r.URL.String()
	params := endpoint.RestParams
	resMap := make(map[string]any)
	queryParams := strings.Split(query, "/")[2:]
	for i := 0; i < min(len(params), len(queryParams)); i++ {
		resMap[params[i]] = queryParams[i]
	}
	data, err := json.Marshal(resMap)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

func LogErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(r.URL.String(), err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
