package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

func NewId(prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, ksuid.New().String())
}

func NewBucket(template string) (string, int64) {
	return NewBucketFromTime(template, time.Now().UTC())
}

func NewBucketFromTime(template string, t time.Time) (string, int64) {
	return t.Format(template), t.UnixMilli()
}

func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
