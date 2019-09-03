package hash

import (
	"fmt"
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const size = 12

// Hash - return random hash for string val
func Hash() string {
	b := make([]byte, size)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	fmt.Println(string(b))
	return string(b)
}