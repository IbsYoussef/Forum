package methods

import (
	"database/sql"
	"html/template"
	"net/http"
)

var tpl *template.Template

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

	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	db, err := sql.Open("sqlite3", "userdb.db")
	if err != nil {
		http.Error(w, "Error opening the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO USERS (username, password, email) VALUES (?,?,?)", username, password, email)
	if err != nil {
		http.Error(w, "Error inserting data into the database", http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "message.html", nil)
}
