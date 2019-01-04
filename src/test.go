package main
import "fmt"
import _ "github.com/go-sql-driver/mysql"
import "database/sql"

func main() {
	fmt.Println("Go Mysql tutorial")

	db, err := sql.Open("mysql","root:1@tcp(127.0.0.1:3306)/user")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("connect success")

	insert, err := db.Query("INSERT INTO Accounts VALUES('4','testid','hhhh')")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("INSERT success")
}
