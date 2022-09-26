package controller

import (
	"github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/utils"
)

type Registry struct {
	ChatController *ChatController
	AuthController *AuthController
}

func NewRegistry(log *logrus.Entry, repository *db.Repository) *Registry {
	//serviceRegistry := service.NewRegistry(log, repository)

	registry := &Registry{}

	registry.ChatController = NewChatController(utils.NewTokenizer([]byte(randstr.Hex(10))))
	registry.AuthController = NewAuthController(utils.NewTokenizer([]byte(randstr.Hex(10))))
	return registry
}
