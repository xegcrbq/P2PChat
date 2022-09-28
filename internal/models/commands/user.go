package commands

import "github.com/xegcrbq/P2PChat/internal/models"

type ReadUserByUserName struct {
	UserName string
}

type ReadUserByUserId struct {
	UserId int32
}
type CreateUserByUser struct {
	User *models.User
}
