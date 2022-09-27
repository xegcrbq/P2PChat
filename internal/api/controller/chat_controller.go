package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xegcrbq/P2PChat/internal/db"
	models2 "github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
	"github.com/xegcrbq/P2PChat/internal/utils"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	MessageHistory string
}
type ChatController struct {
	dialogues        []*models2.Dialogue
	tknz             *utils.Tokenizer
	dialoguePairs    map[string]string
	dataController   *DataController
	talkmeController *TalkmeController
}

func NewChatController(tknz *utils.Tokenizer, dataController *DataController, talkmeController *TalkmeController) *ChatController {
	return &ChatController{
		tknz:             tknz,
		dialoguePairs:    make(map[string]string),
		dataController:   dataController,
		talkmeController: talkmeController,
	}
}

func (c *ChatController) SendFile(ctx *fiber.Ctx) error {
	fmt.Println(string(ctx.Body()))
	return nil
}
func (c *ChatController) SendMessageToTalkMe(ctx *fiber.Ctx) error {
	data, tkn, err := c.tknz.ParseDataClaims(ctx.Cookies("session_id"))
	if !(tkn.Valid) || err != nil {
		ctx.SendStatus(http.StatusUnauthorized)
		return nil
	}
	message := struct {
		Message string
	}{}
	err = json.Unmarshal(ctx.Body(), &message)
	if err != nil {
		return err
	}
	URL := "https://lcab.talk-me.ru/json/v1.0/chat/message/sendToOperator"
	json := strings.NewReader(fmt.Sprintf(`
	{
		"client": {
			"id": "%v",
			"name": "[defi]%v",
			"phone": "+7-900-000-0000",
			"email": "ivan@gmail.com"
		},
		"message": {
			"text": "%v",
			"tag": "%v",
			"attachments": []
		}
	}
	`, data.Data, data.Data, message.Message, time.Now().Unix()))
	req, _ := http.NewRequest("POST", URL, json)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Token", db.GoDotEnvVariable("XToken"))
	client := &http.Client{}
	client.Do(req)
	//resp, _ := client.Do(req)
	//defer resp.Body.Close()
	//bodyBytes, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//bodyString := string(bodyBytes)
	//log.Println(bodyString)
	//writing message in db
	//answU := cC.dataController.Execute(commands.ReadUserByUserName{UserName: data.Data})
	//if answU.Err != nil || answU.User == nil {
	//	return answU.Err
	//}
	//answA := cC.dataController.Execute(commands.ReadUserByUserName{UserName: "admin"})
	//if answA.Err != nil || answA.User == nil {
	//	return answA.Err
	//}
	//answM := cC.dataController.Execute(commands.CreateMessagesByMessage{Message: &models2.Message{
	//	SenderId:     answU.User.UserId,
	//	ReaderId:     answA.User.UserId,
	//	OrderId:      1,
	//	MessageText:  message.Message,
	//	AttachmentId: 0,
	//	SendTime:     time.Now(),
	//}})
	//if answM.Err != nil {
	//	return answM.Err
	//}
	return nil
}

func (c *ChatController) UpdateV2(ctx *fiber.Ctx) error {
	data, tkn, err := c.tknz.ParseDataClaims(ctx.Cookies("session_id"))
	if !(tkn.Valid) || err != nil {
		ctx.SendStatus(http.StatusUnauthorized)
		return nil
	}
	messageCount := struct {
		MessageCount int
	}{}
	err = json.Unmarshal(ctx.Body(), &messageCount)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil
	}
	answ := c.dataController.Execute(commands.ReadUserByUserName{UserName: data.Data})
	if answ.Err != nil || answ.User == nil {
		ctx.SendStatus(http.StatusUnauthorized)
		return nil
	}
	userId := answ.User.UserId
	answ = c.dataController.Execute(commands.ReadMessagesByUserId{UserId: userId})
	if answ.Err != nil || answ.Messages == nil {
		ctx.SendStatus(http.StatusUnauthorized)
		return nil
	}
	if len(*answ.Messages) <= messageCount.MessageCount || messageCount.MessageCount < 0 {
		return nil
	}
	messagesToSend := *answ.Messages
	messagesToSend = messagesToSend[messageCount.MessageCount:]
	for i := range messagesToSend {
		if userId == messagesToSend[i].SenderId {
			messagesToSend[i].SenderId = 0
		} else {
			messagesToSend[i].SenderId = 1
		}

	}
	sendData, err := json.Marshal(messagesToSend)
	if err != nil {
		return err
	}
	return ctx.Send(sendData)
}
func (c *ChatController) WH(ctx *fiber.Ctx) error {
	return c.talkmeController.MessageFromWHBytes(ctx.Body())
}
