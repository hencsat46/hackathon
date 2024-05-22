package handlers

import (
	"context"

	"hackathon/models"
	"hackathon/pkg/config"
	"hackathon/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type HTTPhandler struct {
	app              *fiber.App
	UserBusiness     IBusinessUser
	MessageBusiness  IBusinessMessage
	ChatroomBusiness IBusinessChatroom
	WsBusiness       IBusinessWS
	hub              map[string]*models.Room
	jwtMiddleware    *jwt.JWT
	addr             string
	port             string
}

type IBusinessWS interface {
	GetUser(ctx context.Context, userData models.User) (*models.User, error)
	GetChatroom(ctx context.Context, chatroomData models.Chatroom) (*models.Chatroom, error)
}

type IBusinessMessage interface {
	CreateMessage(ctx context.Context, messageData models.Message) error
	FetchMessagesForChatroom(ctx context.Context, chatroomData models.Chatroom) ([]models.Message, error)
	UpdateMessage(ctx context.Context, messageData models.Message) error
	DeleteMessage(ctx context.Context, messageData models.Message) error
}

type IBusinessChatroom interface {
	CreateChatroom(ctx context.Context, chatroomData models.Chatroom) error
	UpdateChatroom(ctx context.Context, chatroomData models.Chatroom) error
	DeleteChatroom(ctx context.Context, chatroomData models.Chatroom) error
}

type IBusinessUser interface {
	FetchUserChatrooms(ctx context.Context, userData models.User) ([]models.Chatroom, error)
	LoginUser(ctx context.Context, userData models.User) (*models.User, error)
	CreateUser(ctx context.Context, userData models.User) (*models.User, error)
	UpdateUsername(ctx context.Context, userData models.User) error
	UpdateEmail(ctx context.Context, userData models.User) error
	UpdatePassword(ctx context.Context, userData models.User) error
	DeleteUser(ctx context.Context, userData models.User) error
}

func NewHandler(cfg *config.Config, app *fiber.App, userCh IBusinessUser, msgCh IBusinessMessage, chatroomCh IBusinessChatroom, ws IBusinessWS) *HTTPhandler {
	return &HTTPhandler{
		app:              app,
		UserBusiness:     userCh,
		MessageBusiness:  msgCh,
		ChatroomBusiness: chatroomCh,
		WsBusiness:       ws,
		addr:             cfg.Addr,
		port:             cfg.Port,
	}
}

func (h *HTTPhandler) Start() error {
	h.bindRoutesAndMiddlewares()
	return h.app.Listen(h.addr + ":" + h.port)
}
