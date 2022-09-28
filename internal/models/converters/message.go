package converters

import (
	"github.com/xegcrbq/P2PChat/internal/models"
	"time"
)

// TalkMeMessageToMessage use ONLY for socket
func TalkMeMessageToMessage(tm *models.TalkMeMessage) (m *models.Message) {
	if tm == nil {
		m = &models.Message{}
		return
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", tm.DateTime)
	if err != nil {
		return &models.Message{}
	}
	if tm.WhoSend == "operator" {
		m = &models.Message{
			SenderId:    1,
			ReaderId:    0,
			OrderId:     1,
			MessageText: tm.Text,
			SendTime:    parsedTime,
			TalkMeId:    tm.Id,
		}
	} else {
		m = &models.Message{
			SenderId:    0,
			ReaderId:    1,
			OrderId:     1,
			MessageText: tm.Text,
			SendTime:    parsedTime,
			TalkMeId:    tm.Id,
		}
	}
	return
}

// TalkMeMessagesToMessages use ONLY for socket
func TalkMeMessagesToMessages(tms []models.TalkMeMessage) (ms []models.Message) {
	ms = []models.Message{}
	for i := range tms {
		ms = append(ms, *TalkMeMessageToMessage(&tms[i]))
	}
	return
}
