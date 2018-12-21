//package main
package handler
import "database/sql"
import "golang.org/x/crypto/bcrypt"
import "encoding/json"
import "net/http"
import _ "github.com/lib/pq"

// create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Number int 'json:"number", db:"number"'
	Pw string 'json:"pw", db:"pw"'
	Id string 'json:"id", db:"id"'
)

func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Pw),8)

	if _, err = db.Query("insert into Accounts values ($1, $2)", creds.Id, string(hassedPassword)); err != nil {
	w.WriteHeader(http.StatusInternalServerError)
	return
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
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
	
