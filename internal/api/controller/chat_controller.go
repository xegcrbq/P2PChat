package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	models2 "github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
	"github.com/xegcrbq/P2PChat/internal/utils"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Response struct {
	MessageHistory string
}
type ChatController struct {
	dialogues      []*models2.Dialogue
	tknz           *utils.Tokenizer
	dialoguePairs  map[string]string
	dataController *DataController
}

func NewChatController(tknz *utils.Tokenizer, dataController *DataController) *ChatController {
	return &ChatController{
		tknz:           tknz,
		dialoguePairs:  make(map[string]string),
		dataController: dataController,
	}
}

func (cC *ChatController) GetDialogue(users []string) *models2.Dialogue {
	if len(users) != 2 {
		return nil
	}
	sort.Strings(users)
	dialogue := cC.FindDialogue(users)

	if dialogue == nil {
		dialogue = &models2.Dialogue{
			User1: users[0],
			User2: users[1],
		}
		cC.dialogues = append(cC.dialogues, dialogue)
	}
	return dialogue
}
func (cC *ChatController) FindDialogue(users []string) *models2.Dialogue {
	if len(users) != 2 {
		return nil
	}
	sort.Strings(users)
	for i, _ := range cC.dialogues {
		if users[0] == cC.dialogues[i].User1 && users[1] == cC.dialogues[i].User2 {
			return cC.dialogues[i]
		}
	}
	return nil
}
func (cC *ChatController) AddMessage(message *models2.OldMessage) {
	users := []string{message.Sender, message.Target}
	sort.Strings(users)
	dialogue := cC.GetDialogue(users)
	dialogue.Messages = append(dialogue.Messages, message)
}
func (cC *ChatController) ChatWindow(c *fiber.Ctx) error {
	data := []string{c.Params("anotherUser"), c.Params("you")}
	sort.Strings(data)
	dialogue := cC.GetDialogue(data)
	return c.Render("chatWindow", fiber.Map{
		"messages": dialogue.String(),
	})
}
func (cC *ChatController) Send(c *fiber.Ctx) error {
	resp := &models2.OldMessage{}
	err := json.Unmarshal(c.Body(), &resp)
	if err != nil {
		fmt.Println(err)
	}
	cC.AddMessage(resp)
	//fmt.Println(cC.GetDialogue([]string{resp.Sender, resp.Target}).String())
	return nil
}
func (cC *ChatController) Update(c *fiber.Ctx) error {
	resp := &models2.OldMessage{}
	err := json.Unmarshal(c.Body(), &resp)
	if err != nil {
		fmt.Println(err)
	}
	users := []string{resp.Target, resp.Sender}
	sort.Strings(users)
	dialogue := cC.GetDialogue(users)
	answ := &Response{MessageHistory: dialogue.String()}
	data, _ := json.Marshal(answ)
	return c.Send(data)
}
func (cC *ChatController) Test(c *fiber.Ctx) error {
	fmt.Println(c.Body())
	return nil
}

func (cC *ChatController) SendFile(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()))
	return nil
}
func (cC *ChatController) SendMessageToTalkMe(c *fiber.Ctx) error {
	data, tkn, err := cC.tknz.ParseDataClaims(c.Cookies("session_id"))
	if !(tkn.Valid) || err != nil {
		c.SendStatus(http.StatusUnauthorized)
		return nil
	}
	//fmt.Println(data.Data)
	message := struct {
		Message string
	}{}
	err = json.Unmarshal(c.Body(), &message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(message.Message)
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
	req.Header.Set("X-Token", "xuw9xn7znrz4658f862quecb1p8n1s32vhpo35m61yzrofjepnqk0i2tlum3vhqr")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)
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
func (cC *ChatController) UpdateV2(c *fiber.Ctx) error {
	data, tkn, err := cC.tknz.ParseDataClaims(c.Cookies("session_id"))
	if !(tkn.Valid) || err != nil {
		c.SendStatus(http.StatusUnauthorized)
		return nil
	}
	messageCount := struct {
		MessageCount int
	}{}
	err = json.Unmarshal(c.Body(), &messageCount)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	answ := cC.dataController.Execute(commands.ReadUserByUserName{UserName: data.Data})
	if answ.Err != nil || answ.User == nil {
		c.SendStatus(http.StatusUnauthorized)
		return nil
	}
	userId := answ.User.UserId
	answ = cC.dataController.Execute(commands.ReadMessagesByUserId{UserId: userId})
	if answ.Err != nil || answ.Messages == nil {
		c.SendStatus(http.StatusUnauthorized)
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
	return c.Send(sendData)
}
func (cC *ChatController) WH(c *fiber.Ctx) error {
	fmt.Println(c.String())
	return nil
}
