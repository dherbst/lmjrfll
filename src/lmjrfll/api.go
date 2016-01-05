package lmjrfll

import (
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"net/http"
)

func init() {
	http.HandleFunc("/api/1/user/profile", Api1UserProfileHandler)
	http.HandleFunc("/api/1/expo/current", Api1ExpoCurrentHandler)
}

type UserProfileData struct {
	Name      string
	Email     string
	LogoutUrl string
	LoginUrl  string
	IsAdmin   bool
}

// If the user is not logged in, then return the login url.  Otherwise return a json
// structure containing the user's name and email address, and which team they are on.
func Api1UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	data := UserProfileData{}
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		data.LoginUrl = url
		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Internal Service Error", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s", datajson)
		return
	}
	url, _ := user.LogoutURL(ctx, "/")
	data.LogoutUrl = url
	data.Email = u.Email
	data.IsAdmin = u.Admin
	datajson, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Service Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", datajson)
}

// Returns the details of the current Expo
func Api1ExpoCurrentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	c := appengine.NewContext(r)
	expo, err := GetCurrentExpo(c)
	if err != nil {
		fmt.Fprintf(w, "%s", "{ \"NumTeams\": 0, \"NumRegistered\": 0 }")
		return
	}

	datajson, err := json.Marshal(expo)
	if err != nil {
		http.Error(w, "Internal Service Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", datajson)
}
