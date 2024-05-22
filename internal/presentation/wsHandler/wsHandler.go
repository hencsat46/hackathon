package wsHandler

import (
	"context"
	"log/slog"
	"time"

	"hackathon/internal/presentation/entities"
	hubmanager "hackathon/internal/presentation/hubManager"
	messagehttphandler "hackathon/internal/presentation/messageHTTPhandler"
	"hackathon/models"

	"github.com/gofiber/contrib/websocket"
)

type WSHandler struct {
	hubManager  *hubmanager.HubManager
	WsBusiness  IBusinessWS
	msgBusiness messagehttphandler.IBusinessMessage
}

type IBusinessWS interface {
	GetUser(ctx context.Context, GUID string) (*models.User, error)
	GetChatroom(ctx context.Context, chatroomID string) (*models.Chatroom, error)
}

func New(wsBusiness IBusinessWS, hubmngr *hubmanager.HubManager, msgBusiness messagehttphandler.IBusinessMessage) *WSHandler {
	return &WSHandler{
		hubManager:  hubmngr,
		WsBusiness:  wsBusiness,
		msgBusiness: msgBusiness,
	}
}

func (h *WSHandler) HandleWS(c *websocket.Conn) {
	guid := c.Params("GUID")
	cid := c.Params("cid")

	h.hubManager.AddParticipant(c, guid, cid)
	h.listenUserMessage(c, cid, guid)
}

func (h *WSHandler) listenUserMessage(c *websocket.Conn, cid, guid string) {
	for {
		msg := &entities.Message{}
		if err := c.ReadJSON(msg); err != nil {
			slog.Debug(err.Error())
			h.hubManager.DeleteParticipant(c, cid, guid)
			return
		}
		h.hubManager.SendMessage(msg)

		message := models.Message{
			MessageId:  msg.MessageID,
			ChatroomId: msg.ChatroomID,
			SenderGUID:       msg.GUID,
			Content:    msg.Content,
			Image:      msg.Image,
		}

		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
		defer cancel()
		if err := h.msgBusiness.CreateMessage(ctx, message); err != nil {
			slog.Debug(err.Error())
		}
	}
}
