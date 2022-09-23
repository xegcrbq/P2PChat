package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/thanhpk/randstr"
	"github.com/xegcrbq/P2PChat/controller"
	"github.com/xegcrbq/P2PChat/tokenizer"
	"log"
)

func main() {
	engine := html.New("./templates", ".html")
	cC := controller.NewChatController(tokenizer.NewTokenizer([]byte(randstr.Hex(10))))
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./templates")
	app.Get("from/:you/to/:anotherUser", cC.ChatWindow)
	app.Get("chat/", cC.UserChat)
	app.Post("send-message/", cC.Send)
	app.Post("update/", cC.Update)
	app.Post("uploadFile/", cC.SendFile)
	app.Get("login/:username", cC.UsernameEntered)
	log.Fatal(app.Listen(":1080"))
}
