package main
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "golang.org/x/crypto/bcrypt"
import "net/http"
import "fmt"
//global values
var db *sql.DB
var err error

func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "signup.html")
		return
	}

	username := req.FormValue("id")
	password := req.FormValue("pw")

	var user string
	err := db.QueryRow("SELECT id from Accounts where id?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("hashedpassword error")
			http.Error(res, "Server Error, unable to create your account.", 500)
			return
		}
		_, err = db.Exec("INSERT into Accounts(id,pw) values(?,?)",username,hashedPassword)
		if err != nil {
			fmt.Println("insert error")
			http.Error(res, "Server Error, unable to create your account.",500)
			return
		}

		res.Write([]byte("User Created"))
		return

	case err != nil:
		fmt.Println("aaaa")
		http.Error(res, "Server error, unable to create your account.",500)
		return
	default:
		http.Redirect(res,req, "/",301)
	}
}

func loginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "login.html")
		return
	}
	username := req.FormValue("id")
	password := req.FormValue("pw")

	var databaseUsername string
	var databasePassword string

	err := db.QueryRow("SELECT id,pw FROM Accounts WHERE id=?", username).Scan(&databaseUsername,&databasePassword)

	if err != nil {
		res.Write([]byte("not correct"))
		http.Redirect(res, req, "/login",301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}
	res.Write([]byte("Hello" + databaseUsername))
}


func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func main() {
	db, err = sql.Open("mysql", "root:1@tcp(127.0.0.1:3306)/user")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/signup",signupPage)
	http.HandleFunc("/login",loginPage)
	http.HandleFunc("/",homePage)
	http.ListenAndServe(":8081",nil)
}
