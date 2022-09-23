package cmd

import "github.com/xegcrbq/P2PChat/models"

type CreateMessagesByMessage struct {
	Message *models.Message
}

type ReadMessagesByOrderId struct {
	OrderId int32
}
