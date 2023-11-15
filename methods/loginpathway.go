package methods

import "net/http"

var (
	SELECT string = "SELECT pass WHERE username, password = (?,?);"
)

func LogInHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logIn" {
		http.Error(w, "Error 404 Page not found", 404)
	} else {
		tpl.ExecuteTemplate(w, "logIn.html", nil)
	}

	
}

func UserLoggedInHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/loggedIn" {
		http.Error(w, "Error 404 Page not found", 404)
	} else {
		tpl.ExecuteTemplate(w, "loggedIn.html", nil)
	}
}

func CreateCookie() {

}
