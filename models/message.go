package models

import (
	"fmt"
	"time"
)

type Message struct {
	Sender string `json:"sender"`
	Target string `json:"target"`
	Text   string `json:"text"`
	Time   int64  `json:"time"`
}

func (m *Message) String() string {
	time := time.UnixMilli(m.Time)
	return fmt.Sprintf("[%v:%v]%v: %v", time.Hour(), time.Minute(), m.Sender, m.Text)
}
