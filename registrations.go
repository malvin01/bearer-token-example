package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"golang.org/x/crypto/bcrypt"
)

func registerUser(username string, password string) (string, error) {

	db, err := sql.Open("mysql", "root:Mrzero=12@tcp(127.0.0.1:3306)/sample_db")

	if err != nil {

		return "", err

	}

	queryString := "insert into system_users(username, password) values (?, ?)"

	stmt, err := db.Prepare(queryString)

	if err != nil {

		return "", err

	}

	defer stmt.Close()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	_, err = stmt.Exec(username, hashedPassword)

	if err != nil {

		return "", err

	}

	return "Success\r\n", nil

}
