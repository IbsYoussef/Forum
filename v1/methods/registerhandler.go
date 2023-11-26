package methods

import "net/http"

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.Error(w, "Error 404 Page not found", 404)
	} else {
		Tpl.ExecuteTemplate(w, "register.html", nil)
	}
}
