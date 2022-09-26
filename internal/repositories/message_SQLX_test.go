package repositories

import (
	"fmt"
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/cmd"
	"testing"
	"time"
)

var messageRepo = NewMessageRepoSQLX(db.ConnectSQLXTest())

func TestMessageRepoSQLX(t *testing.T) {
	message := &models.Message{
		SenderId:     1,
		ReaderId:     2,
		OrderId:      1,
		AttachmentId: 1,
		MessageText:  "hello",
		SendTime:     time.Now(),
	}
	fmt.Println(messageRepo.CreateMessage(&cmd.CreateMessagesByMessage{Message: message}))
	res, _ := messageRepo.ReadMessagesByOrderId(&cmd.ReadMessagesByOrderId{OrderId: 1})
	for _, r := range *res {
		fmt.Println(r, r.AttachmentId == -1)
	}
}
