package dataaccess

import (
	"context"
	"fmt"
	"hackathon/migrations"
	"hackathon/models"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dao *DataAccess) CreateChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	slog.Debug(fmt.Sprintf("creating chatroom %v\n", chatroomData))

	mongoChatroom := migrations.MongoChatroom{
		ChatroomId: chatroomData.ChatroomId,
		Name:       chatroomData.Name,
		IsPrivate:  chatroomData.IsPrivate,
	}

	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.D{{"guid", chatroomData.OwnerGUID}}

	update := bson.D{{"$add", bson.D{{"chatrooms", mongoChatroom}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		slog.Debug(err.Error())
		return err
	}

	coll = dao.mongoConnection.Database("ringo").Collection("messages")

	data := bson.D{{"chatroom_id", chatroomData.ChatroomId}, {"chatroom_data", primitive.A{}}}

	if _, err := coll.InsertOne(context.TODO(), data); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (dao *DataAccess) UpdateChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	slog.Debug(fmt.Sprintf("updating chatroom %v\n", chatroomData))

	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.D{{"guid", chatroomData.OwnerGUID}, {"chatrooms.ChatroomId", chatroomData.ChatroomId}}

	update := bson.D{{"$set", bson.D{{"chatrooms.$.name", chatroomData.Name}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (dao *DataAccess) DeleteChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	slog.Debug(fmt.Sprintf("deleting chatroom %v\n", chatroomData))
	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.D{{"guid", chatroomData.OwnerGUID}, {"chatrooms.ChatroomId", chatroomData.ChatroomId}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}
