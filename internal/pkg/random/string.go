package random

import "math/rand/v2"

type StringRandom interface {
	StringN(n int) string
}

var charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
var charsetLen = len(charset)

type strrand struct {
	// src := rand.NewSource()
	// rand.NewChaCha8()
	rand *rand.Rand
}

func NewStringRandom() StringRandom {

	// src := rand.NewSource(rand.Int63())
	b := [32]byte{}

	return &strrand{
		rand: rand.New(rand.NewChaCha8(b)),
	}
}

// StringN create random string with n characters
// if n is zero of negative, return empty string
func (r *strrand) StringN(n int) string {
	if n <= 0 {
		return ""
	}
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[0] = charset[r.rand.IntN(charsetLen)]
	}
	return string(s)
}
