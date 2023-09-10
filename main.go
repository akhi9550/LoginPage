package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

type UserSt struct {
	Username string
	Password string
}

var users = make(map[string]UserSt)
var session = make(map[string]string)

type errors struct {
	UsernameError string
	PasswordError string
}

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))

	users["akhu@gmail.com"] = UserSt{"Akhanya", "9552"}
	users["shandu@gmail.com"] = UserSt{"Shandrima", "9551"}
	users["akhil@gmail.com"] = UserSt{"Akhil", "9550"}

}

func main() {

	http.HandleFunc("/", loginPage)
	http.HandleFunc("/home", home)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8000", nil)

}
