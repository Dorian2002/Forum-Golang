package controllers

import (
	"html/template"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"time"
)

var post struct {
	TitleTooLong bool
	TitleTooShort bool
	ContentTooLong bool
}

func PostCreate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/postCreate.html")
	Ucookie, errC := r.Cookie("cookie")
	CurrentUUID := Ucookie.Value
	if errC != nil {
		fmt.Println(errC)
		t, _ = template.ParseFiles("template/postCreate.html")
	}
	if CurrentUUID == "0" {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	}
    
	var id int
	if CurrentUUID != "" {
		db, err := sql.Open("sqlite3", "./dbFile/database.db")
		rows, err := db.Query("SELECT user_id FROM cookie WHERE uuid=$1;", CurrentUUID)
		if err != nil{
			fmt.Println(err)
		}
		for rows.Next() {
			rows.Scan(&id)
		}
		rows.Close()
		db.Close()
	}

	var contentTooLong bool
	var titleTooLong bool
	var titleTooShort bool

	title := r.FormValue("title")
	content := r.FormValue("content")
	tag_viande := r.FormValue("tag2")
	tag_sauce := r.FormValue("tag1")
	tag_vegan := r.FormValue("tag3")
	tag_equipement := r.FormValue("tag4")
	tag_accompagnement := r.FormValue("tag5")

	if len(title) <=0 {
		titleTooShort = true
	}else if len(title) >=32 {
		titleTooLong = true
	}

	if len(content) >=2000 {
		contentTooLong = true
	}

	now := time.Now()
	time := string(now.Format("2006-01-02"))

	post.TitleTooLong = titleTooLong
	post.TitleTooShort = titleTooShort
	post.ContentTooLong = contentTooLong

	if !contentTooLong && !titleTooLong && !titleTooShort {
		db, err := sql.Open("sqlite3", "./dbFile/database.db")
		db.Exec("INSERT INTO post(user_id, title, content, created_at, tag_viande, tag_sauce, tag_accompagnement, tag_vegan, tag_equipement) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);",id , title, content, time, tag_viande, tag_sauce, tag_accompagnement, tag_vegan, tag_equipement)
		if err != nil{
			fmt.Println(err)
		}
		db.Close()
	}
	t.Execute(w, post)
}
