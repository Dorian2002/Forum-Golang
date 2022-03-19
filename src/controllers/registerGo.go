package controllers

import (
	"html/template"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"database/sql"
	"strconv"
	"golang.org/x/crypto/bcrypt"
)

var Ri struct {
	Id int
	WrongPswd bool
	WrongAge bool
	AAE bool
	IsRegistering bool
	IsRegister bool
}

// fonction appelé dans la fonction main, éxécution de la template html
func Register(w http.ResponseWriter, r *http.Request) {
	Ri.Id = 0
	Ri.AAE = false
	Ri.IsRegistering = false
	Ri.WrongPswd = false
	Ri.WrongAge = false
	t, _ := template.ParseFiles("template/register.html")

	name := ""
	pswd := ""
	pswdVerif := ""
	Nom := ""
	Prenom := ""
	Email := ""
	var Age int
	var errAge error

	Age, errAge = strconv.Atoi(r.FormValue("age"))
	Email = r.FormValue("Email")
	Prenom = r.FormValue("Prénom")
	Nom = r.FormValue("Nom")
	name = r.FormValue("username")
	pswd = r.FormValue("password")
	pswdVerif = r.FormValue("passwordV")

	if errAge != nil {
		Ri.WrongAge = true
	}

	if pswd != pswdVerif {
		Ri.WrongPswd = true
	}

	if name == "" || pswd == "" || Nom == "" || Prenom == "" || Email == "" {
		Ri.IsRegistering = false
	}else{
		Ri.IsRegistering = true
	}

	if Ri.IsRegistering == true && Ri.WrongAge == false{
		Ri.Id = 0
		Ri.IsRegister = false
		RegisterQuery(name, pswd, Nom, Email, Prenom, Age)
	
		if Ri.Id == 0 {
			Ri.IsRegister = false
		} else {
			Ri.IsRegister = true
		}
	}
	fmt.Println(Ri.Id)
	t.Execute(w, Ri)
}

func RegisterQuery(username string, password string, Nom string, Email string, Prénom string, Age int) {

	if Ri.WrongPswd == false {
		db, err := sql.Open("sqlite3", "./dbFile/database.db")
		if err != nil {
  			panic(err)
		}
		
		pswd, _ := bcrypt.GenerateFromPassword([]byte(password),10)

		rows, err := db.Query("SELECT id FROM users WHERE username= $1 OR email = $2", username, Email, Nom, Prénom, Age)
		if err != nil{
			panic(err)
		}
		for rows.Next() {
			rows.Scan(&Ri.Id)
		}
		if Ri.Id != 0 {
			Ri.AAE = true
		}
		fmt.Println(Ri.Id)
		if Ri.AAE == false && Ri.IsRegistering == true {
			rows, err = db.Query("INSERT INTO users(email, psswrd, username, Nom, Prénom, Age) VALUES($1, $2, $3, $4, $5, $6);", Email, pswd, username, Nom, Prénom, Age)
			if err != nil{
				fmt.Println(err)
			}

			for rows.Next() {
				rows.Scan(&Ri.Id)
			}
		}
		rows, err = db.Query("SELECT id FROM users WHERE username= $1 AND psswrd = $2 AND email = $3 AND Nom = $4 AND Prénom = $5 AND Age = $6;", username, pswd, Email, Nom, Prénom, Age)
		if err != nil{
			panic(err)
		}
		for rows.Next() {
			rows.Scan(&Ri.Id)
		}
		fmt.Println(Ri.Id)
		db.Close()
	}
}
