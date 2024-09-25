package util

import (
	"strings"
)

const (
	BASE_58   = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	RANG      = 58
	HASH_SIZE = 7
)

// The base58 encoding format is similar to base62 encoding except that base58 avoids non-distinguishable
// characters such as O (uppercase O), 0 (zero), and I (capital I), l (lowercase L).
//
//	The characters in base62 encoding consume 6 bits (2⁶ = 64).
//	A short URL of 7 characters in length in base62 encoding consumes 42 bits.
func IntToBase58(value int) string {
	sb := strings.Builder{}
	rem := 0
	for i := 0; value > 0; i-- {
		rem = value % RANG
		sb.WriteByte(BASE_58[rem])
		value = value / RANG
	}
	hash := Reverse(sb.String())
	sb.Reset()
	for i := 0; i < HASH_SIZE-len(hash); i++ {
		sb.WriteString("0")
	}
	sb.WriteString(hash)
	return sb.String()
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
