package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max (for int32 type)
func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb string
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb = sb + string(c)
	}

	return sb
}

// RandomCategory generates a random category name
func RandomCategory() string {
	return RandomString(6)
}
