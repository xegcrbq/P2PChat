package models

type Answer struct {
	User     *User
	Messages *[]Message
	Err      error
}
