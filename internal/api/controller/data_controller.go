package controller

import (
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
	"github.com/xegcrbq/P2PChat/internal/services"
)

type DataController struct {
	messageService *services.MessageService
	usersService   *services.UserService
}

func NewDataController(messageService *services.MessageService, usersService *services.UserService) *DataController {
	return &DataController{
		messageService: messageService,
		usersService:   usersService,
	}
}
func (c *DataController) Execute(command interface{}) *models.Answer {
	switch command.(type) {
	//messages
	case commands.CreateMessagesByMessage:
		err := c.messageService.CreateMessagesByMessage(command.(commands.CreateMessagesByMessage))
		return &models.Answer{
			Err: err,
		}
	case commands.ReadMessagesByOrderId:
		messages, err := c.messageService.ReadMessagesByOrderId(command.(commands.ReadMessagesByOrderId))
		return &models.Answer{
			Messages: messages,
			Err:      err,
		}
	case commands.ReadMessagesByUserId:
		messages, err := c.messageService.ReadMessagesByUserId(command.(commands.ReadMessagesByUserId))
		return &models.Answer{
			Messages: messages,
			Err:      err,
		}
	case commands.MessageReadNewest:
		message, err := c.messageService.MessageReadNewestBySenderId(command.(commands.MessageReadNewest))
		return &models.Answer{
			Messages: &[]models.Message{*message},
			Err:      err,
		}
	//users
	case commands.ReadUserByUserName:
		user, err := c.usersService.ReadUserByUserName(command.(commands.ReadUserByUserName))
		return &models.Answer{
			User: user,
			Err:  err,
		}
	case commands.ReadUserByUserId:
		user, err := c.usersService.ReadUserByUserId(command.(commands.ReadUserByUserId))
		return &models.Answer{
			User: user,
			Err:  err,
		}
	default:
		return &models.Answer{Err: models.ErrComandNotFound}
	}
}
