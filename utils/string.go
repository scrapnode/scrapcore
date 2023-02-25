package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"
	"unsafe"

	"github.com/segmentio/ksuid"
	"github.com/sethvargo/go-password/password"
)

func NewId(prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, ksuid.New().String())
}

func NewBucket(template string, t time.Time) (string, int64) {
	return t.Format(template), t.UnixMilli()
}

func MD5(obj any) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	hash := md5.Sum(bytes)
	return hex.EncodeToString(hash[:]), nil
}

func Censor(value string, show int) string {
	if len(value) < show {
		return value
	}

	s := value[0:show]
	hide := strings.Repeat("*", len(value[show:]))
	return s + hide
}

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func RandomString(len int) string {
	return password.MustGenerate(len, int(math.Round(float64(len/2))), 0, false, true)
}
