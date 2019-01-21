package main

import "github.com/gorilla/mux"

import _ "github.com/go-sql-driver/mysql"
import "database/sql"
import "golang.org/x/crypto/bcrypt"
import "encoding/json"
import "net/http"
//import _ "github.com/lib/pq"
const hashCost = 8
var db *sql.DB

type Credentials struct {
	Number int `json:"number", db:"number"`
	Pw string `json:"pw", db:"pw"`
	Id string `json:"id", db:"id"`
}

func connDB()(*sql.DB) {
	db, err:= sql.Open("mysql","root:1@tcp(127.0.0.1:3306)/user")
	if err != nil {
		panic(err)
	}
	return db
}

func Signup(w http.ResponseWriter, r *http.Request) {
	db := connDB()
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Pw),8)

	if _, err = db.Query("insert into Accounts values ($1, $2)", creds.Id, string(hashedPassword)); err != nil {
	w.WriteHeader(http.StatusInternalServerError)
	return
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	db := connDB()
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := db.QueryRow("select pw from Accounts where id=$1",creds.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	storedCreds  := &Credentials{}
	err = result.Scan(&storedCreds.Pw)
	if err != nil {
		if err == sql.ErrNoRows{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Pw), []byte(creds.Pw)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
}


func main() {
	// Signin and Signup handler that we will implement
	port := ":8081"
	r := mux.NewRouter()
	r.HandleFunc("/signin",Signin)
	r.HandleFunc("/singup",Signup)
	// initalize our database connection
	http.ListenAndServe(port,r)
}
// gma.. why not working? 
