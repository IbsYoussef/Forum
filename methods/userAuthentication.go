package methods

import (
	"database/sql"
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

func EmailExists(db *sql.DB, email string) (bool, error) {
	// Check if user email signed up already exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func Checkcredentials(db *sql.DB, email, password string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ? AND password = ?", email, password).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateCookie() {
}
