package methods

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (hashedPassword string) {
	//Bcrypt Password
	Hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error Hashing Password: Error Code %v", err)
	}
	return string(Hash)
}

func CheckPassword(hashedkey string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedkey), []byte(password))
	return err == nil
}
