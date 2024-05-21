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

func (b *business) CreateUser(ctx context.Context, userData models.User) (*models.User, error) {
	user, err := b.UserDataAccess.CreateUser(ctx, userData)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}
	return user, nil
}

func (b *business) UpdateUsername(ctx context.Context, userData models.User) error {
	if err := b.UserDataAccess.UpdateUsername(ctx, userData); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (b *business) UpdateEmail(ctx context.Context, userData models.User) error {
	if err := b.UserDataAccess.UpdateEmail(ctx, userData); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (b *business) UpdatePassword(ctx context.Context, userData models.User) error {
	if err := b.UserDataAccess.UpdatePassword(ctx, userData); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (b *business) DeleteUser(ctx context.Context, userData models.User) error {
	if err := b.UserDataAccess.DeleteUser(ctx, userData); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}
