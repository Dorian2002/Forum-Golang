package controllers

import (
	"html/template"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"time"
	"github.com/google/uuid"
)

var Cookie struct {
	uuid 	string
	name    string
	user_id int
}

var row struct {
	Id int
	Pswd string
	IsLoging bool
	IsLog bool
}

func LogIn(w http.ResponseWriter, r *http.Request) {

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "cookie", Value: "0", Expires: expiration}
	http.SetCookie(w, &cookie)

	Cookie.user_id = 0
	row.IsLoging = false
	cookieExist := true
	t, _ := template.ParseFiles("template/login.html")
	name := ""
	pswd := ""
	name = r.FormValue("username")
	pswd = r.FormValue("password")

	if name == "" || pswd == ""{
		row.IsLoging = false
	}else{
		row.IsLoging = true
	}
	
	if row.IsLoging == true {
		row.Id = 0
		row.IsLog = false
		cookieExist = ConnectQuery(name, pswd)
	
		if row.Id == 0 {
			row.IsLog = false
		} else {
			row.IsLog = true
		}
	}
	
	if cookieExist == false{
		db, err := sql.Open("sqlite3", "./dbFile/database.db")
		if err != nil {
  			panic(err)
		}
		UUIDtype := uuid.New()
		uuid := UUIDtype.String()
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "cookie", Value: uuid, Expires: expiration}
		db.Exec("INSERT INTO cookie(uuid, name, user_id) VALUES($1, $2, $3);", uuid, "cookie", row.Id)
		http.SetCookie(w, &cookie)
		fmt.Println("cookieCreated")
	}
	if cookieExist == true{
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "cookie", Value: Cookie.uuid, Expires: expiration}
		http.SetCookie(w, &cookie)
		fmt.Println("cookieUpdated")
	}
	
	Cookie.uuid = "0"
	t.Execute(w, row)

}

func ConnectQuery(username string, password string) bool {

	cookieExist := true
	fmt.Println(cookieExist)

	db, err := sql.Open("sqlite3", "./dbFile/database.db")
	if err != nil {
  		panic(err)
	}
	var mdp string
	rows, err := db.Query("SELECT psswrd FROM users WHERE username=$1;", username)
	if err != nil{
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&mdp)
	}
	rows.Close()

	P := []byte(mdp)

	if bcrypt.CompareHashAndPassword(P, []byte(password)) == nil {
		fmt.Println("right password")
		cookieExist = false
		var id int
		rows, err = db.Query("SELECT id FROM users WHERE username=$1;", username)
		if err != nil{
			fmt.Println(err)
		}
		for rows.Next() {
			rows.Scan(&id)
		}
		rows.Close()
		
		row.Id = id

		var uuid string
		rows, err = db.Query("SELECT uuid FROM cookie WHERE user_id=$1;", id)
		for rows.Next() {
			rows.Scan(&uuid)
		}
		rows.Close()
		fmt.Print("uuid : ")
		fmt.Println(uuid)
		fmt.Print("id : ")
		fmt.Println(id)
		Cookie.uuid = uuid
		if uuid != "" {
			cookieExist = true
		}
	}

	db.Close()
	return cookieExist
}
