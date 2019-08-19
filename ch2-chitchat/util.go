package main

import (
	"errors"
	"net/http"
)

func authenitcate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)
		publicTmplFiles := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html",
		}
		privateTmplFiles := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html",
		}
	}
	var templates *template.Template
	if err != nil {
		templates = template.Must(template.Parse - Files(privateTmplFiles...))
	} else {
		templates = template.Must(template.ParseFiles(publicTmplFiles...))
	}
}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
