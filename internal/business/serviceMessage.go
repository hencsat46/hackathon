package business

import (
	"context"
	"hackathon/models"
	"log/slog"
)

func (b *Business) FetchMessagesForChatroom(ctx context.Context, chatroomData models.Chatroom) ([]models.Message, error) {
	messages, err := b.MessageDataAccess.FetchMessagesForChatroom(ctx, chatroomData)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	return messages, nil
}

func (b *Business) CreateMessage(ctx context.Context, messageData models.Message) error {
	if err := b.MessageDataAccess.CreateMessage(ctx, messageData); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}
func (b *Business) UpdateMessage(ctx context.Context, messageData models.Message) error {
	if err := b.MessageDataAccess.UpdateMessage(ctx, messageData); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}
func (b *Business) DeleteMessage(ctx context.Context, messageData models.Message) error {
	if err := b.MessageDataAccess.DeleteMessage(ctx, messageData); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}
