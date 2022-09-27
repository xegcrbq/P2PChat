package controller

import (
	"encoding/json"
	"fmt"
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type TalkmeController struct {
	xToken         string
	dataController *DataController
	dateStart      time.Time
}

func NewTalkmeController(xToken string, controller *DataController) *TalkmeController {
	return &TalkmeController{
		xToken:         xToken,
		dataController: controller,
		dateStart:      time.Now().Add(-time.Hour * 24 * 180),
	}
}

// AutoUpdate в начале обновляет бд до новейшего состояния, затем каждые duration отправляет на talkme запрос на получение последних сообщений и вносит их в бд
func (c *TalkmeController) Update(duration time.Duration, infinity bool) {
	answA := c.dataController.Execute(commands.ReadUserByUserName{UserName: "admin"})
	if answA.Err != nil || answA.User == nil {
		return
	}
	answM := c.dataController.Execute(commands.MessageReadNewest{})
	if answM.Err == nil && answM.Messages != nil {
		message := *answM.Messages
		c.dateStart = message[0].SendTime
	}
	if duration < time.Second*3 {
		duration = time.Second * 10
	}
	for c.dateStart.Add(duration).Unix() < time.Now().Unix() {
		dateEnd := time.Now()
		if c.dateStart.Add(time.Hour*24*7).Unix() < dateEnd.Unix() {
			dateEnd = c.dateStart.Add(time.Hour * 24 * 7)
		}
		c.readAndUpdateDB(dateEnd, answA.User.UserId)
		c.dateStart = dateEnd
	}
	if infinity {
		for range time.Tick(duration) {
			dateEnd := time.Now()
			if c.dateStart.Add(time.Hour*24*7).Unix() < dateEnd.Unix() {
				dateEnd = c.dateStart.Add(time.Hour * 24 * 7)
			}
			c.readAndUpdateDB(dateEnd, answA.User.UserId)
			c.dateStart = dateEnd
		}
	}
}

// readMessagesForPeriod отправляет на talkme запрос на получение последних сообщений
func (c *TalkmeController) readMessagesForPeriod(start, end time.Time) (*models.TalkMeMessageGetListAnswer, error) {
	URL := "https://lcab.talk-me.ru/json/v1.0/chat/message/getList"
	startString := fmt.Sprintf("%.4v-%.2v-%.2v %.2v:%.2v:%.2v", start.Year(), int(start.Month()), start.Day(), start.Hour(), start.Minute(), start.Second())
	endString := fmt.Sprintf("%.4v-%.2v-%.2v %.2v:%.2v:%.2v", end.Year(), int(end.Month()), end.Day(), end.Hour(), end.Minute(), end.Second())
	bodyJson := strings.NewReader(fmt.Sprintf(`
	{
		"dateRange": {
			"start": "%v",
			"stop": "%v"
		}
	}
	`, startString, endString))
	req, _ := http.NewRequest("POST", URL, bodyJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Token", c.xToken)
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	talkMeAnswer := models.TalkMeMessageGetListAnswer{}
	err = json.Unmarshal(bodyBytes, &talkMeAnswer)
	if err != nil {
		return nil, err
	}
	return &talkMeAnswer, nil
}

// messagesFromOperator отфильтровывает сообщения, оставляет только те, отправители которых есть в нашей бд
func (c *TalkmeController) messagesFromOperator(messageList *models.TalkMeMessageGetListAnswer) *[]models.TalkMeMessageGetListResult {
	var result []models.TalkMeMessageGetListResult
	for _, inputResult := range messageList.Result {
		answ := c.dataController.Execute(commands.ReadUserByUserName{UserName: inputResult.ClientId})
		if answ.Err != nil || answ.User == nil {
			continue
		}
		var messages []models.TalkMeMessage
		for _, message := range inputResult.Messages {
			if message.MessageType != "comment" {
				//message.WhoSend == "operator" &&
				messages = append(messages, message)
			}
		}
		result = append(result, models.TalkMeMessageGetListResult{
			Messages: messages,
			ClientId: inputResult.ClientId,
		})
	}
	return &result
}

// readAndUpdateDB запрос на talkme + внесение в бд
func (c *TalkmeController) readAndUpdateDB(dateEnd time.Time, adminId int32) {
	answ, err := c.readMessagesForPeriod(c.dateStart, dateEnd)
	if err != nil {
		log.Fatal(err)
		return
	}
	messagesByClients := c.messagesFromOperator(answ)
	for _, client := range *messagesByClients {
		answU := c.dataController.Execute(commands.ReadUserByUserName{UserName: client.ClientId})
		if answU.Err != nil || answU.User == nil {
			log.Fatal("AutoUpdate user not found")
			continue
		}
		for _, message := range client.Messages {
			c.WriteMessageFromTmeMessage(&message, answU.User.UserId, adminId)
		}
	}
}
func (c *TalkmeController) WriteMessageFromTmeMessage(tmeM *models.TalkMeMessage, userId, adminId int32) *models.Answer {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", tmeM.DateTime)
	if err != nil {
		log.Fatal(err)
	}
	if tmeM.WhoSend == "operator" {
		return c.dataController.Execute(commands.CreateMessagesByMessage{Message: &models.Message{
			SenderId:     adminId,
			ReaderId:     userId,
			OrderId:      1,
			MessageText:  tmeM.Text,
			AttachmentId: 0,
			SendTime:     parsedTime,
			TalkMeId:     tmeM.Id,
		}})
	} else {
		return c.dataController.Execute(commands.CreateMessagesByMessage{Message: &models.Message{
			SenderId:     userId,
			ReaderId:     adminId,
			OrderId:      1,
			MessageText:  tmeM.Text,
			AttachmentId: 0,
			SendTime:     parsedTime,
			TalkMeId:     tmeM.Id,
		}})
	}
}
func (c *TalkmeController) MessageFromWHBytes(data []byte) error {
	var twh models.TalkMeWebHook
	err := json.Unmarshal(data, &twh)
	if err != nil {
		return err
	}
	if !twh.Validate() {
		return models.ErrInvalidSC
	}
	answA := c.dataController.Execute(commands.ReadUserByUserName{UserName: "admin"})
	if answA.Err != nil {
		return answA.Err
	}
	answU := c.dataController.Execute(commands.ReadUserByUserName{UserName: twh.Data.Client.ClientId})
	return c.WriteMessageFromTmeMessage(&twh.Data.Message, answU.User.UserId, answA.User.UserId).Err
}
