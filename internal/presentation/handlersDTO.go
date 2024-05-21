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
	ID         string `json:"id"`
	SenderGUID string `json:"sender_GUID"`
	Content    string `json:"message_content"`
	Image      bool   `json:"image_sent"`
}

type Response struct {
	Error   string `json:"error"`
	Content any    `json:"content"`
}
