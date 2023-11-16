package methods

import (
	"database/sql"
	"net/http"
)

var (
	SELECT string = "SELECT USERS WHERE email, password = (?,?);"
)

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logIn" {
		http.Error(w, "Error 404 Page not found", 404)
	} else {
		tpl.ExecuteTemplate(w, "logIn.html", nil)
	}
}

func UserLoggedInHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/loggedIn" {
		http.Error(w, "Error 404 Page not found", 404)
	} else {
		tpl.ExecuteTemplate(w, "loggedIn.html", nil)
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error Loggin In User Credentials", 500)
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	db, err := sql.Open("sqlite3", "userdb.db")
	if err != nil {
		http.Error(w, "Error opening the database", http.StatusInternalServerError)
		return
	}

	check, err := Checkcredentials(db, email, password)
	if err != nil {
		http.Error(w, "Error checking email and password existence in the database", http.StatusInternalServerError)
		return
	}

	if check {
		tpl.ExecuteTemplate(w, "loggedIn.html", nil)
	} else {
		http.Error(w, "Invalid Email or Password", http.StatusInternalServerError)
	}
}
