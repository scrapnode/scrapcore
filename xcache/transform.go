package xcache

import jsoniter "github.com/json-iterator/go"

func Encode(value any) ([]byte, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(value)
}

func Decode[T any](data []byte) (*T, error) {
	cursor := new(T)
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, cursor)
	return cursor, err
}
