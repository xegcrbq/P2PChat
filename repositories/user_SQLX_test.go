package repositories

import (
	"fmt"
	"github.com/xegcrbq/P2PChat/db"
	"github.com/xegcrbq/P2PChat/models/cmd"
	"testing"
)

var userRepo = NewUserRepoSQLX(db.ConnectSQLXTest())

func TestUserRepoSQLX(t *testing.T) {
	fmt.Println(userRepo.ReadUserByUserIdAndPassword(&cmd.ReadUserByUserNameAndPassword{
		UserName: "admin",
		Password: "admin",
	}))
}
