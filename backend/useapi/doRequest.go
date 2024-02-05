package useapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/eonias189/calculationService/backend/config"
)

type RequestParams[B interface{}] struct {
	Endpoint config.Endpoint
	Body     B
}

type None struct{}

func applyQueryParams(url string, params ...string) string {
	pattern, _ := regexp.Compile("(<[^>]+>)")
	toReplace := 0
	for pattern.MatchString(url) {
		if len(params) == toReplace {
			return url
		}
		matched := pattern.FindString(url)
		url = strings.Replace(url, matched, params[toReplace], 1)
		toReplace++
	}
	return url

}

func DoRequest[B interface{}, R interface{}](cli *http.Client, url string, params RequestParams[B], queryParams ...string) (R, error) {
	var resp R

	path := url + "/" + params.Endpoint.Url
	if params.Endpoint.Method == "GET" {
		path = applyQueryParams(path, queryParams...)
	}

	bodyBytes, err := json.Marshal(params.Body)
	if err != nil {
		return resp, err
	}

	bodyReader := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest(params.Endpoint.Method, path, bodyReader)
	if err != nil {
		return resp, err
	}

	req.Header.Set("Content-Type", "application/json")
	respReader, err := cli.Do(req)
	if err != nil {
		return resp, err
	}

	defer respReader.Body.Close()
	respBytes, err := io.ReadAll(respReader.Body)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(respBytes, &resp)
	return resp, err

}
