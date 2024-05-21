package handlers

import (
	"context"
	"errors"
	"hackathon/models"
	"net/http"
	"time"

	"hackathon/exceptions"

	"github.com/gofiber/fiber/v2"
)

type IChallenChatroom interface {
	FetchUserChatrooms(ctx context.Context, chatroomData models.Chatroom) ([]models.Chatroom, error)
	CreateChatroom(ctx context.Context, chatroomData models.Chatroom) error
	UpdateChatroom(ctx context.Context, chatroomData models.Chatroom) error
	DeleteChatroom(ctx context.Context, chatroomData models.Chatroom) error
}

func (h *HTTPhandler) fetchUserChatrooms(c *fiber.Ctx) error {
	userGUID := c.Params("guid")

	if len(userGUID) == 0 {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}

	chatroomData := models.Chatroom{
		OwnerGUID: userGUID,
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	chatrooms, err := h.ChatroomChallen.FetchUserChatrooms(ctx, chatroomData)
	if err != nil {
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

	return c.Status(http.StatusOK).JSON(Response{
		Error:   "nil",
		Content: chatrooms,
	})
}

func (h *HTTPhandler) createChatroom(c *fiber.Ctx) error {
	var request Chatroom

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}

	chatroomData := models.Chatroom{
		ChatroomId:        request.ID,
		Name:              request.Name,
		OwnerGUID:         request.OwnerGUID,
		IsPrivate:         request.IsPrivate,
		ParticipantsLimit: request.ParticipantsLimit,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.ChatroomChallen.CreateChatroom(ctx, chatroomData); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Error:   exceptions.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusOK).JSON(Response{
		Error:   "nil",
		Content: "Chatroom created",
	})
}

func (h *HTTPhandler) updateChatroom(c *fiber.Ctx) error {
	var request Chatroom

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}

	chatroomData := models.Chatroom{
		ChatroomId:        request.ID,
		Name:              request.Name,
		OwnerGUID:         request.OwnerGUID,
		IsPrivate:         request.IsPrivate,
		ParticipantsLimit: request.ParticipantsLimit,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.ChatroomChallen.UpdateChatroom(ctx, chatroomData); err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			return c.Status(http.StatusNotFound).JSON(Response{
				Error:   exceptions.ErrNotFound.Error(),
				Content: "Chatroom not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Error:   exceptions.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusOK).JSON(Response{
		Error:   "nil",
		Content: "Chatroom updated",
	})
}

func (h *HTTPhandler) deleteChatroom(c *fiber.Ctx) error {
	var request Chatroom

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}

	chatroomData := models.Chatroom{
		ChatroomId: request.ID,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.ChatroomChallen.DeleteChatroom(ctx, chatroomData); err != nil {
		if errors.Is(err, exceptions.ErrNotFound) {
			return c.Status(http.StatusNotFound).JSON(Response{
				Error:   exceptions.ErrNotFound.Error(),
				Content: "Chatroom not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Error:   exceptions.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusOK).JSON(Response{
		Error:   "nil",
		Content: "Chatroom deleted",
	})
}
