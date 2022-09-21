package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xegcrbq/P2PChat/models"
	"sort"
)

type Response struct {
	MessageHistory string
}
type ChatController struct {
	dialogues []*models.Dialogue
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
func (cC *ChatController) AddMessage(message *models.Message) {
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
	resp := &models.Message{}
	err := json.Unmarshal(c.Body(), &resp)
	if err != nil {
		fmt.Println(err)
	}
	cC.AddMessage(resp)
	//fmt.Println(cC.GetDialogue([]string{resp.Sender, resp.Target}).String())
	return nil
}
func (cC *ChatController) Update(c *fiber.Ctx) error {
	resp := &models.Message{}
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
