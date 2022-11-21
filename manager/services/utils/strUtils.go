package utils

import (
	"math/rand"
	"time"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
