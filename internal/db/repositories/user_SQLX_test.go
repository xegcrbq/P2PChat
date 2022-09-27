package repositories

import (
	"fmt"
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
	"testing"
)

var userRepo = NewUserRepoSQLX(db.ConnectSQLXTest())

func TestUserRepoSQLX(t *testing.T) {
	fmt.Println(userRepo.ReadUserByUserName(&commands.ReadUserByUserName{
		UserName: "admin",
	}))
	fmt.Println(userRepo.ReadUserByUserId(&commands.ReadUserByUserId{
		UserId: 2,
	}))
}
