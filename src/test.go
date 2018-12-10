package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	//"os"
	"io/ioutil"
)
func setMd5Hash()string{
	//fp, err := os.Open("./db")
	hasher := md5.Sum([]byte("rootL:1@tcp(127.0.0.1:3306)/user"))
	return hex.EncodeToString(hasher[:])
}

func decryptMD5(){
	fp, err := ioutil.ReadFile("./md5_db")
	if err != nil {
		panic(err)
	}
	//fmt.Println(hex.EncodeToString(fp))
	fmt.Printf("%s",fp)
}
func main() {
	decryptMD5()
}

