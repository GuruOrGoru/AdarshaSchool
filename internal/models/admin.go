package models

import "net/http"

const (
    AdminEmail    = "mail@mailymailyadmin.com"
    AdminPassword = "woahiamadmin"
    CookieName    = "admin_session"
)

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
