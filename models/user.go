package models

type (
	// User represents the structure of our resource
	User struct {
		ID        int64
		FirstName string
		LastName  string
		UserName  string
		Email     string
		Password  string
	}
)
