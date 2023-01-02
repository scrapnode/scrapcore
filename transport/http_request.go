package transport

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

func ParseBody(r *http.Request, dest any) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.NewDecoder(r.Body).Decode(dest)
}
