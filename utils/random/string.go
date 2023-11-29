package random

import (
	"math/rand"
	"time"
	"unsafe"
)

const letterBytes = "0123456789abcdef"

var src = rand.NewSource(time.Now().UnixNano())

// RandString generate random string with given length
// ref: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), 4; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), 4
		}
		if idx := int(cache & 15); idx < 16 {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= 4
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
