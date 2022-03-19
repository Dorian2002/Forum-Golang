package main

import (
	"net/http"
	"./controllers"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	fs := http.FileServer(http.Dir("./template/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/profil",controllers.Profil)
	http.HandleFunc("/contact",controllers.Contact)
	http.HandleFunc("/register",controllers.Register)
	http.HandleFunc("/login",controllers.LogIn)
	http.HandleFunc("/posts",controllers.Posts)
	http.HandleFunc("/postCreate",controllers.PostCreate)
	http.HandleFunc("/logout",controllers.Logout)
	http.HandleFunc("/", controllers.Home)
	http.ListenAndServe("localhost:1111", nil)
}
