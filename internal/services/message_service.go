package services

import (
	"github.com/xegcrbq/P2PChat/internal/db/repositories"
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
)

type MessageRepo interface {
	ReadMessagesByOrderId(command *commands.ReadMessagesByOrderId) (*[]models.Message, error)
	ReadMessagesByUserId(command *commands.ReadMessagesByUserId) (*[]models.Message, error)
	CreateMessage(command *commands.CreateMessagesByMessage) error
	MessageReadNewestBySenderId(command *commands.MessageReadNewest) (*models.Message, error)
}
type MessageService struct {
	messageRepo *repositories.MessageRepoSQLX
}

func NewMessageService(messageRepo *repositories.MessageRepoSQLX) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

func (s *MessageService) IsMessageAvailable(command commands.ReadMessagesByOrderId) (bool, error) {
	message, err := s.messageRepo.ReadMessagesByOrderId(&command)
	if err == nil && message != nil {
		return true, nil
	}
	return false, err
}
func (s *MessageService) CreateMessagesByMessage(command commands.CreateMessagesByMessage) error {
	err := s.messageRepo.CreateMessage(&command)
	return err
}
func (s *MessageService) ReadMessagesByOrderId(command commands.ReadMessagesByOrderId) (*[]models.Message, error) {
	messages, err := s.messageRepo.ReadMessagesByOrderId(&command)
	return messages, err
}
func (s *MessageService) ReadMessagesByUserId(command commands.ReadMessagesByUserId) (*[]models.Message, error) {
	messages, err := s.messageRepo.ReadMessagesByUserId(&command)
	return messages, err
}
func (s *MessageService) MessageReadNewestBySenderId(command commands.MessageReadNewest) (*models.Message, error) {
	message, err := s.messageRepo.MessageReadNewestBySenderId(&command)
	return message, err
}

//func (s *MessageService) DeleteCredentials(command models.CommandDeleteCredentialsByUsername) error {
//	found, err := s.IsUserAvailable(commands)
//	if !found {
//		return ErrDataNotFound
//	}
//	if err != nil {
//		return err
//	}
//	err = s.credentialsRepo.DeleteCredentialsByUsername(&commands)
//	return err
//}
