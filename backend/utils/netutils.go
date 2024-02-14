package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func DoRequest[B, R any](cli *http.Client, url, endpoint, method string, body B, response *R) error {
	path := url + "/" + endpoint

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest(method, path, bodyReader)
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
	err = json.Unmarshal(respBytes, response)
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
	// fmt.Println(string(reqBodyData))

	err = json.Unmarshal(reqBodyData, &body)
	return body, err
}

func SendResponse[R any](resp R, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(resp)
	fmt.Fprint(w, string(data))
}

func SendError(err error, w http.ResponseWriter) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func CheckMethodMiddleware(next http.Handler, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.ToUpper(method) != strings.ToUpper(r.Method) {
			SendError(fmt.Errorf("handling only %v requests", method), w)
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
				SendError(fmt.Errorf(fmt.Sprint(err)), w)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
