package controller

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/db/repositories"
	"github.com/xegcrbq/P2PChat/internal/services"
	"github.com/xegcrbq/P2PChat/internal/utils"
	"time"
)

type Registry struct {
	ChatController   *ChatController
	AuthController   *AuthController
	SocketController *SocketController
}

func NewRegistry(log *logrus.Entry, repository *db.Repository) *Registry {
	//serviceRegistry := service.NewRegistry(log, repository)

	registry := &Registry{}
	//tokenizer := utils.NewTokenizer([]byte(randstr.Hex(10)))
	tokenizer := utils.NewTokenizer([]byte(db.GoDotEnvVariable("TokenizerKey")))
	ms := services.NewMessageService(repositories.NewMessageRepoSQLX(db.ConnectSQLXTest()))
	us := services.NewUserService(repositories.NewUserRepoSQLX(db.ConnectSQLXTest()))
	socketService := services.NewSocketService()
	dataController := NewDataController(ms, us)
	talkmeController := NewTalkMeController(db.GoDotEnvVariable("XToken"), dataController, socketService)
	registry.ChatController = NewChatController(tokenizer, dataController, talkmeController)
	registry.AuthController = NewAuthController(tokenizer, dataController)
	registry.SocketController = NewSocketController(tokenizer, socketService, dataController)
	fmt.Println("talkmeController.Update")
	talkmeController.Update(time.Second*5, true)
	fmt.Println("talkmeController.End")
	return registry
}
