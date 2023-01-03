package transport

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/scrapnode/scrapcore/utils"
	"net/http"
)

func WriteJSON(writer http.ResponseWriter, data any) {
	writer.Header().Set("content-type", "application/json")

	bytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		panic(err)
	}
	if _, err = writer.Write(bytes); err != nil {
		panic(err)
	}
}

func WriteString(writer http.ResponseWriter, data string) {
	writer.Header().Set("content-type", "text/plain; charset=utf-8")
	if _, err := writer.Write(utils.StringToBytes(data)); err != nil {
		panic(err)
	}
}

func WriteErr(writer http.ResponseWriter, err error, status int) {
	// Changing the header map after a call to WriteHeader (or Write) has no effect
	// unless the HTTP status code was of the 1xx class or the modified headers are trailers.
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(status)
	data := map[string]interface{}{"error": err.Error()}

	bytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		panic(err)
	}
	if _, err = writer.Write(bytes); err != nil {
		panic(err)
	}
}

func WriteErr400(writer http.ResponseWriter, err error) {
	WriteErr(writer, err, http.StatusBadRequest)
}

func WriteErr404(writer http.ResponseWriter, err error) {
	WriteErr(writer, err, http.StatusNotFound)
}

func WriteErr500(writer http.ResponseWriter, err error) {
	WriteErr(writer, err, http.StatusInternalServerError)
}
