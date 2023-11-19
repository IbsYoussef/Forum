package methods

import "database/sql"

func Checkcredentials(db *sql.DB, email, password string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ? AND password = ?", email, password).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
