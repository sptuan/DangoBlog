package main

import (
	"DangoBlog/data"
	"net/http"
)

// GET /login
// Show the login page
func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// GET /signup
// Show the signup page
func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		logger.Error("[Auth] Cannot parse form")
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		logger.Warning("[Auth] Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		logger.Warning("[Auth] Parse Auth Form failed.")
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			logger.Warning("[Auth] Create Session failed.")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}

}

// GET /logout
// Logs the user out
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		logger.Infof("[LOGOUT] session: %s", cookie.Value)
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
