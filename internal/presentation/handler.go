package handlers

import (
	"context"
	"log"
	"log/slog"

	"hackathon/models"
	"hackathon/pkg/config"
	"hackathon/pkg/jwt"

	"github.com/gofiber/contrib/websocket"
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
	GetChatrooms(ctx context.Context) ([]models.Chatroom, error)
	CreateChatroom(ctx context.Context, chatroomData models.Chatroom) (*models.Chatroom, error)
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
		hub:              make(map[string]*models.Room),
		jwtMiddleware:    jwt.New(cfg),
		addr:             cfg.Addr,
		port:             cfg.Port,
	}
}

func (h *HTTPhandler) Start() error {
	rooms, err := h.ChatroomBusiness.GetChatrooms(context.TODO())
	if err != nil {
		slog.Debug(err.Error())
	}
	rms := make([]models.Room, 0, len(rooms))

	for _, r := range rooms {
		log.Println(r)
		rms = append(rms, models.Room{CID: r.ChatroomId, Participants: make(map[string]*websocket.Conn)})
	}

	for _, r := range rms {
		h.hub[r.CID] = &r
		log.Println(r)
	}

	h.bindRoutesAndMiddlewares()
	return h.app.Listen(h.addr + ":" + h.port)
}
