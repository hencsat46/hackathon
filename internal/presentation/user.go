package handlers

import (
	"context"
	"errors"
	"fmt"
	"hackathon/models"
	"log"
	"log/slog"
	"net/http"
	"time"

	"hackathon/exceptions"

	"github.com/gofiber/fiber/v2"
)

func (h *HTTPhandler) createUser(c *fiber.Ctx) error {
	userDTO := new(User)

	if err := c.BodyParser(userDTO); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("create user endpoint called: %v\n", userDTO))

	userEntity := models.User{
		Username:       userDTO.Username,
		HashedPassword: userDTO.Password,
		Email:          userDTO.Email,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	challenData, err := h.UserBusiness.CreateUser(ctx, userEntity)
	slog.Debug(err.Error())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Error:   exceptions.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(Response{
		Error: "nil",
		Content: User{
			GUID: challenData.GUID,
		},
	})
}

func (h *HTTPhandler) updateUsername(c *fiber.Ctx) error {
	userDTO := new(User)

	if err := c.BodyParser(userDTO); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("update username endpoint called: %v\n", userDTO))

	userEntity := models.User{
		GUID:     userDTO.GUID,
		Username: userDTO.Username,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserBusiness.UpdateUsername(ctx, userEntity); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Error:   exceptions.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(Response{
		Error:   "nil",
		Content: "User updated",
	})
}

func (h *HTTPhandler) updateEmail(c *fiber.Ctx) error {
	userDTO := new(User)

	if err := c.BodyParser(userDTO); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("update email endpoint called: %v\n", userDTO))

	userEntity := models.User{
		GUID:  userDTO.GUID,
		Email: userDTO.Email,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserBusiness.UpdateEmail(ctx, userEntity); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Error:   exceptions.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(Response{
		Error:   "nil",
		Content: "Email updated",
	})
}

func (h *HTTPhandler) updatePassword(c *fiber.Ctx) error {
	userDTO := new(User)

	if err := c.BodyParser(userDTO); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("update password endpoint called: %v\n", userDTO))

	userEntity := models.User{
		GUID:              userDTO.GUID,
		OldHashedPassword: userDTO.OldPassword,
		HashedPassword:    userDTO.Password,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserBusiness.UpdatePassword(ctx, userEntity); err != nil {
		slog.Debug(err.Error())
		if errors.Is(err, exceptions.ErrPasswordIncorrect) {
			return c.Status(http.StatusBadRequest).JSON(Response{
				Error:   exceptions.ErrPasswordIncorrect.Error(),
				Content: nil,
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Error:   exceptions.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(Response{
		Error:   "nil",
		Content: "Password updated",
	})
}

func (h *HTTPhandler) DeleteUser(c *fiber.Ctx) error {
	userDTO := new(User)

	if err := c.BodyParser(userDTO); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("delete user endpoint called: %v\n", userDTO))

	userEntity := models.User{
		GUID: userDTO.GUID,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserBusiness.DeleteUser(ctx, userEntity); err != nil {
		slog.Debug(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(Response{
			Error:   exceptions.ErrInternalServerError.Error(),
			Content: nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(Response{
		Error:   "nil",
		Content: "User deleted",
	})
}

func (h *HTTPhandler) fetchUserChatrooms(c *fiber.Ctx) error {
	userGUID := c.Params("guid")

	if len(userGUID) == 0 {
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}
	slog.Debug(fmt.Sprintf("fetch user's chatrooms endpoint called: %v\n", userGUID))

	chatroomData := models.Chatroom{
		OwnerGUID: userGUID,
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	chatrooms, err := h.UserBusiness.FetchUserChatrooms(ctx, chatroomData)
	if err != nil {
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
		Content: chatrooms,
	})
}
