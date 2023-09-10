package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var errorStruct errors

func loginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	cookie, err := r.Cookie("session")
	if err == nil && session[cookie.Value] != "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	//getting formvalues to variables
	if r.Method == "POST" {
		username := r.FormValue("userName")
		password := r.FormValue("passWord")
		// Authenticate User
		if _, ok := users[username]; !ok {
			errorStruct.UsernameError = "Invalid Username"
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		if password != users[username].Password {
			errorStruct.PasswordError = "Invalid Password"
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		errorStruct.UsernameError = ""
		errorStruct.PasswordError = ""
		if password == users[username].Password {
			//create cookie and set cookie using uuid or new random session token
			uid := uuid.NewV4().String()
			cookie := &http.Cookie{
				Name:  "session",
				Value: uid,
			}
			http.SetCookie(w, cookie)
			session[cookie.Value] = username
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
	}
	// invalid username or password, redirect to login page
	tmpl.ExecuteTemplate(w, "login.html", errorStruct)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	cookie, err := r.Cookie("session")
	if err != nil || session[cookie.Value] == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	username := session[cookie.Value]
	us := users[username]
	tmpl.ExecuteTemplate(w, "home.html", us)
}

func logout(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	cookie, _ := r.Cookie("session")
	cookie.MaxAge = -1
	session[cookie.Value] = ""
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
