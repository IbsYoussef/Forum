package main

import (
	"fmt"
	"forum/packages/methods"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", methods.IndexHandler)
	http.HandleFunc("/register", methods.RegisterPageHandler)
	http.HandleFunc("/message", methods.RegisterUserInfo)

	assets := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	fmt.Printf("Listening... on port 👉 http://localhost:8080/ \n")
	fmt.Printf("Use 👉 Control+C to stop server \n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error code: %s", err.Error())
	}
}
