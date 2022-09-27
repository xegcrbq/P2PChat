package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
)

type UserRepoSQLX struct {
	db *sqlx.DB
}

func NewUserRepoSQLX(db *sqlx.DB) *UserRepoSQLX {
	return &UserRepoSQLX{
		db: db,
	}
}
func (r *UserRepoSQLX) ReadUserByUserName(command *commands.ReadUserByUserName) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user,
		`select * from users where username=$1;`, command.UserName)
	return &user, err
}
func (r *UserRepoSQLX) ReadUserByUserId(command *commands.ReadUserByUserId) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user,
		`select * from users where userid=$1;`, command.UserId)
	return &user, err
}
