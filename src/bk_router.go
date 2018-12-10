package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"text/template"
	_ "github.com/go-sql-driver/mysql"
	"os"
	b64 "encoding/base64"
	"io"
	"database/sql"
	"fmt"
)

var mainTemplate, _ = template.ParseFiles("/home/sangsoo/html/main.html")
var loginTemplate, _ = template.ParseFiles("/home/sangsoo/html/bk_login.html")
var enrollTemplate, _ = template.ParseFiles("/home/sangsoo/html/enroll.html")
var forgotTemplate, _ = template.ParseFiles("/home/sangsoo/html/forgot.html")

func forgotPage(w http.ResponseWriter, r *http.Request) {
	forgotTemplate.Execute(w,nil)
}
func mainPage(w http.ResponseWriter, r *http.Request) {
	mainTemplate.Execute(w,nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	loginTemplate.Execute(w,nil)

}

func enrollPage(w http.ResponseWriter, r *http.Request) {
	enrollTemplate.Execute(w,nil)
}


func main() {
	fp, err := os.Open("./b32_db")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	buff := make([]byte, 32)
	
	for {
		cnt, err := fp.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if cnt == 0 {
			break
		}
	}

	encodeData := b64.StdEncoding.EncodeToString([]byte(buff))
	fmt.Println(encodeData)
	decodeData, _ := b64.StdEncoding.DecodeString(encodeData)

	fmt.Println(string(decodeData))
	db, err := sql.Open("mysql", string(decodeData))

	defer db.Close()


	//res, err := stmt.Exec("test","1234")
	//checkErr(err)


	r := mux.NewRouter()
	r.HandleFunc("/main", mainPage)
	r.HandleFunc("/login", loginPage)
	r.HandleFunc("/enroll",enrollPage)
	r.HandleFunc("/forgot",forgotPage)
	http.ListenAndServe(":8081",r)
}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}
