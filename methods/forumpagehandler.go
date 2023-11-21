package methods

import "net/http"

func ForumPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/mainpage" {
		http.Error(w, "Error 404 Page not found", 404)
	} else {
		Tpl.ExecuteTemplate(w, "index.html", nil)
	}
}
