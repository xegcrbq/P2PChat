package commands

import (
	"github.com/xegcrbq/P2PChat/internal/models"
)

type CreateMessagesByMessage struct {
	Message *models.Message
}

type ReadMessagesByOrderId struct {
	OrderId int32
}

type ReadMessagesByUserId struct {
	UserId int32
}
type MessageReadNewest struct {
}
