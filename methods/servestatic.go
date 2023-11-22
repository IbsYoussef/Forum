package methods

import "net/http"

func ServeStatic() {
	assets := http.FileServer(http.Dir("assets"))
	static := http.FileServer(http.Dir("static"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.Handle("/static/", http.StripPrefix("/static/", static))
}
