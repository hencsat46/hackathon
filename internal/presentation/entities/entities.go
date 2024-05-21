package entities

import "github.com/gofiber/contrib/websocket"

type Room struct {
	CID          string
	Participants map[string]*websocket.Conn
}

type Message struct {
	GUID       string `json:"guid"`
	ChatroomID string `json:"chatroom_id"`
	Content    string `json:"content"`
}
