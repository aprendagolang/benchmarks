package main

import (
	"log"

	"github.com/aprendagolang/connections"
)

func main() {
	db, err := connections.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = connections.CreateUserTable()
	if err != nil {
		log.Fatal(err)
	}

	// insert 3 users
	users := []connections.User{
		{Username: "user1", Password: "pass1"},
		{Username: "user2", Password: "pass2"},
		{Username: "user3", Password: "pass3"},
	}

	// sql for insert user
	sql := `INSERT INTO users (username, password) VALUES (?, ?)`

	for _, user := range users {
		_, err = db.Exec(sql, user.Username, user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}
}
