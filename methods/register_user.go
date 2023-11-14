package methods

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template
var db *sql.DB

func init() {
	tpl = template.Must(tpl.ParseGlob("templates/*.html"))
}

const INSERT = "INSERT INTO USERS (username, password, email) VALUES (?,?,?)"
const FIND = "SELECT COUNT(*) FROM users WHERE email = ?"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Error 404 Page not found", 404)
	} else {
		tpl.ExecuteTemplate(w, "index.html", nil)
	}
}

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

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	db, err := sql.Open("sqlite3", "userdb.db")
	if err != nil {
		http.Error(w, "Error opening the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var count int
	err = db.QueryRow(FIND, email).Scan(&count)
	if err != nil {
		http.Error(w, "Error checking email existence in the database", http.StatusInternalServerError)
		return
	}

	if count > 0 {
		http.Error(w, "Email is already taken", http.StatusBadRequest)
		return
	} else {
		hashedPassword, Herr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if Herr != nil {
			fmt.Printf("Error Hashing Password: Error Code %v", Herr)
		}
		password = string(hashedPassword)

		_, err = db.Exec(INSERT, username, password, email)
		if err != nil {
			http.Error(w, "Error inserting data into the database", http.StatusInternalServerError)
			return
		}
	}

	tpl.ExecuteTemplate(w, "message.html", nil)
}
