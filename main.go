package main

import (
	"fmt"
	methods "forum/packages/methods"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//Index Page
	http.HandleFunc("/", methods.IndexHandler)

	methods.ServeStatic()

	fmt.Printf("Listening... on port 👉 http://localhost:8080/ \n")
	fmt.Printf("Use 👉 Control+C to stop server \n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error code: %s", err.Error())
	}
}
