package main
import (
	"github.com/gorilla/mux"
	"net/http"
	"text/template"
	"os"
	"io"
	"database/sql"
	"fmt"
	helper "./helper"
	_ "github.com/go-sql-driver/mysql"
	b64 "encoding/base64"
)

type  Data struct {
	number int
	id string
	pw string
}

var mainTemplate, _ = template.ParseFiles("/home/sangsoo/html/main.html")
var loginTemplate, _ = template.ParseFiles("/home/sangsoo/html/bk_login.html")
var enrollTemplate, _ = template.ParseFiles("/home/sangsoo/html/enroll.html")
var forgotTemplate, _ = template.ParseFiles("/home/sangsoo/html/forgot.html")
//var printTemplate, _ = template.ParseFiles("/home/sangsoo/html/print.html")

func connDB() (*sql.DB) {
	db,err := sql.Open("mysql","root:1@tcp(127.0.0.1:3306)/user")
        if err != nil {
                panic(err)
        }
	defer db.Close()
	return db
}

func forgotPage(w http.ResponseWriter, r *http.Request) {
	forgotTemplate.Execute(w,nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	mainTemplate.Execute(w,nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {

	login_info := connDB()

	rows, err := login_info.Query("select * from Accounts")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		data := Data{}
		err = rows.Scan(&data.number, &data.id, &data.pw)
		if err != nil {
			panic(err)
		}
		fmt.Println(data)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("login page success")
	// parsing
	fmt.Println( r.ParseForm())
	userid := r.FormValue("username")
	userpw := r.FormValue("password")
	userid_check := helper.IsEmpty(userid)
	userpw_check := helper.IsEmpty(userpw)
	if userid_check  == true || userpw_check == true {
		fmt.Println("error")
	}
	loginTemplate.Execute(w,nil)
	//findId := db.Query("SELECT id FROM Accounts where ")
}

func enrollPage(w http.ResponseWriter, r *http.Request) {
	enrollTemplate.Execute(w,nil)
}

func printPage(w http.ResponseWriter, r *http.Request) {
}

func main() {
	fp, err := os.Open("./b32_db")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	buff := make([]byte, 32) // length file(./db)  


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
	decodeData, _ := b64.StdEncoding.DecodeString(encodeData)
	db, err := sql.Open("mysql", string(decodeData))

	defer db.Close()

	port := ":8081"
	r := mux.NewRouter()
	r.HandleFunc("/main", mainPage)
	r.HandleFunc("/login", loginPage)
	r.HandleFunc("/enroll",enrollPage)
	r.HandleFunc("/forgot",forgotPage)
	http.ListenAndServe(port,r)
}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}
