package auth

import (
	// standard library
	// "fmt"
	// "bytes"
	"net/http"
	// "reflect"
    "time"
	//imported libraries
	// "github.com/gorilla/sessions"
	// project files
	// "github.com/patientplatypus/webserver/utilities"

)

func addCookie(w http.ResponseWriter, email string, password string) {
	expire := time.Now().AddDate(0, 0, 29)
    cookieEmail := http.Cookie{Name: "zennifyEmail", Value: email, Expires: expire, Path: "/", MaxAge: 90000}
	http.SetCookie(w, &cookieEmail)
	cookiePassword := http.Cookie{Name: "zennifyPassword", Value: password, Expires: expire, Path: "/", MaxAge: 90000}
	http.SetCookie(w, &cookiePassword)
}

func CookiesMain(email string, password string, loginType string, w http.ResponseWriter) {
	addCookie(w, email, password)
}

