package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xegcrbq/P2PChat/db"
	"github.com/xegcrbq/P2PChat/models"
	"github.com/xegcrbq/P2PChat/tokenizer"
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
	dialogues     []*models.Dialogue
	tknz          *tokenizer.Tokenizer
	dialoguePairs map[string]string
}

func NewChatController(tknz *tokenizer.Tokenizer) *ChatController {
	return &ChatController{
		tknz:          tknz,
		dialoguePairs: make(map[string]string),
	}
}
func (cC *ChatController) GetDialogue(users []string) *models.Dialogue {
	if len(users) != 2 {
		return nil
	}
	sort.Strings(users)
	dialogue := cC.FindDialogue(users)

	if dialogue == nil {
		dialogue = &models.Dialogue{
			User1: users[0],
			User2: users[1],
		}
		cC.dialogues = append(cC.dialogues, dialogue)
	}
	return dialogue
}
func (cC *ChatController) FindDialogue(users []string) *models.Dialogue {
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
func (cC *ChatController) AddMessage(message *models.OldMessage) {
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
	resp := &models.OldMessage{}
	err := json.Unmarshal(c.Body(), &resp)
	if err != nil {
		fmt.Println(err)
	}
	cC.AddMessage(resp)
	//fmt.Println(cC.GetDialogue([]string{resp.Sender, resp.Target}).String())
	return nil
}
func (cC *ChatController) Update(c *fiber.Ctx) error {
	resp := &models.OldMessage{}
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

func (cC *ChatController) UserChat(c *fiber.Ctx) error {
	sessionId := c.Cookies("session_id")
	if sessionId == "" {
		c.SendStatus(http.StatusUnauthorized)
		return nil
	}
	_, token, err := cC.tknz.ParseDataClaims(sessionId)
	if err != nil || !token.Valid {
		c.SendStatus(http.StatusUnauthorized)
		return nil
	}

	return c.Render("chatWindow2", fiber.Map{})
}
func (cC *ChatController) newSession(username string) *fiber.Cookie {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	cC.dialoguePairs[username] = db.GetRandomName()
	return cC.tknz.NewJWTCookie("session_id", username, expirationTime)
}
func (cC *ChatController) UsernameEntered(c *fiber.Ctx) error {
	username := c.Params("username")
	unCookie := cC.newSession(username)
	c.Cookie(unCookie)
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
	return nil
}
