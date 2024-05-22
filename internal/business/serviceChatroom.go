package business

import (
	"context"
	"fmt"
	"hackathon/models"
	"log/slog"

	"github.com/google/uuid"
)

func (b *Business) CreateChatroom(ctx context.Context, chatroomData models.Chatroom) (*models.Chatroom, error) {
	slog.Debug(fmt.Sprintf("creating chatroom: %v\n", chatroomData))

	chatroomData.ChatroomId = uuid.NewString()

	if err := b.ChatroomDataAccess.CreateChatroom(ctx, chatroomData); err != nil {
		slog.Debug(err.Error())
		return nil, err
	}
	return &models.Chatroom{ChatroomId: chatroomData.ChatroomId}, nil
}

func (b *Business) UpdateChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	slog.Debug(fmt.Sprintf("upodating chatroom: %v\n", chatroomData))
	if err := b.ChatroomDataAccess.UpdateChatroom(ctx, chatroomData); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (b *Business) DeleteChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	slog.Debug(fmt.Sprintf("deleting chatroom: %v\n", chatroomData))
	if err := b.ChatroomDataAccess.DeleteChatroom(ctx, chatroomData); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}
