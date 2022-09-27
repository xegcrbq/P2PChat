package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
)

type MessageRepoSQLX struct {
	db *sqlx.DB
}

func NewMessageRepoSQLX(db *sqlx.DB) *MessageRepoSQLX {
	return &MessageRepoSQLX{
		db: db,
	}
}
func (r *MessageRepoSQLX) ReadMessagesByOrderId(command *commands.ReadMessagesByOrderId) (*[]models.Message, error) {
	var messages []models.Message
	err := r.db.Select(&messages,
		`select messageid, senderid, readerid, orderid, messagetext, coalesce(attachmentid, 0) as attachmentid, sendtime, isread from messages where orderid=$1;`, command.OrderId)
	return &messages, err
}

func (r *MessageRepoSQLX) ReadMessagesByUserId(command *commands.ReadMessagesByUserId) (*[]models.Message, error) {
	var messages []models.Message
	err := r.db.Select(&messages,
		`select messageid, senderid, readerid, orderid, messagetext, coalesce(attachmentid, 0) as attachmentid, sendtime, isread, talkmeid from messages where senderid=$1 or readerid=$2 ORDER BY SendTime;`, command.UserId, command.UserId)
	return &messages, err
}

func (r *MessageRepoSQLX) MessageReadNewestBySenderId(command *commands.MessageReadNewest) (*models.Message, error) {
	var message models.Message
	err := r.db.Get(&message,
		`select messageid, senderid, readerid, orderid, messagetext, coalesce(attachmentid, 0) as attachmentid, sendtime, isread, talkmeid from messages ORDER BY SendTime desc LIMIT 1;`)
	return &message, err
}

func (r *MessageRepoSQLX) CreateMessage(command *commands.CreateMessagesByMessage) error {
	message := command.Message
	var err error
	if message.AttachmentId == 0 {
		_, err = r.db.Exec(`
	insert into 
	    messages(senderid, readerid, orderid, messagetext, attachmentid, sendtime, talkmeid) 
	VALUES
		($1, $2, $3, $4, null, $5, $6);`,
			message.SenderId, message.ReaderId, message.OrderId, message.MessageText, message.SendTime, message.TalkMeId)
	} else {
		_, err = r.db.Exec(`
	insert into 
	    messages(senderid, readerid, orderid, messagetext, attachmentid, sendtime, talkmeid) 
	VALUES
		($1, $2, $3, $4, $5, $6, $7);`,
			message.SenderId, message.ReaderId, message.OrderId, message.MessageText, message.AttachmentId, message.SendTime, message.TalkMeId)
	}

	return err
}
