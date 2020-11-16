package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(pwd string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
