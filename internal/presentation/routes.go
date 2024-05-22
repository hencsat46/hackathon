package handlers

import (
	"context"
	"net/http"
	"time"

	e "hackathon/exceptions"
	"hackathon/internal/presentation/entities"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (h *HTTPhandler) bindRoutesAndMiddlewares() {
	h.app.Use("/ws", func(c *fiber.Ctx) error {
		GUID := c.Params("GUID")
		chatroomID := c.Params("chatroomID")

		if len(GUID) == 0 || len(chatroomID) == 0 {
			return c.Status(http.StatusBadRequest).JSON(entities.Response{
				Error:   e.ErrBadRequest.Error(),
				Content: nil,
			})
		}

		ctxUser, cancel := context.WithTimeout(context.TODO(), time.Second*5)
		defer cancel()

		ctxChatroom, cancel := context.WithTimeout(context.TODO(), time.Second*5)
		defer cancel()

		if _, err := h.WsBusiness.GetChatroom(ctxChatroom, chatroomID); err != nil {
			return c.Status(http.StatusNotFound).JSON(entities.Response{
				Error:   e.ErrNotFound.Error(),
				Content: nil,
			})
		}

		if _, err := h.WsBusiness.GetUser(ctxUser, GUID); err != nil {
			return c.Status(http.StatusNotFound).JSON(entities.Response{
				Error:   e.ErrNotFound.Error(),
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

	userRoutes.Post("/create", h.CreateUser)
	userRoutes.Post("/login", h.Login)
	userRoutes.Put("/updateUsername", h.jwtMiddleware.ValidateToken(h.UpdateChatroom))
	userRoutes.Put("/updateEmail", h.jwtMiddleware.ValidateToken(h.UpdateEmail))
	userRoutes.Put("/updatePassword", h.jwtMiddleware.ValidateToken(h.UpdatePassword))
	userRoutes.Delete("/delete", h.jwtMiddleware.ValidateToken(h.DeleteUser))
	userRoutes.Get("/userChatrooms/:guid", h.jwtMiddleware.ValidateToken(h.FetchUserChatrooms))

	wsRoutes.Get("/:GUID/:cid", h.jwtMiddleware.ValidateToken(websocket.New(h.HandleWS)))

	chatroomRoutes.Post("/create", h.jwtMiddleware.ValidateToken(h.CreateChatroom))
	chatroomRoutes.Put("/", h.jwtMiddleware.ValidateToken(h.UpdateChatroom))
	chatroomRoutes.Delete("/:chatroomID", h.jwtMiddleware.ValidateToken(h.DeleteChatroom))

	messageRoutes.Get("/:cid", h.jwtMiddleware.ValidateToken(h.FetchMessagesForChatroom))
	messageRoutes.Put("/", h.jwtMiddleware.ValidateToken(h.UpdateMessage))
	messageRoutes.Delete("/", h.jwtMiddleware.ValidateToken(h.DeleteMessage))
}
