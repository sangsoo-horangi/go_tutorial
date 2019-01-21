package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func MustConnectDB() {
	if err := ConnectDatabase(); err != nil {
		panic(err)
	}
}


