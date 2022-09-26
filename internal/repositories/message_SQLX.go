package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/cmd"
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
		`select messageid, senderid, readerid, orderid, messagetext, coalesce(attachmentid, 0) as attachmentid, sendtime, isread from messages where orderid=$1;`, cmd.OrderId)
	return &messages, err
}

func (r *MessageRepoSQLX) CreateMessage(cmd *cmd.CreateMessagesByMessage) error {
	message := cmd.Message
	var err error
	if message.AttachmentId == 0 {
		_, err = r.db.Exec(`
	insert into 
	    messages(senderid, readerid, orderid, messagetext, attachmentid, sendtime) 
	VALUES
		($1, $2, $3, $4, null, $5);`,
			message.SenderId, message.ReaderId, message.OrderId, message.MessageText, message.SendTime)
	} else {
		_, err = r.db.Exec(`
	insert into 
	    messages(senderid, readerid, orderid, messagetext, attachmentid, sendtime) 
	VALUES
		($1, $2, $3, $4, $5, $6);`,
			message.SenderId, message.ReaderId, message.OrderId, message.MessageText, message.AttachmentId, message.SendTime)
	}

	return err
}
