package models

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
	Content    string
	Image      bool
	ChatroomId string
}

type Chatroom struct {
	ChatroomId        string
	Name              string
	OwnerGUID         string
	IsPrivate         bool
	ParticipantsLimit int
}
