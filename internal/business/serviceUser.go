package business

import (
	"context"
	"hackathon/models"
	"log"
)

func (b *business) FetchUserChatrooms(ctx context.Context, userData models.User) ([]models.Chatroom, error) {
	chatrooms, err := b.UserDataAccess.FetchUserChatrooms(ctx, userData)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return chatrooms, nil
}

func (b *business) CreateUser(ctx context.Context, userData models.User) error {
	if err := b.UserDataAccess.CreateUser(ctx, userData); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (b *business) UpdateUsername(ctx context.Context, userData models.User) error {
	if err := b.UserDataAccess.UpdateUsername(ctx, userData); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (b *business) UpdateEmail(ctx context.Context, userData models.User) error {
	if err := b.UserDataAccess.UpdateEmail(ctx, userData); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (b *business) UpdatePassword(ctx context.Context, userData models.User) error {
	if err := b.UserDataAccess.UpdatePassword(ctx, userData); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (b *business) DeleteUser(ctx context.Context, userData models.User) error {
	if err := b.UserDataAccess.DeleteUser(ctx, userData); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
