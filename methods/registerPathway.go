package methods

import (
	"database/sql"
	"net/http"
)

var db *sql.DB

var (
	INSERT string = "INSERT INTO USERS (username, password, email) VALUES (?,?,?)"
	FIND   string = "SELECT COUNT(*) FROM users WHERE email = ?"
)

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.Error(w, "Error 404 Page not found", 404)
	} else {
		tpl.ExecuteTemplate(w, "register.html", nil)
	}
}

func RegisterUserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Error Page 404 not found", 404)
		return
	}

	//Check Form Parsing error
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	//Parse user registration details
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	//Open Database
	db, err := sql.Open("sqlite3", "userdb.db")
	if err != nil {
		http.Error(w, "Error opening the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check if the email already exists
	exists, err := EmailExists(db, email)
	if err != nil {
		http.Error(w, "Error checking email existence in the database", http.StatusInternalServerError)
		return
	}

	// If email already exists, return an error
	if exists {
		http.Error(w, "Email is already taken", http.StatusBadRequest)
		return
	}

	//Bcrypt Password
	password = HashPassword(password)

	//Insert Values to Data
	_, err = db.Exec(INSERT, username, password, email)
	if err != nil {
		http.Error(w, "Error inserting data into the database", http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "message.html", nil)
}
