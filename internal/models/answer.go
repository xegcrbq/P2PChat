package models

type Answer struct {
	User     *User
	UserName string
	Messages *[]Message
	Err      error
}
