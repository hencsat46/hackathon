package migrations

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoUsers struct {
	GUID      string          `bson:"guid"`
	Chatrooms []MongoChatroom `bson:"chatrooms"`
}

type MongoChatrooms struct {
	ChatroomId   string         `bson:"chatroom_id"`
	ChatroomData []MongoMessage `bson:"chatroom_data"`
}

type MongoChatroomData struct {
	UserId    string              `bson:"user_id"`
	Text      string              `bson:"text"`
	Timestamp primitive.Timestamp `bson:"timestamp"`
}

type MongoChatroom struct {
	ChatroomId string `bson:"chatroom_id"`
	Name       string `bson:"name"`
	IsPrivate  bool   `bson:"isPrivate"`
}

type MongoUser struct {
	GUID           string `bson:"guid"`
	Username       string `bson:"username"`
	HashedPassword string `bson:"password"`
	Email          string `bson:"email"`
}

type MongoMessage struct {
	MessageId  string `bson:"message_id"`
	ChatroomId string `bson:"chatroom_id"`
	SenderGUID string `bson:"sender_guid"`
	Content    string `bson:"content"`
	Image      bool   `bson:"image"`
}
