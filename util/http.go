package util

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Request struct
type Request struct {
	Host    string
	Port    string
	Path    string
	Headers map[string]interface{}
	Method  string
	Params  map[string]interface{}
}

// Response struct
type Response struct {
	Body string
}

// SendRequest is to send the request
func (request *Request) SendRequest() string {
	client := &http.Client{}
	// Set params
	params := url.Values{}
	for key, value := range request.Params {
		params.Set(key, value.(string))
	}
	data := params.Encode()
	// Set method, url, params for http request
	req, err := http.NewRequest(request.Method, strings.Join([]string{"http://", request.Host, ":", request.Port, request.Path}, ""), strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	// Set the headers
	for key, value := range request.Headers {
		req.Header.Set(key, value.(string))
	}
	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// Get the body of the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return string(body)
}
