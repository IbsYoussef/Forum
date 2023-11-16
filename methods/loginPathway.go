package methods

import (
	"net/http"
)

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/log_in" {
		http.Error(w, "Error 404 Page not found", 404)
	} else {
		tpl.ExecuteTemplate(w, "log_in.html", nil)
	}
}
