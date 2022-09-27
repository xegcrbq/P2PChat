package services

import (
	"github.com/xegcrbq/P2PChat/internal/db/repositories"
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
)

type UserRepo interface {
	ReadUserByUserName(cmd *commands.ReadUserByUserName) (*models.User, error)
	ReadUserByUserId(cmd *commands.ReadUserByUserId) (*models.User, error)
}

type UserService struct {
	userRepo *repositories.UserRepoSQLX
}

func NewUserService(userRepo *repositories.UserRepoSQLX) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) ReadUserByUserId(command commands.ReadUserByUserId) (*models.User, error) {
	user, err := s.userRepo.ReadUserByUserId(&command)
	return user, err
}
func (s *UserService) ReadUserByUserName(command commands.ReadUserByUserName) (*models.User, error) {
	user, err := s.userRepo.ReadUserByUserName(&command)
	return user, err
}
