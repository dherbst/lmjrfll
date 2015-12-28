package lmjrfll

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/api/1/user/profile", Api1UserProfileHandler)
	http.HandleFunc("/api/1/expo/current", Api1ExpoCurrentHandler)
}

// If the user is not logged in, then return the login url.  Otherwise return a json
// structure containing the user's name and email address, and which team they are on.
func Api1UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	fmt.Fprintf(w, "%s", "{}")
}

// Returns the details of the current Expo
func Api1ExpoCurrentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	fmt.Fprintf(w, "%s", "{}")
}
