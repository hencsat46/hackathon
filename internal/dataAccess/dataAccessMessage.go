package dataaccess

import (
	"context"
	"hackathon/migrations"
	"hackathon/models"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DataAccess) CreateMessage(ctx context.Context, messageData models.Message) error {
	coll := dao.mongoConnection.Database("ringo").Collection("messages")

	filter := bson.D{{"chatroom_id", messageData.ChatroomId}}

	data := migrations.MongoMessage{
		MessageId:  messageData.MessageId,
		ChatroomId: messageData.ChatroomId,
		SenderGUID: messageData.SenderGUID,
		Content:    messageData.Content,
		Image:      messageData.Image,
	}

	update := bson.D{{"$addToSet", bson.D{{"chatroom_data", data}}}}

	if _, err := coll.UpdateOne(context.TODO(), filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}
func (dao *DataAccess) UpdateMessage(ctx context.Context, messageData models.Message) error {
	coll := dao.mongoConnection.Database("ringo").Collection("messages")

	filter := bson.D{{"chatroom_id", messageData.ChatroomId}, {"chatroom_data.message_id", messageData.MessageId}}
	update := bson.D{{"$set", bson.D{{"chatroom_data.$.content", messageData.Content}}}}

	if _, err := coll.UpdateOne(context.TODO(), filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}
func (dao *DataAccess) DeleteMessage(ctx context.Context, messageData models.Message) error {
	coll := dao.mongoConnection.Database("ringo").Collection("messages")

	filter := bson.D{{"chatroom_id", messageData.ChatroomId}, {"chatroom_data.message_id", messageData.MessageId}}

	if _, err := coll.DeleteOne(context.TODO(), filter); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}
