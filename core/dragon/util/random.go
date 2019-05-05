package util

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// generate random string
func RandomStr(length int) string {
	runes := make([]rune, length)
	lettersLen := len(letters)
	for i := 0; i < length; i++ {
		runes[i] = letters[rand.Intn(lettersLen)]
	}
	return string(runes)
}
