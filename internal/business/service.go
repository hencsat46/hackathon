package business

import (
	"context"
	"hackathon/models"
)

type business struct {
	ChatroomDataAccess IDataAccessChatroom
	UserDataAccess     IDataAccessUser
	MessageDataAccess  IDataAccessMessage
}

type IDataAccessChatroom interface {
	CreateChatroom(ctx context.Context, chatroomData models.Chatroom) error
	UpdateChatroom(ctx context.Context, chatroomData models.Chatroom) error
	DeleteChatroom(ctx context.Context, chatroomData models.Chatroom) error
}

type IDataAccessUser interface {
	FetchUserChatrooms(ctx context.Context, userData models.User) ([]models.Chatroom, error)
	CreateUser(ctx context.Context, userData models.User) (*models.User, error)
	UpdateUsername(ctx context.Context, userData models.User) error
	UpdateEmail(ctx context.Context, userData models.User) error
	UpdatePassword(ctx context.Context, userData models.User) error
	DeleteUser(ctx context.Context, userData models.User) error
}

type IDataAccessMessage interface {
}

