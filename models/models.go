package models

import "github.com/gofiber/contrib/websocket"

type Room struct {
	CID          string
	Participants map[string]*websocket.Conn
}

type User struct {
	GUID              string
	Username          string
	OldHashedPassword string
	HashedPassword    string
	Email             string
}

type Message struct {
	MessageId  string
	SenderGUID string
	SenderName string
	Content    string
	Image      bool
	ChatroomId string
}

type Chatroom struct {
	ChatroomId string
	Name       string
	OwnerGUID  string
	IsPrivate  bool
}
