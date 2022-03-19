package controllers

import (
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "cookie", Value: "0", Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Println("cookieUpdated")
	http.Redirect(w, r, "home", http.StatusSeeOther)
}
