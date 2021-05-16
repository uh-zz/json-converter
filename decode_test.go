package main

import (
	"encoding/json"
	"log"
	"testing"
)

type JsonUser struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func BenchmarkEncodingJson(b *testing.B) {
	var (
		user JsonUser
		str  string = `{"id":5,"name":"hoge","address":"東京"}`
	)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal([]byte(str), &user); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkJsonConverter(b *testing.B) {
	var (
		user User
		str  string = `{"id":5,"name":"hoge","address":"東京"}`
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Decode(&user, str)
	}
}
