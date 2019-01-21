package main
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
//import "golang.org/x/crypto/bcrypt"
import "net/http"
import "fmt"
//global values
var db *sql.DB
var err error
type Data struct {
	number int
	id string
	pw string
}

func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "enroll.html")
	} else {
		return
	}

	fmt.Println("Signup page connection")

	fmt.Println("insert success")
}

func loginPage(res http.ResponseWriter, req *http.Request) {
	// my login style is GET
	var username string
	if req.Method != "POST" {
		http.ServeFile(res, req, "login.html")
		fmt.Println("login page connection")
		//username = req.FormValue("id")
		//password = req.FormValue("pw")

			} else {
		fmt.Println("POST connection")
		return
	}

	rows, err := db.Query("SELECT * FROM Accounts")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()


	for rows.Next() {
		data := Data{}
		// select success
		err = rows.Scan(&data.number, &data.id, &data.pw)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(data)
		err = db.QueryRow("SELECT id,pw FROM Accounts WHERE id=?", username).Scan(&data.id,&data.pw)

		if err != nil {
			res.Write([]byte("not correct"))
			http.Redirect(res, req, "/login",301)
			return
		}

		//res.Write([]byte("Hello" + &data.id))

	}

	//err = db.QueryRow("SELECT id,pw FROM Accounts WHERE id=?", username).Scan(&databaseUsername,&databasePassword)
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

	http.HandleFunc("/enroll",signupPage)
	http.HandleFunc("/login",loginPage)
	http.HandleFunc("/",homePage)
	http.ListenAndServe(":8081",nil)
}
