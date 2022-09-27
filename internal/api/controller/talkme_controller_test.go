package controller

import (
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/db/repositories"
	"github.com/xegcrbq/P2PChat/internal/services"
	"testing"
	"time"
)

func TestReadMessagesForPeriod(t *testing.T) {
	ms := services.NewMessageService(repositories.NewMessageRepoSQLX(db.ConnectSQLXTest()))
	us := services.NewUserService(repositories.NewUserRepoSQLX(db.ConnectSQLXTest()))
	dataController := NewDataController(ms, us)
	tc := NewTalkmeController(db.GoDotEnvVariable("XToken"), dataController)
	tc.Update(time.Second*10, false)
	//answ, _ := tc.readMessagesForPeriod(time.Now().Add(-time.Hour*24*6), time.Now())
	//fmt.Println(tc.messagesFromOperator(answ))
}
