package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xegcrbq/P2PChat/internal/utils"
	"net/http"
	"time"
)

type AuthController struct {
	tknz *utils.Tokenizer
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
	return c.Render("chatWindow2", fiber.Map{})
}

func (cC *AuthController) UsernameEntered(c *fiber.Ctx) error {
	username := c.Params("username")
	unCookie := cC.newSession(username)
	c.Cookie(unCookie)
	return nil
}

func NewAuthController(tknz *utils.Tokenizer) *AuthController {
	return &AuthController{
		tknz: tknz,
	}
}
