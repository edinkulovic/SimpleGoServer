package models

import (
	"fmt"

	"github.com/edinkulovic/SimpleGoServer/app/db"
)

type (
	// User represents the structure of our resource
	User struct {
		ID        int64
		FirstName string
		LastName  string
		UserName  string
		Email     string
		Password  string
		Age       int
	}
)

// Login method retrieves User if Credentials are Correct
func (u User) Login(username, password string) (User, error) {
	err := db.DB.QueryRow("SELECT ID, Username, FirstName, LastName, Age FROM Users WHERE Username=? AND Password=?", username, password).Scan(
		&u.ID, &u.UserName, &u.FirstName, &u.LastName, &u.Age)

	return u, err
}

// GetByUsername retrieves User by username
func (u User) GetByUsername(username string) (User, error) {
	err := db.DB.QueryRow("SELECT ID, Username, FirstName, LastName, Age FROM Users WHERE Username=?", username).Scan(
		&u.ID, &u.UserName, &u.FirstName, &u.LastName, &u.Age)
	fmt.Println(err)
	return u, err
}
