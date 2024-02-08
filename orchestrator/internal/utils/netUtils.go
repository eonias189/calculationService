package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	c "github.com/eonias189/calculationService/orchestrator/contract"
	"github.com/eonias189/calculationService/orchestrator/internal/config"
)

func ObjToMap(o any) (map[string]string, error) {
	res := make(map[string]string)
	data, err := json.Marshal(o)
	if err != nil {
		return res, err
	}
	resAny := make(map[string]any)
	err = json.Unmarshal(data, &resAny)
	if err != nil {
		return res, err
	}
	for key, val := range resAny {
		res[key] = fmt.Sprint(val)
	}
	return res, nil
}

func applyRestParams(url string, endpoint config.Endpoint, restParams any) string {
	restParamsMap, _ := ObjToMap(restParams)
	toAdd := []string{url}
	for _, param := range endpoint.RestParams {
		value, ok := restParamsMap[param]
		if ok {
			toAdd = append(toAdd, value)
		}
	}
	return strings.Join(toAdd, "/")
}

type RequestParams[B interface{}, RP interface{}, R interface{}] struct {
	Body       B
	RestParams RP
	Response   *R
}

func NewRequestParams[B, RP, R any](body B, restParams RP, response *R) RequestParams[B, RP, R] {
	return RequestParams[B, RP, R]{Body: body, RestParams: restParams, Response: response}
}

func DoRequest[B interface{}, RP interface{}, R interface{}](cli *http.Client, url string, endpoint config.Endpoint, params struct {
	Body       B
	RestParams RP
	Response   *R
}) error {
	path := url + "/" + endpoint.Url

	path = applyRestParams(path, endpoint, params.RestParams)

	bodyBytes, err := json.Marshal(params.Body)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest(endpoint.Method, path, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(path)

	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf(string(respBytes))
	}
	err = json.Unmarshal(respBytes, params.Response)
	return err

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

func SendResponse[R interface{ GetError() error }](resp R, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if resp.GetError() != nil {
		http.Error(w, resp.GetError().Error(), http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(resp)
	fmt.Fprint(w, string(data))
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

func CheckMethodMiddleware(next http.Handler, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.ToUpper(method) != strings.ToUpper(r.Method) {
			SendResponse(&c.ErrorResponse{Error: fmt.Errorf("handling only %v requests", method)}, w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.String())
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				SendResponse(&c.ErrorResponse{Error: fmt.Errorf(fmt.Sprint(err))}, w)

			}
		}()
		next.ServeHTTP(w, r)
	})
}
