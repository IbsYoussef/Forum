package methods

import (
	"net/http"
)

func CreateCookie(w http.ResponseWriter, username string) {
	user := http.Cookie{
		Name:  "session-ID",
		Value: username,
	}
	http.SetCookie(w, &user)
}

// func CreateSession(w http.ResponseWriter, userID int) {
// 	sessionID := uuid.New().String()
// 	cookie := http.Cookie{
// 		Name:     "session-name",
// 		Value:    sessionID,
// 		Path:     "/",
// 		MaxAge:   86400, // 24 hours
// 		HttpOnly: true,
// 		Secure:   true, // Enable only in production with HTTPS
// 	}
// 	http.SetCookie(w, &cookie)
// 	// Store the session ID and user ID in a map
// 	// sessions[sessionID] = userID
// }
