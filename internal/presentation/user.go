package handlers

import (
	"context"
	"errors"
	"hackathon/models"
	"log"
	"net/http"
	"time"

	"hackathon/exceptions"

	"github.com/gofiber/fiber/v2"
)

type HTTPhandler struct {
	app             *fiber.App
	UserChallen     IChallenUser
	MessageChallen  IChallenMessage
	ChatroomChallen IChallenChatroom
	WsChallen       IChallenWS
}

type IChallenUser interface {
	CreateUser(ctx context.Context, userData models.User) (*models.User, error)
	UpdateUsername(ctx context.Context, userData models.User) error
	UpdateEmail(ctx context.Context, userData models.User) error
	UpdatePassword(ctx context.Context, userData models.User) error
	DeleteUser(ctx context.Context, userData models.User) error
}

// func NewHandler(challen IChallenUser) *HTTPhandler {
// 	return
// }

func (h *HTTPhandler) createUser(c *fiber.Ctx) error {
	userDTO := new(User)

	if err := c.BodyParser(userDTO); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}

	userEntity := models.User{
		Username:       userDTO.Username,
		HashedPassword: userDTO.Password,
		Email:          userDTO.Email,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	challenData, err := h.UserChallen.CreateUser(ctx, userEntity)
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

	userEntity := models.User{
		GUID:     userDTO.GUID,
		Username: userDTO.Username,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserChallen.UpdateUsername(ctx, userEntity); err != nil {
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
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(Response{
			Error:   exceptions.ErrBadRequest.Error(),
			Content: nil,
		})
	}

	userEntity := models.User{
		GUID:  userDTO.GUID,
		Email: userDTO.Email,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserChallen.UpdateEmail(ctx, userEntity); err != nil {
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

	userEntity := models.User{
		GUID:              userDTO.GUID,
		OldHashedPassword: userDTO.OldPassword,
		HashedPassword:    userDTO.Password,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserChallen.UpdatePassword(ctx, userEntity); err != nil {
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

	userEntity := models.User{
		GUID: userDTO.GUID,
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	if err := h.UserChallen.DeleteUser(ctx, userEntity); err != nil {
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
