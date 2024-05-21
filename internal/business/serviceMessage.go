package business

import (
	"context"
	"hackathon/models"
	"log"
)

func (b *business) CreateMessage(ctx context.Context, messageData models.Message) error {
	if err := b.MessageDataAccess.CreateMessage(ctx, messageData); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (b *business) UpdateMessage(ctx context.Context, messageData models.Message) error {
	if err := b.MessageDataAccess.UpdateMessage(ctx, messageData); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (b *business) DeleteMessage(ctx context.Context, messageData models.Message) error {
	if err := b.MessageDataAccess.DeleteMessage(ctx, messageData); err != nil {
		log.Println(err)
		return err
	}
	return nil
}