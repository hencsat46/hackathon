package handlers

import (
	"context"
	"errors"
	"fmt"
	"hackathon/models"
	"hackathon/pkg/exceptions"
	"log/slog"
	"net/http"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (h *HTTPhandler) createChatroom(c *fiber.Ctx) error {
	var request Chatroom

	if err := c.BodyParser(&request); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("create chatroom endpoint called: %v\n", request))

	chatroomData := models.Chatroom{
		Name:      request.Name,
		OwnerGUID: request.OwnerGUID,
		IsPrivate: request.IsPrivate,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	returnData, err := h.ChatroomBusiness.CreateChatroom(ctx, chatroomData)
	if err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Error:   exceptions.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	if _, ok := h.hub[returnData.ChatroomId]; !ok {
		h.hub[returnData.ChatroomId] = &models.Room{
			CID:          returnData.ChatroomId,
			Participants: make(map[string]*websocket.Conn),
		}
	}

	slog.Debug(fmt.Sprintf("%v", h.hub))

	return c.Status(http.StatusOK).JSON(Response{
		Error:   "nil",
		Content: returnData.ChatroomId,
	})
}

func (h *HTTPhandler) updateChatroom(c *fiber.Ctx) error {
	var request Chatroom

	if err := c.BodyParser(&request); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("update chatroom endpoint called: %v\n", request))

	chatroomData := models.Chatroom{
		ChatroomId: request.ID,
		Name:       request.Name,
		OwnerGUID:  request.OwnerGUID,
		IsPrivate:  request.IsPrivate,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.ChatroomBusiness.UpdateChatroom(ctx, chatroomData); err != nil {
		slog.Debug(err.Error())
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
	slog.Debug(fmt.Sprintf("delete chatroom endpoint called: %v\n", request))

	chatroomData := models.Chatroom{
		ChatroomId: request.ID,
		OwnerGUID:  request.OwnerGUID,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.ChatroomBusiness.DeleteChatroom(ctx, chatroomData); err != nil {
		slog.Debug(err.Error())
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
