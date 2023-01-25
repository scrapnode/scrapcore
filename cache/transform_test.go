package cache_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/scrapnode/scrapcore/cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeDecodeInt_Ok(t *testing.T) {
	value := gofakeit.IntRange(1, 100)

	data, err := cache.Encode(value)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	decoded, err := cache.Decode[int](data)
	assert.Nil(t, err)
	assert.Equal(t, value, *decoded)
}

func TestEncodeDecodeString_Ok(t *testing.T) {
	value := gofakeit.Username()

	data, err := cache.Encode(value)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	decoded, err := cache.Decode[string](data)
	assert.Nil(t, err)
	assert.Equal(t, value, *decoded)
}

func TestEncodeDecodeMap_Ok(t *testing.T) {
	value := map[string]interface{}{
		"username": gofakeit.Username(),
		"age":      gofakeit.Float64Range(10, 30),
		"active":   gofakeit.Bool(),
	}

	data, err := cache.Encode(value)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	decoded, err := cache.Decode[map[string]interface{}](data)
	assert.Nil(t, err)
	assert.Equal(t, value, *decoded)
}

type EncodeDecodeStruct struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Active   bool   `json:"active"`
}

func TestEncodeDecodeStruct_Ok(t *testing.T) {
	value := EncodeDecodeStruct{
		Username: gofakeit.Username(),
		Age:      gofakeit.IntRange(10, 30),
		Active:   gofakeit.Bool(),
	}

	data, err := cache.Encode(value)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	decoded, err := cache.Decode[EncodeDecodeStruct](data)
	assert.Nil(t, err)
	assert.Equal(t, value, *decoded)
}
