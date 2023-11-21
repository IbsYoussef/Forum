package main

import (
	"fmt"
	methods "forum/packages/methods"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	methods.Init()

	//Index Page
	http.HandleFunc("/", methods.IndexHandler)

	//Login and Register Page
	http.HandleFunc("/login", methods.LoginPageRegHandler)
	http.HandleFunc("/register", methods.RegisterPageHandler)

	methods.ServeStatic()

	fmt.Printf("Listening... on port 👉 http://localhost:8080/ \n")
	fmt.Printf("Use 👉 Control+C to stop server \n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error code: %s", err.Error())
	}
}
