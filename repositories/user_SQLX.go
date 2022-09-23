package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/xegcrbq/P2PChat/models"
	"github.com/xegcrbq/P2PChat/models/cmd"
)

type UserRepoSQLX struct {
	db *sqlx.DB
}

func NewUserRepoSQLX(db *sqlx.DB) *UserRepoSQLX {
	return &UserRepoSQLX{
		db: db,
	}
}
func (r *UserRepoSQLX) ReadUserByUserIdAndPassword(cmd *cmd.ReadUserByUserNameAndPassword) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user,
		`select * from users where username=$1 and password=$2;`, cmd.UserName, cmd.Password)
	return &user, err
}
