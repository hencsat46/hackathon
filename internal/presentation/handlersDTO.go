package handlers

type Chatroom struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	OwnerGUID         string `json:"owner_GUID"`
	IsPrivate         bool   `json:"is_private"`
	ParticipantsLimit int    `json:"participants_limit"`
}

type User struct {
	GUID        string `json:"GUID"`
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	Password    string `json:"password"`
	Email       string `json:"email"`
}

type Message struct {
	MessageId  string `bson:"message_id" json:"message_id"`
	ChatroomId string `bson:"chatroom_id" json:"chatroom_id"`
	SenderGUID string `bson:"sender_guid" json:"sender_guid"`
	SenderName string `bson:"sender_name" json:"sender_name"`
	Content    string `bson:"content" json:"content"`
	Image      bool   `bson:"image" json:"image"`
}

type Response struct {
	Error   string `json:"error"`
	Content any    `json:"content"`
}
