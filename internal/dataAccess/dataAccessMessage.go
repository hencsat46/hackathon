package dataaccess

import (
	"context"
	"hackathon/migrations"
	"hackathon/models"
	"log"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DataAccess) FetchMessagesForChatroom(ctx context.Context, chatroomID string) ([]models.Message, error) {
	var messages []models.Message

	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.M{"chatroom_id": chatroomID}

	if err := coll.FindOne(context.TODO(), filter).Decode(&messages); err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	log.Println(messages)

	return messages, nil
}

func (dao *DataAccess) CreateMessage(ctx context.Context, message models.Message) error {
	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.M{"chatroom_id": message.ChatroomId}

	data := migrations.MongoMessage{
		MessageId:  message.MessageId,
		ChatroomId: message.ChatroomId,
		SenderGUID: message.SenderGUID,
		SenderName: message.SenderName,
		Content:    message.Content,
		Image:      message.Image,
	}

	update := bson.D{{"$push", bson.D{{"messages", data}}}}

	if _, err := coll.UpdateOne(context.TODO(), filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (dao *DataAccess) UpdateMessage(ctx context.Context, newContent, messageID string) error {
	coll := dao.mongoConnection.Database("ringo").Collection("messages")

	filter := bson.D{{"chatroom_id", messageData.ChatroomId}, {"chatrooms.message_id", messageData.MessageId}}
	update := bson.D{{"$set", bson.D{{"chatrooms.$.content", messageData.Content}}}}

	if _, err := coll.UpdateOne(context.TODO(), filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (dao *DataAccess) DeleteMessage(ctx context.Context, messageData models.Message) error {
	coll := dao.mongoConnection.Database("ringo").Collection("messages")

	filter := bson.D{{"chatroom_id", messageData.ChatroomId}, {"chatrooms.message_id", messageData.MessageId}}

	if _, err := coll.DeleteOne(context.TODO(), filter); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}
