package entities

import "github.com/gofiber/contrib/websocket"

type WSHub map[string]*WSRoom

type WSRoom struct {
	CID          string
	Participants map[string]*websocket.Conn
}

type Message struct {
	MessageID  string `json:"message_id"`
	ChatroomID string `json:"chatroom_id"`
	GUID       string `json:"guid"`
	Content    string `json:"content"`
	Image      bool   `json:"image"`
}
