package useapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/eonias189/calculationService/backend/config"
	types "github.com/eonias189/calculationService/backend/interfaces"
)

type RequestParams[B interface{}, RP interface{}, R interface{}] struct {
	endpoint config.Endpoint
	api      types.ApiMethod[B, RP, R]
}

func NewRequestParams[B interface{}, RP interface{}, R interface{}](endpoint config.Endpoint, apiParams types.ApiMethod[B, RP, R]) RequestParams[B, RP, R] {
	return RequestParams[B, RP, R]{endpoint: endpoint, api: apiParams}
}

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

func DoRequest[B interface{}, RP interface{}, R interface{}](cli *http.Client, url string, endpoint config.Endpoint, apiParams struct {
	Body       B
	RestParams RP
	Response   R
}) error {
	path := url + "/" + endpoint.Url

	path = applyRestParams(path, endpoint, apiParams.RestParams)

	bodyBytes, err := json.Marshal(apiParams.Body)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest(endpoint.Method, path, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	respReader, err := cli.Do(req)
	if err != nil {
		return err
	}
	// fmt.Println(path)

	defer respReader.Body.Close()
	respBytes, err := io.ReadAll(respReader.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBytes, apiParams.Response)
	return err

}
