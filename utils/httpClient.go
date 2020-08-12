package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type ResponseData struct {
	Data interface{} `json:"data"`
}
type ResponseStatus struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_msg"`
}

// Response cnode API response struct
type Response struct {
	ResponseStatus
	ResponseData
}

type IHttpClient interface {
	Get(url string) (interface{}, error)
	Post(url string, body interface{}) (interface{}, error)
	Decode(input, output interface{}) interface{}
}

type HttpClient struct {
	IHttpClient
}

// RequestGet send GET HTTP request
func (h *HttpClient) Get(url string) (interface{}, error) {
	var data Response
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "http.Get(url)")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "ioutil.ReadAll(resp.Body). resp.Body: %+v", resp.Body)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, errors.Wrapf(err, "json.Unmarshal. body: %+v", body)
	}
	if !data.Success {
		return nil, errors.New("API error: " + data.ErrorMessage)
	}
	return data.Data, nil
}

// RequestPost send POST HTTP request
func (h *HttpClient) Post(url string, body interface{}) (interface{}, error) {
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data interface{}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (h *HttpClient) Decode(input, output interface{}) interface{} {
	err := mapstructure.Decode(input, &output)
	if err != nil {
		fmt.Printf("mapstructure.Decode error: %v\n input: %+v\n output: %+v\n", err, input, output)
	}
	return &output
}
