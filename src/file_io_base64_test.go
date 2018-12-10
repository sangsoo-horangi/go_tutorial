package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	b64 "encoding/base64"
	"os"
	"io"
)

func main() {
	fmt.Println("testing")
	//importData :=  "root:1@tcp(127.0.0.1:3306)/user"
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
	//fmt.Print("buf ")
	//fmt.Println(string(buff))

	encodeData := b64.StdEncoding.EncodeToString([]byte(buff))
	//fmt.Println(encodeData)
	decodeData, _ := b64.StdEncoding.DecodeString(encodeData)
	//fmt.Println(string(decodeData))
	//db, err := sql.Open("mysql", "root:1@tcp(127.0.0.1:3306)/user")
}
