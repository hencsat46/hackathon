package dataaccess

import (
	"context"
	"fmt"
	"hackathon/migrations"
	"hackathon/models"
	"log"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dao *DataAccess) CreateChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	slog.Debug(fmt.Sprintf("creating chatroom %v\n", chatroomData))

	mongoChatroom := migrations.MongoChatroom{
		ChatroomId: chatroomData.ChatroomId,
		Name:       chatroomData.Name,
		OwnerGUID:  chatroomData.OwnerGUID,
		IsPrivate:  chatroomData.IsPrivate,
		Messages:   primitive.A{},
	}

	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	if _, err := coll.InsertOne(ctx, mongoChatroom); err != nil {
		slog.Debug(err.Error())
		return err
	}

	coll = dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.M{"guid": chatroomData.OwnerGUID}

	update := bson.M{"$push": bson.M{"chatrooms": chatroomData.ChatroomId}}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (dao *DataAccess) UpdateChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	slog.Debug(fmt.Sprintf("updating chatroom %v\n", chatroomData))

	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.M{"chatroom_id": chatroomData.ChatroomId}

	update := bson.M{"$set": bson.M{"name": chatroomData.Name}}

	if _, err := coll.UpdateOne(context.TODO(), filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (dao *DataAccess) DeleteChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	slog.Debug(fmt.Sprintf("deleting chatroom %v\n", chatroomData))
	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.M{"chatroom_id": chatroomData.ChatroomId}

	if _, err := coll.DeleteOne(context.TODO(), filter); err != nil {
		slog.Debug(err.Error())
		return err
	}

	coll = dao.mongoConnection.Database("ringo").Collection("users")
	log.Println(chatroomData.ChatroomId)
	log.Println(chatroomData.OwnerGUID)
	filter = bson.M{"guid": chatroomData.OwnerGUID}
	update := bson.M{"$pull": bson.M{"chatrooms": chatroomData.ChatroomId}}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (dao *DataAccess) GetChatrooms(ctx context.Context) ([]models.Chatroom, error) {
	slog.Debug("getting chatrooms")
	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")
	var chats []models.Chatroom

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	for cursor.Next(ctx) {
		var chat models.Chatroom
		if err := cursor.Decode(&chat); err != nil {
			slog.Debug(err.Error())
			return nil, err
		}

		chats = append(chats, chat)
	}

	return chats, nil
}
