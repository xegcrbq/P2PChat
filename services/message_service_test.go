package services

import (
	"fmt"
	"github.com/xegcrbq/P2PChat/db"
	"github.com/xegcrbq/P2PChat/repositories"
	"testing"
)

func TestMessageService(t *testing.T) {
	fmt.Println(NewMessageService(repositories.NewMessageRepoSQLX(db.ConnectSQLXTest())))
}
