package model

/**
* user struct for representing a user in the system.
* This struct is used to store and manipulate user data within the application.
* Using the username as the primary key
 */

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{Name: name}
}
