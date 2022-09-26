package services

import (
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/cmd"
	"github.com/xegcrbq/P2PChat/internal/repositories"
)

type MessageRepo interface {
	ReadMessagesByOrderId(cmd *cmd.ReadMessagesByOrderId) (*[]models.Message, error)
	CreateMessage(cmd *cmd.CreateMessagesByMessage) error
}
type MessageService struct {
	messageRepo *repositories.MessageRepoSQLX
}

func NewMessageService(messageRepo *repositories.MessageRepoSQLX) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

//func (s *CredentialsService) IsUserAvailable(cmd models.CommandDeleteCredentialsByUsername) (bool, error) {
//	session, err := s.credentialsRepo.ReadCredentialsByUsername(&models.QueryReadCredentialsByUsername{Username: cmd.Username})
//	if err == nil && session != nil {
//		return true, nil
//	}
//	return false, err
//}
//func (s *CredentialsService) GetCredentials(cmd models.QueryReadCredentialsByUsername) (*models.Credentials, error) {
//	session, err := s.credentialsRepo.ReadCredentialsByUsername(&cmd)
//	return session, err
//}
//func (s *CredentialsService) InsertCredentials(cmd models.CommandCreateCredentials) error {
//	err := s.credentialsRepo.SaveCredentials(&cmd)
//	return err
//}
//func (s *CredentialsService) DeleteCredentials(cmd models.CommandDeleteCredentialsByUsername) error {
//	found, err := s.IsUserAvailable(cmd)
//	if !found {
//		return ErrDataNotFound
//	}
//	if err != nil {
//		return err
//	}
//	err = s.credentialsRepo.DeleteCredentialsByUsername(&cmd)
//	return err
//}
