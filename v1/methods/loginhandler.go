package methods

import "net/http"

func LoginPageRegHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "Error 404 page not found", 404)
	} else {
		Tpl.ExecuteTemplate(w, "login.html", nil)
	}
}
