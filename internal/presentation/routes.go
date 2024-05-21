package handlers

import (
	"context"
	"hackathon/exceptions"
	"hackathon/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *HTTPhandler) BindRoutesAndMiddlewares() {
	h.app.Use("/ws", func(c *fiber.Ctx) error {
		userGUID := c.Params("GUID")
		chatroomId := c.Params("chatroomID")

		if len(userGUID) == 0 || len(chatroomId) == 0 {
			return c.Status(http.StatusBadRequest).JSON(Response{
				Error:   exceptions.ErrBadRequest.Error(),
				Content: nil,
			})
		}

		userData := models.User{
			GUID: userGUID,
		}

		chatroomData := models.Chatroom{
			ChatroomId: chatroomId,
		}

		ctxUser, cancel := context.WithTimeout(context.TODO(), time.Second*5)
		defer cancel()

		ctxChatroom, cancel := context.WithTimeout(context.TODO(), time.Second*5)
		defer cancel()

		if chatroom := h.WsBusiness.GetChatroom(ctxChatroom, chatroomData); chatroom == nil {
			return c.Status(http.StatusNotFound).JSON(Response{
				Error:   exceptions.ErrNotFound.Error(),
				Content: nil,
			})
		}

		if user := h.WsBusiness.GetUser(ctxUser, userData); user == nil {
			return c.Status(http.StatusNotFound).JSON(Response{
				Error:   exceptions.ErrNotFound.Error(),
				Content: nil,
			})
		}

		c.Next()
		return nil
	})
}
