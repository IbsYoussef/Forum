package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var tpl *template.Template

func init() {
	tpl = template.Must(tpl.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/register", registerPageHandler)
	http.HandleFunc("/message", getUserInfo)

	assets := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	fmt.Printf("Listening... on port 👉 :8080 \n")
	fmt.Printf("Use 👉 Control+C to stop server \n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error code: %s", err.Error())
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Error 404 Page not found", 404)
	} else {
		tpl.ExecuteTemplate(w, "index.html", nil)
	}
}

func registerPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.Error(w, "Error 404 Page not found", 404)
	} else {
		tpl.ExecuteTemplate(w, "register.html", nil)
	}
}

func getUserInfo(w http.ResponseWriter, r *http.Request) {
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
