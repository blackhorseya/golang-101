package stringx

import (
	"strings"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// EncodeBase62 encodes a number to a base62 string.
func EncodeBase62(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(alphabet[n%62]) + s
		n /= 62
	}
	return s
}

// DecodeBase62 decodes a base62 string to a number.
func DecodeBase62(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		n = n*62 + strings.IndexByte(alphabet, s[i])
	}
	return n
}
