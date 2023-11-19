package methods

import "database/sql"

func EmailExists(db *sql.DB, email string) (bool, error) {
	// Check if user email signed up already exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
