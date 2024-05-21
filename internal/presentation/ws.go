package handlers

import (
	"hackathon/models"
	"log/slog"

	"github.com/gofiber/contrib/websocket"
)

func (h *HTTPhandler) handleWS(c *websocket.Conn) {
	guid := c.Params("GUID")
	cid := c.Params("cid")

	h.hub[cid].Participants[guid] = c

	h.listenUserMessage(c, cid, guid)
}

func (h *HTTPhandler) listenUserMessage(c *websocket.Conn, cid string, guid string) {
	for {
		msg := &models.Message{}
		if err := c.ReadJSON(msg); err != nil {
			delete(h.hub[cid].Participants, guid)
			return
		}

		chatroomID := msg.ChatroomId
		room := h.hub[chatroomID]

		for GUID, conn := range room.Participants {
			if GUID != guid {
				if err := conn.WriteJSON(msg); err != nil {
					slog.Error(err.Error())
				}
			}
		}
	}
}
