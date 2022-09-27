package controller

import (
	"github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/db/repositories"
	"github.com/xegcrbq/P2PChat/internal/services"
	"github.com/xegcrbq/P2PChat/internal/utils"
	"time"
)

type Registry struct {
	ChatController *ChatController
	AuthController *AuthController
}

func NewRegistry(log *logrus.Entry, repository *db.Repository) *Registry {
	//serviceRegistry := service.NewRegistry(log, repository)

	registry := &Registry{}
	tokenBytes := utils.NewTokenizer([]byte(randstr.Hex(10)))
	ms := services.NewMessageService(repositories.NewMessageRepoSQLX(db.ConnectSQLXTest()))
	us := services.NewUserService(repositories.NewUserRepoSQLX(db.ConnectSQLXTest()))
	dataController := NewDataController(ms, us)
	talkmeController := NewTalkmeController(db.GoDotEnvVariable("XToken"), dataController)
	registry.ChatController = NewChatController(tokenBytes, dataController, talkmeController)
	registry.AuthController = NewAuthController(tokenBytes)
	talkmeController.Update(time.Second*10, false)
	return registry
}
