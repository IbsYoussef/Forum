package methods

import "database/sql"

// Checks if User exits
func Checkcredentials(db *sql.DB, email, password string) (bool, error) {
	var count int
	err := db.QueryRow(FIND, email, password).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Checks if email has already been taken
func EmailExists(db *sql.DB, email string) (bool, error) {
	// Check if user email signed up already exists
	var count int
	err := db.QueryRow(FIND, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
