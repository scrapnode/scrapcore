package transport

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/scrapnode/scrapcore/utils"
	"net/http"
)

func WriteJSON(writer http.ResponseWriter, data any) error {
	writer.Header().Set("content-type", "application/json")

	bytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		return err
	}

	_, err = writer.Write(bytes)
	return err
}

func WriteString(writer http.ResponseWriter, data string) error {
	writer.Header().Set("content-type", "text/plain; charset=utf-8")
	_, err := writer.Write(utils.StringToBytes(data))
	return err
}

func WriteErr400(writer http.ResponseWriter, err error) error {
	// Changing the header map after a call to WriteHeader (or Write) has no effect
	// unless the HTTP status code was of the 1xx class or the modified headers are trailers.
	writer.Header().Set("content-type", "application/json")

	writer.WriteHeader(http.StatusBadRequest)
	data := map[string]interface{}{"error": err.Error()}
	return WriteJSON(writer, data)
}

func WriteErr404(writer http.ResponseWriter, err error) error {
	writer.Header().Set("content-type", "application/json")

	writer.WriteHeader(http.StatusNotFound)
	data := map[string]interface{}{"error": err.Error()}
	return WriteJSON(writer, data)
}

func WriteErr500(writer http.ResponseWriter, err error) error {
	writer.Header().Set("content-type", "application/json")

	writer.WriteHeader(http.StatusInternalServerError)
	data := map[string]interface{}{"error": err.Error()}
	return WriteJSON(writer, data)
}
