package handlers

import (
	"context"
	"hackathon/models"
	"hackathon/pkg/exceptions"
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

		if _, err := h.WsBusiness.GetChatroom(ctxChatroom, chatroomData); err != nil {
			return c.Status(http.StatusNotFound).JSON(Response{
				Error:   exceptions.ErrNotFound.Error(),
				Content: nil,
			})
		}

		if _, err := h.WsBusiness.GetUser(ctxUser, userData); err != nil {
			return c.Status(http.StatusNotFound).JSON(Response{
				Error:   exceptions.ErrNotFound.Error(),
				Content: nil,
			})
		}

		c.Next()
		return nil
	})

	userRoutes := h.app.Group("/user")
	chatroomRoutes := h.app.Group("/chatroom")
	messageRoutes := h.app.Group("/message")
	wsRoutes := h.app.Group("/ws")

	userRoutes.Post("/create", h.createUser)
}
