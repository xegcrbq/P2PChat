package controller

import (
	"github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/db/repositories"
	"github.com/xegcrbq/P2PChat/internal/services"
	"github.com/xegcrbq/P2PChat/internal/utils"
)

type Registry struct {
	ChatController   *ChatController
	AuthController   *AuthController
	TalkmeController *TalkmeController
}

func NewRegistry(log *logrus.Entry, repository *db.Repository) *Registry {
	//serviceRegistry := service.NewRegistry(log, repository)

	registry := &Registry{}
	tokenBytes := utils.NewTokenizer([]byte(randstr.Hex(10)))
	ms := services.NewMessageService(repositories.NewMessageRepoSQLX(db.ConnectSQLXTest()))
	us := services.NewUserService(repositories.NewUserRepoSQLX(db.ConnectSQLXTest()))
	dataController := NewDataController(ms, us)
	registry.ChatController = NewChatController(tokenBytes, dataController)
	registry.AuthController = NewAuthController(tokenBytes)

	registry.TalkmeController = NewTalkmeController("xuw9xn7znrz4658f862quecb1p8n1s32vhpo35m61yzrofjepnqk0i2tlum3vhqr", dataController)
	return registry
}
