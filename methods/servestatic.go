package methods

import "net/http"

func ServeStatic() {
	assets := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
}
