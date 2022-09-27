package models

// Credentials структура для данных пользователя
type User struct {
	UserId   int32
	UserName string
	Password string
	IsTrader bool
}

func (c User) IsValid() bool {
	if c.UserName != "" && c.Password != "" {
		return true
	}
	return false
}
func (c User) Equal(c2 User) bool {
	if c.UserName != c2.UserName {
		return false
	}
	if c.Password != c2.Password {
		return false
	}
	return true
}

type UsersRepository interface {
	SaveCredentials(c *User) error
	ReadCredentialsByUsername(username string) (*User, error)
	DeleteCredentialsByUsername(username string) error
}
