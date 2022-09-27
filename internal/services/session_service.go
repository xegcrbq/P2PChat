package services

//
//import (
//	"github.com/xegcrbq/auth/models"
//	"github.com/xegcrbq/auth/repositories"
//)
//
//type SessionService struct {
//	sessionRepo repositories.SessionRepo
//}
//
//func NewSessionService(sessionRepo repositories.SessionRepo) *SessionService {
//	return &SessionService{
//		sessionRepo: sessionRepo,
//	}
//}
//
//func (s *SessionService) IsSessionAvailable(commands models.CommandDeleteSessionByRefreshToken) (bool, error) {
//	session, err := s.sessionRepo.ReadSessionByRefreshToken(&models.QueryReadSessionByRefreshToken{RefreshToken: commands.RefreshToken})
//	if err == nil && session != nil {
//		return true, nil
//	}
//	return false, err
//}
//func (s *SessionService) GetSession(commands models.QueryReadSessionByRefreshToken) (*models.Session, error) {
//	session, err := s.sessionRepo.ReadSessionByRefreshToken(&commands)
//	return session, err
//}
//func (s *SessionService) InsertSession(commands models.CommandCreateSession) error {
//	err := s.sessionRepo.SaveSession(&commands)
//	return err
//}
//func (s *SessionService) DeleteSession(commands models.CommandDeleteSessionByRefreshToken) error {
//	found, err := s.IsSessionAvailable(commands)
//	if !found {
//		return ErrDataNotFound
//	}
//	if err != nil {
//		return err
//	}
//	err = s.sessionRepo.DeleteSessionByRefreshToken(&commands)
//	return err
//}
