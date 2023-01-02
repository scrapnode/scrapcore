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
	writer.WriteHeader(http.StatusBadRequest)
	data := H{"error": err.Error()}
	return WriteJSON(writer, data)
}

func WriteErr404(writer http.ResponseWriter, err error) error {
	writer.WriteHeader(http.StatusNotFound)
	data := H{"error": err.Error()}
	return WriteJSON(writer, data)
}

func WriteErr500(writer http.ResponseWriter, err error) error {
	writer.WriteHeader(http.StatusInternalServerError)
	data := H{"error": err.Error()}
	return WriteJSON(writer, data)
}
