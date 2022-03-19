package controllers

import (
	"html/template"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"database/sql"
	"time"
)

type PostContent struct {
	Title  string
	Content string
}

type Data []struct {
	postContent PostContent
}

func Home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/home.html")

	Ucookie, errC := r.Cookie("cookie")
	if errC != nil {
		fmt.Println(errC)
		t, _ = template.ParseFiles("template/home.html")
	}
    CurrentUUID := Ucookie.Value

	if CurrentUUID == "" {
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "cookie", Value: "0", Expires: expiration}
		http.SetCookie(w, &cookie)
	}

	var data [1]PostContent

	tag_viande := r.FormValue("tag2")
	tag_sauce := r.FormValue("tag1")
	tag_vegan := r.FormValue("tag3")
	tag_equipement := r.FormValue("tag4")
	tag_accompagnement := r.FormValue("tag5")
	fmt.Println(tag_viande)
	fmt.Println(tag_sauce)
	fmt.Println(tag_vegan)
	fmt.Println(tag_equipement)
	fmt.Println(tag_accompagnement)

	fmt.Println(CurrentUUID)
	fmt.Println("step 1")

	data = postDspl(tag_viande, tag_sauce, tag_vegan, tag_equipement, tag_accompagnement)
	t.Execute(w, data)
}

func postDspl(viande string, sauce string, vegan string, equipement string, accompagnement string) [1]PostContent {

	var content PostContent

	var data [1]PostContent

	if viande =="" && sauce =="" && equipement=="" && accompagnement=="" && vegan=="" {
		db, err := sql.Open("sqlite3", "./dbFile/database.db")
		if err!= nil {
			fmt.Println(err)
		}
		rows,err2 := db.Query("SELECT * FROM post;")
		if err2!= nil{
			fmt.Println(err)
		}
		for rows.Next() {
			rows.Scan(&content)
		}
		rows.Close()
		db.Close()
	}

	if viande !="" || sauce !="" || equipement !="" || accompagnement !="" || vegan !="" {
		db, err := sql.Open("sqlite3", "./dbFile/database.db")
		if err!= nil {
			fmt.Println(err)
		}
		rows,err2 := db.Query("SELECT * FROM post WHERE tag_viande=$1 AND tag_sauce=$2 AND tag_vegan=$3 AND tag_equipement=$4 AND tag_accompagnement=$5);", viande, sauce, vegan, equipement, accompagnement)
		if err2!= nil{
			fmt.Println(err)
		}
		for rows.Next() {
			rows.Scan(&content)
		}
		rows.Close()
		db.Close()
	}

	data[0] = content

	fmt.Println(data)
	return data
}
