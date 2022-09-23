package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/xegcrbq/P2PChat/models"
	"github.com/xegcrbq/P2PChat/models/cmd"
)

type MessageRepoSQLX struct {
	db *sqlx.DB
}

func NewMessageRepoSQLX(db *sqlx.DB) *MessageRepoSQLX {
	return &MessageRepoSQLX{
		db: db,
	}
}
func (r *MessageRepoSQLX) ReadMessagesByOrderId(cmd *cmd.ReadMessagesByOrderId) (*[]models.Message, error) {
	var messages []models.Message
	err := r.db.Select(&messages,
		`select * from messages where orderid=$1;`, cmd.OrderId)
	return &messages, err
}

//func (r *MessageRepoSQLX) CreateMessage(cmd.)
