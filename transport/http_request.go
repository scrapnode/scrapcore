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
