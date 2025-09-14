package models

import (
	"log"
	"net/http"
	"os"
)

var (
	AdminEmail    string
	AdminPassword string
	CookieName    string
)

func InitAdmin() {
	AdminEmail = os.Getenv("ADMIN_EMAIL")
	AdminPassword = os.Getenv("ADMIN_PASSWORD")
	CookieName = os.Getenv("COOKIE_NAME")
	if AdminEmail == "" || AdminPassword == "" || CookieName == "" {
		log.Fatalln("ADMIN_EMAIL, ADMIN_PASSWORD, or COOKIE_NAME not set in environment")
	}
}

func IsAdmin(email, password string) bool {
	return email == AdminEmail && password == AdminPassword
}

func IsLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return false
	}
	return cookie.Value == "true"
}

func SetSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    "true",
		Path:     "/",
		HttpOnly: true,
	})
}

func ClearSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
}
