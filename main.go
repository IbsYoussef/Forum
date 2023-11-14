package main

import (
	"fmt"
	"forum/packages/methods"
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
	http.HandleFunc("/register", methods.RegisterPageHandler)
	http.HandleFunc("/message", methods.RegisterUserInfo)

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
