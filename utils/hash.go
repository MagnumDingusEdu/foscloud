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

func CheckPassword(hashedPwd string, plainPwd []byte) (bool, error) {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}

}
