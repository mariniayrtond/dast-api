package model

const GuestUsername = "guest"

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func NewUser(ID string, name string, email string) *User {
	return &User{ID: ID, Name: name, Email: email}
}
