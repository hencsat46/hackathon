package business

import (
	"context"
	"fmt"
	"hackathon/models"
	"log/slog"
)

func (b *business) FetchUserChatrooms(ctx context.Context, userData models.User) ([]models.Chatroom, error) {
	slog.Debug(fmt.Sprintf("fetching chatrooms: %v\n", userData))
	chatrooms, err := b.UserDataAccess.FetchUserChatrooms(ctx, userData)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	return chatrooms, nil
}
