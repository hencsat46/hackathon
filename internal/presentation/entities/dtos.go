package entities

type ChatroomDTO struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	OwnerGUID         string `json:"owner_GUID"`
	IsPrivate         bool   `json:"is_private"`
	ParticipantsLimit int    `json:"participants_limit"`
}

type UserDTO struct {
	GUID     string `json:"guid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpdatePasswordDTO struct {
	GUID        string `json:"GUID"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"password"`
}

type MessageDTO struct {
	MessageId  string `json:"message_id"`
	ChatroomID string `json:"chatroom_id"`
	Content    string `json:"content"`
}

type Response struct {
	Error   string `json:"error"`
	Content any    `json:"content"`
}
