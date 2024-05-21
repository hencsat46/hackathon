package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"hackathon/models"

	"hackathon/exceptions"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type IChallenWS interface {
	CheckUser(ctx context.Context, userData models.User) error
	CheckChatroom(ctx context.Context, chatroomData models.Chatroom) error
	GetUser(ctx context.Context, userData models.User) models.User
	GetChatroom(ctx context.Context, chatroomData models.Chatroom) models.Chatroom
}

func (h *HTTPhandler) CreateChatroomRoutes() {
	h.app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	h.app.Use("/ws/:chatroomID/:GUID", func(c *fiber.Ctx) error {
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

		if err := h.WsChallen.CheckChatroom(ctxChatroom, chatroomData); err != nil {
			if errors.Is(err, exceptions.ErrNotFound) {
				return c.Status(http.StatusNotFound).JSON(Response{
					Error:   exceptions.ErrNotFound.Error(),
					Content: nil,
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(Response{
				Error:   exceptions.ErrInternalServerError.Error(),
				Content: nil,
			})
		}

		if err := h.WsChallen.CheckUser(ctxUser, userData); err != nil {
			if errors.Is(err, exceptions.ErrNotFound) {
				return c.Status(http.StatusNotFound).JSON(Response{
					Error:   exceptions.ErrNotFound.Error(),
					Content: nil,
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(Response{
				Error:   exceptions.ErrInternalServerError.Error(),
				Content: nil,
			})
		}

		c.Next()
		return nil

	})
}

// func (h *HTTPhandler) handleWs(conn *websocket.Conn) {
// 	userGUID := conn.Params("GUID")
// 	chatroomId := conn.Params("chatroomID")

// 	userData := models.User{GUID: userGUID}
// 	chatroomData := models.Chatroom{ChatroomId: chatroomId}

// 	ctxUser, cancel := context.WithTimeout(context.TODO(), time.Second*5)
// 	defer cancel()

// 	ctxChatroom, cancel := context.WithTimeout(context.TODO(), time.Second*5)
// 	defer cancel()

// 	user := h.WsChallen.GetUser(ctxUser, userData)
// 	chatroom := h.WsChallen.GetChatroom(ctxChatroom, chatroomData)
// 	user.Conn = c
// 	ctrl.listenUsersMessages(user, chatroom)
// }
