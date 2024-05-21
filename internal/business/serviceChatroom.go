package business

import (
	"context"
	"hackathon/models"
	"log"
)

func (b *business) CreateChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	if err := b.ChatroomDataAccess.CreateChatroom(ctx, chatroomData); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (b *business) UpdateChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	if err := b.ChatroomDataAccess.UpdateChatroom(ctx, chatroomData); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (b *business) DeleteChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	if err := b.ChatroomDataAccess.DeleteChatroom(ctx, chatroomData); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
