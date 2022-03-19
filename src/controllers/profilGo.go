package controllers

import (
	"html/template"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
)

type UData struct{
	Email string
	Username string
	Nom string
	Prénom string
	Age int
}

func Profil(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/profil.html")

	var data UData

	Ucookie, errC := r.Cookie("cookie")
	CurrentUUID := Ucookie.Value
	if errC != nil {
		fmt.Println(errC)
	}
	if CurrentUUID == "0" {
		http.Redirect(w, r, "login", http.StatusSeeOther)
	}else{
		data = UDataGetter(CurrentUUID)
	}
	
	t.Execute(w, data)
}

func UDataGetter(uuid string) UData{
	db, err := sql.Open("sqlite3", "./dbFile/database.db")
	if err != nil {
  		panic(err)
	}

	var id int
	var data UData

	rows, err := db.Query("SELECT user_id FROM cookie WHERE uuid=$1;", uuid)
	if err != nil{
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&id)
	}
	rows.Close()

	rows, err = db.Query("SELECT email,username,Nom,Prénom,age FROM users WHERE id=$1;", id)
	if err != nil{
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&data.Email, &data.Username, &data.Nom, &data.Prénom, &data.Age)
	}
	fmt.Println(rows)
	rows.Close()

	db.Close()
	
	fmt.Println(id)
	fmt.Println(data)
	return data
}
