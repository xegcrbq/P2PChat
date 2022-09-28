package services

import (
	"fmt"
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/db/repositories"
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
	"testing"
)

func TestCreateUser(t *testing.T) {
	us := NewUserService(repositories.NewUserRepoSQLX(db.ConnectSQLXTest()))
	s, err := us.CreateUserByUser(commands.CreateUserByUser{
		User: &models.User{
			UserName: "admin",
			Password: "s",
		},
	})
	fmt.Println(s)
	fmt.Println(err)
}
