package handlers

import (
	chatroomhttphandler "hackathon/internal/presentation/chatroomHTTPhandler"
	hubmanager "hackathon/internal/presentation/hubManager"
	messagehttphandler "hackathon/internal/presentation/messageHTTPhandler"
	userhttphandler "hackathon/internal/presentation/userHTTPhandler"
	"hackathon/internal/presentation/wsHandler"
	"hackathon/pkg/config"
	"hackathon/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type HTTPhandler struct {
	app           *fiber.App
	jwtMiddleware *jwt.JWT
	userhttphandler.UserHandler
	messagehttphandler.MessageHandler
	chatroomhttphandler.ChatroomHandler
	wsHandler.WSHandler
	*hubmanager.HubManager
	addr string
	port string
}

func NewHandler(cfg *config.Config, app *fiber.App, loader hubmanager.ILoader, messageBusiness messagehttphandler.IBusinessMessage, userBusiness userhttphandler.IBusinessUser, jwt *jwt.JWT, wsBusiness wsHandler.IBusinessWS) *HTTPhandler {
	hubmngr := hubmanager.New(loader)
	return &HTTPhandler{
		HubManager:      hubmngr,
		jwtMiddleware:   jwt,
		ChatroomHandler: *chatroomhttphandler.New(hubmngr),
		MessageHandler:  *messagehttphandler.New(messageBusiness),
		UserHandler:     *userhttphandler.New(userBusiness, jwt),
		WSHandler:       *wsHandler.New(wsBusiness, hubmngr, messageBusiness),
		app:             app,
		addr:            cfg.Addr,
		port:            cfg.Port,
	}
}

func (h *HTTPhandler) Start() error {
	h.LoadAllChatroomsToHub()
	h.bindRoutesAndMiddlewares()
	return h.app.Listen(h.addr + ":" + h.port)
}

func (h *HTTPhandler) Shutdown() error {
	return h.app.Shutdown()
}
