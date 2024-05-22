package business

import (
	"context"
	"hackathon/models"
)

type Business struct {
	ChatroomDataAccess IDataAccessChatroom
	UserDataAccess     IDataAccessUser
	MessageDataAccess  IDataAccessMessage
}

type IDataAccessChatroom interface {
	GetChatrooms(ctx context.Context) ([]models.Chatroom, error)
	CreateChatroom(ctx context.Context, chatroomData models.Chatroom) error
	UpdateChatroom(ctx context.Context, chatroomData models.Chatroom) error
	DeleteChatroom(ctx context.Context, chatroomData models.Chatroom) error
}

type IDataAccessUser interface {
	FetchUserChatrooms(ctx context.Context, userData models.User) ([]models.Chatroom, error)
	LoginUser(ctx context.Context, userData models.User) (*models.User, error)
	CreateUser(ctx context.Context, userData models.User) (*models.User, error)
	UpdateUsername(ctx context.Context, userData models.User) error
	UpdateEmail(ctx context.Context, userData models.User) error
	UpdatePassword(ctx context.Context, userData models.User) error
	DeleteUser(ctx context.Context, userData models.User) error
}

type IDataAccessMessage interface {
	FetchMessagesForChatroom(ctx context.Context, chatroomData models.Chatroom) ([]models.Message, error)
	CreateMessage(ctx context.Context, messageData models.Message) error
	UpdateMessage(ctx context.Context, messageData models.Message) error
	DeleteMessage(ctx context.Context, messageData models.Message) error
}

func NewService(chatroomDataAccess IDataAccessChatroom, userDataAccess IDataAccessUser, messageDataAccess IDataAccessMessage) *Business {
	return &Business{
		ChatroomDataAccess: chatroomDataAccess,
		UserDataAccess:     userDataAccess,
		MessageDataAccess:  messageDataAccess,
	}
}

func (b *Business) GetUser(ctx context.Context, userData models.User) (*models.User, error) {
	panic("not implemented")
}
func (b *Business) GetChatroom(ctx context.Context, chatroomData models.Chatroom) (*models.Chatroom, error) {
	panic("not implemented")
}
