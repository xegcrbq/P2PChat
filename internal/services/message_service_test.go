package services

import (
	"fmt"
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/repositories"
	"testing"
)

func TestMessageService(t *testing.T) {
	fmt.Println(NewMessageService(repositories.NewMessageRepoSQLX(db.ConnectSQLXTest())))
}
