package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Timeline struct {
	number int
	id string
	pw string
}


func main(){
	db,err := sql.Open("mysql","root:1@tcp(127.0.0.1:3306)/user")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select number,id,pw from Accounts")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		timeline := Timeline {}
		err = rows.Scan(&timeline.number,&timeline.id,&timeline.pw)
		if err != nil {
			panic(err)
		}
		fmt.Println(timeline)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
