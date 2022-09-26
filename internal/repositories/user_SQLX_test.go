package repositories

import (
	"fmt"
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/models/cmd"
	"testing"
)

var userRepo = NewUserRepoSQLX(db.ConnectSQLXTest())

func TestUserRepoSQLX(t *testing.T) {
	fmt.Println(userRepo.ReadUserByUserNameAndPassword(&cmd.ReadUserByUserNameAndPassword{
		UserName: "admin",
		Password: "admin",
	}))
	fmt.Println(userRepo.ReadUserByUserId(&cmd.ReadUserByUserId{
		UserId: 2,
	}))
}
