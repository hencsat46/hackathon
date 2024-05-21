package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"hackathon/exceptions"
	"hackathon/models"

	"github.com/gofiber/fiber/v2"
)

func (h *HTTPhandler) updateMessage(c *fiber.Ctx) error {
	var request Message

	if err := c.BodyParser(&request); err != nil {
		c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("update message endpoint called: %v\n", request))

	messageData := models.Message{
		MessageId: request.ID,
		Content:   request.Content,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.MessageBusiness.UpdateMessage(ctx, messageData); err != nil {
		slog.Debug(err.Error())
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
		Content: "Message updated",
	})

}

func (h *HTTPhandler) DeleteMessage(c *fiber.Ctx) error {
	var request Message

	if err := c.BodyParser(&request); err != nil {
		c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("delete message endpoint called: %v\n", request))

	messageData := models.Message{
		MessageId: request.ID,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	if err := h.MessageBusiness.UpdateMessage(ctx, messageData); err != nil {
		slog.Debug(err.Error())
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
		Content: "Message deleted",
	})

}
