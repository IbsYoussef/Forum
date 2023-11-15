package methods

import "net/http"

var (
	SELECT string = "SELECT pass WHERE username, password = (?,?);"
)

func UserLoggedInHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateCookie() {

}
