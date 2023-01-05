package xsender

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type Request struct {
	Uri     string      `json:"uri"`
	Method  string      `json:"method"`
	Headers http.Header `json:"headers"`
	Body    string      `json:"body"`
}

func (req *Request) SetHeaders(data string) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(data), &req.Headers)
}

func (req *Request) GetHeaders() (string, error) {
	bytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(req.Headers)
	return string(bytes), err
}

type Response struct {
	Uri     string      `json:"uri"`
	Status  int         `json:"status"`
	Headers http.Header `json:"headers"`
	Body    string      `json:"body"`
}

func (res *Response) SetHeaders(data string) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(data), &res.Headers)
}

func (res *Response) GetHeaders() (string, error) {
	bytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(res.Headers)
	return string(bytes), err
}
