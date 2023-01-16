package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/segmentio/ksuid"
)

func NewId(prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, ksuid.New().String())
}

func NewBucketFromTime(template string, t time.Time) (string, int64) {
	return t.Format(template), t.UnixMilli()
}

func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
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
