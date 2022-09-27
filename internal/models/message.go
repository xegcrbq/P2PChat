package models

import (
	"fmt"
	"time"
)

type Message struct {
	MessageId    int32     `db:"messageid"`
	SenderId     int32     `db:"senderid"`
	ReaderId     int32     `db:"readerid"`
	OrderId      int32     `db:"orderid"`
	MessageText  string    `db:"messagetext"`
	AttachmentId int32     `db:"attachmentid"`
	SendTime     time.Time `db:"sendtime"`
	IsRead       bool      `db:"isread"`
	TalkMeId     int32     `db:"talkmeid"`
}

type OldMessage struct {
	Sender string `json:"sender"`
	Target string `json:"target"`
	Text   string `json:"text"`
	Time   int64  `json:"time"`
}

func (m *OldMessage) String() string {
	time := time.UnixMilli(m.Time)
	return fmt.Sprintf("[%v:%v]%v: %v", time.Hour(), time.Minute(), m.Sender, m.Text)
}
