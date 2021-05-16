package main

import "fmt"

type User struct {
	Id      int    `json-converter:"json:id"`
	Name    string `json-converter:"json:name"`
	Address string `json-converter:"json:address"`
}

func main() {
	var (
		hogeUser User
		hogeStr  string = `{"id":5,"name":"hoge","address":"東京"}`
	)

	Decode(&hogeUser, hogeStr)

	fmt.Println(hogeUser)
}
