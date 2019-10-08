package hash

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const size = 12

var MySigningKey = []byte("thisIsMySigningKeyLol")

// Hash - return random hash for string val
func Hash() string {
	b := make([]byte, size)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	fmt.Println(string(b))
	return string(b)
}

// GenerateJWT - generate a jwt token
func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// set claims
	claims["authorized"] = true
	claims["client"] = username
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	// generate the token
	tokenString, err := token.SignedString(MySigningKey)
	if err != nil {
		log.Println("Couldn't generate signed key", err)
	}
	log.Println("JWT token", tokenString)

	return tokenString, err
}

func init() {
	// seed once
	rand.Seed(time.Now().UnixNano())
}