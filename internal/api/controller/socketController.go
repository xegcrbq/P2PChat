package controller

import (
	"encoding/json"
	"fmt"
	"github.com/antoniodipinto/ikisocket"
	"github.com/xegcrbq/P2PChat/internal/db"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
	"github.com/xegcrbq/P2PChat/internal/services"
	"github.com/xegcrbq/P2PChat/internal/utils"
	"net/http"
	"strings"
	"time"
)

type SocketController struct {
	tknz           *utils.Tokenizer
	socketService  *services.SocketService
	dataController *DataController
}

func NewSocketController(tknz *utils.Tokenizer, socketService *services.SocketService, dataController *DataController) *SocketController {
	c := &SocketController{
		tknz:           tknz,
		socketService:  socketService,
		dataController: dataController,
	}
	c.Init()
	return c
}
func (c *SocketController) Init() {
	// Multiple event handling supported
	ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
		c.socketService.AddSocket(ep.SocketAttributes["username"].(string), ep.SocketUUID, ep.Kws)
		ep.Kws.Emit(c.loadPreviousMessages(ep.SocketAttributes["username"].(string)))
	})
	//old
	//ikisocket.On(ikisocket.EventMessage, func(ep *ikisocket.EventPayload) {
	//	fmt.Println("message: " + string(ep.Data) + " from " + ep.SocketAttributes["username"].(string))
	//	ep.Kws.Emit(ep.Data)
	//	sockets, _ := c.socketService.Get(ep.SocketAttributes["username"].(string))
	//	for k, v := range sockets {
	//		fmt.Println(fmt.Sprintf("User: %v\nUUID%v", k, v))
	//	}
	//})

	ikisocket.On(ikisocket.EventMessage, c.sendMessageToTalkMe)

	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		c.socketService.DeleteSocket(ep.SocketAttributes["username"].(string), ep.SocketUUID)
	})
}
func (c *SocketController) SocketReaderCreate(kws *ikisocket.Websocket) {
	data, tkn, _ := c.tknz.ParseDataClaims(kws.Params("session_id"))
	if !tkn.Valid {
		return
	}
	kws.SetAttribute("username", data.Data)
	return
}
func (c *SocketController) sendMessageToTalkMe(ep *ikisocket.EventPayload) {
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
	`, ep.SocketAttributes["username"].(string), ep.SocketAttributes["username"].(string), string(ep.Data), time.Now().Unix()))
	req, _ := http.NewRequest("POST", URL, json)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Token", db.GoDotEnvVariable("XToken"))
	client := &http.Client{}
	client.Do(req)
}
func (c *SocketController) loadPreviousMessages(username string) []byte {
	var userId int32
	{
		answU := c.dataController.Execute(commands.ReadUserByUserName{UserName: username})
		if answU.Err != nil || answU.User == nil {
			return nil
		}
		userId = answU.User.UserId
	}
	answM := c.dataController.Execute(commands.ReadMessagesByUserId{UserId: userId})
	if answM.Err != nil || answM.Messages == nil {
		return nil
	}
	messagesToSend := *answM.Messages
	for i := range messagesToSend {
		if userId == messagesToSend[i].SenderId {
			messagesToSend[i].SenderId = 0
		} else {
			messagesToSend[i].SenderId = 1
		}
	}
	sendData, err := json.Marshal(messagesToSend)
	if err != nil {
		return nil
	}
	return sendData

}
