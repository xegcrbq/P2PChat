package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xegcrbq/P2PChat/internal/models"
	"github.com/xegcrbq/P2PChat/internal/models/commands"
	"github.com/xegcrbq/P2PChat/internal/utils"
	"net/http"
	"time"
)

type AuthController struct {
	tknz           *utils.Tokenizer
	dataController *DataController
}

func NewAuthController(tknz *utils.Tokenizer, dataController *DataController) *AuthController {
	return &AuthController{
		tknz:           tknz,
		dataController: dataController,
	}
}

func (cC *AuthController) newSession(username string) *fiber.Cookie {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	return cC.tknz.NewJWTCookie("session_id", username, expirationTime)
}

func (cC *AuthController) UserChat(c *fiber.Ctx) error {
	sessionId := c.Cookies("session_id")
	if sessionId == "" {
		err := c.SendStatus(http.StatusUnauthorized)
		if err != nil {
			return err
		}
		return nil
	}
	_, token, err := cC.tknz.ParseDataClaims(sessionId)
	if err != nil || !token.Valid {
		err := c.SendStatus(http.StatusUnauthorized)
		if err != nil {
			return err
		}
		return nil
	}
	return c.Render("chat", fiber.Map{})
}

func (cC *AuthController) UsernameEntered(c *fiber.Ctx) error {

	username := c.Params("username")
	fmt.Println(username)
	answ := cC.dataController.Execute(commands.CreateUserByUser{User: &models.User{
		UserName: username,
		Password: username,
	}})
	fmt.Println(answ.Err)
	if answ.Err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return nil
	}
	unCookie := cC.newSession(answ.UserName)
	c.Cookie(unCookie)
	return nil
}
