package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxz"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

// Generates a random integer between min and max values
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generates a random string of 'n' characters
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Generates a random username
func RandomUserName() string {
	return RandomString(6)
}

// Generates a random full name
func RandomName() string {
	return RandomString(6) + " " + RandomString(6)
}

// Generates a random email
func RandomEmail() string {
	return RandomString(6) + "@" + RandomString(4) + ".com"
}

// Generates a hashed and salted password
func RandomPassword() string {
	return HashAndSalt(RandomString(10))
}

// Generate a random link
func RandomLink() string {
	return "https://example.com/" + RandomString(6)
}

// Generate a random filesize
func RandomFilesize() int64 {
	return RandomInt(1,1000)
}