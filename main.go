package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/xegcrbq/P2PChat/controller"
	"log"
)

func main() {
	engine := html.New("./templates", ".html")
	cC := &controller.ChatController{}
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("from/:you/to/:anotherUser", cC.ChatWindow)
	app.Post("send-message/", cC.Send)
	app.Post("update/", cC.Update)

	log.Fatal(app.Listen(":1080"))
}
