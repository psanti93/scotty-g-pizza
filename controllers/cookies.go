package controllers

import (
	"fmt"
	"net/http"
)

const (
	CookieSession = "Session"
)

func newCookie(name string, value string, r *http.Request) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     r.URL.Path,
		HttpOnly: false,
	}
}

func setCookie(w http.ResponseWriter, r *http.Request, name string, value string) {
	cookie := newCookie(name, value, r)
	http.SetCookie(w, cookie)
}

func readCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", fmt.Errorf("Reading cookie error: %v", err)
	}
	token := cookie.Value
	return token, nil
}
