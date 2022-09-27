package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/xegcrbq/P2PChat/internal/api/controller"
	"github.com/xegcrbq/P2PChat/internal/db"
)

type APIService struct {
	log    *logrus.Entry
	router *fiber.App
}

func (svc *APIService) Serve(addr string) {
	svc.log.Fatal(svc.router.Listen(addr))
}

func (svc *APIService) Shutdown(ctx context.Context) error {
	return svc.router.Shutdown()
}

func NewAPIService(log *logrus.Entry, dbConn *pgxpool.Pool) (*APIService, error) {
	engine := html.New("./templates", ".html")

	svc := &APIService{
		log: log,
		router: fiber.New(fiber.Config{
			Views: engine,
		}),
	}
	repository, err := db.NewRepository(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	controllersRegistry := controller.NewRegistry(log, repository)
	api := svc.router.Group("/api/")
	api.Static("/", "./templates")
	api.Get("chat/", controllersRegistry.AuthController.UserChat)
	api.Post("uploadFile/", controllersRegistry.ChatController.SendFile)
	api.Get("login/:username", controllersRegistry.AuthController.UsernameEntered)
	//talkMe
	api.Post("send/", controllersRegistry.ChatController.SendMessageToTalkMe)
	api.Post("update/v2/", controllersRegistry.ChatController.UpdateV2)
	api.Post("webhook/", controllersRegistry.ChatController.WH)

	return svc, nil
}
