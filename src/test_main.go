package main
import (
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/securecookie"
	"net/http"
	"text/template"
	"os"
	"io"
	"encoding/json"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	b64 "encoding/base64"
)

type  Data struct {
	number int
	id string
	pw string
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var mainTemplate, _ = template.ParseFiles("/home/sangsoo/html/main.html")
var loginTemplate, _ = template.ParseFiles("/home/sangsoo/html/bk_login.html")
var enrollTemplate, _ = template.ParseFiles("/home/sangsoo/html/enroll.html")
var forgotTemplate, _ = template.ParseFiles("/home/sangsoo/html/forgot.html")
var loginsuccessTemplate, _ = template.ParseFiles("/home/sangsoo/html/login_success.html")
//var printTemplate, _ = template.ParseFiles("/home/sangsoo/html/print.html")

func connDB() (*sql.DB) {
	db,err := sql.Open("mysql","root:1@tcp(127.0.0.1:3306)/user")
        if err != nil {
                panic(err)
        }
	//defer db.Close()
	return db
}

func forgotPage(w http.ResponseWriter, r *http.Request) {
	forgotTemplate.Execute(w,nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	mainTemplate.Execute(w,nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	loginTemplate.Execute(w,nil)

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
			fmt.Println("test")
			panic(err)
		}
		fmt.Println(data)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}else {
		fmt.Println("login page success")
	}
	// parsing
}

func enrollPage(w http.ResponseWriter, r *http.Request) {
	enrollTemplate.Execute(w,nil)
	db := connDB()
	creds := &Data{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.pw),8)

	_, err = db.Query("INSERT INTO Accounts (id,pw) values ($1, $2)",creds.id,string(hashedPassword))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("insert success")
		return
	}
}

func loginsuccessPage(w http.ResponseWriter, r *http.Request) {
	loginsuccessTemplate.Execute(w,nil)
	fmt.Println("login success!")
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
	r.HandleFunc("/login_success",loginsuccessPage)
	http.ListenAndServe(port,r)
}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}
