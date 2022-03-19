package controllers

import (
	"html/template"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/contact.html")
	t.Execute(w, nil)
}