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
