package transport

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type Body map[string]interface{}

func (body *Body) FromHttpRequest(r *http.Request) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.NewDecoder(r.Body).Decode(body)
}
func (body *Body) ToString() string {
	bytes, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(body)
	return string(bytes)
}

type Headers struct {
	keyvalues map[string][]string
	keyvalue  map[string]string
}

func (headers *Headers) FromHttpRequest(r *http.Request) error {
	for key, values := range r.Header {
		headers.keyvalues[key] = values
		headers.keyvalue[key] = values[0]
	}
	return nil
}
func (headers *Headers) ToString() string {
	bytes, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(headers.keyvalues)
	return string(bytes)
}
