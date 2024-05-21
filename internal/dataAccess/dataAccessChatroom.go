package dataaccess

import (
	"context"
	"hackathon/migrations"
	"hackathon/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func (dao *DataAccess) CreateChatroom(ctx context.Context, chatroomData models.Chatroom) error {

	mongoChatroom := migrations.MongoChatroom{
		ChatroomId: chatroomData.ChatroomId,
		Name:       chatroomData.Name,
		IsPrivate:  chatroomData.IsPrivate,
	}

	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.D{{"guid", chatroomData.OwnerGUID}}

	update := bson.D{{"$add", bson.D{{"chatrooms", mongoChatroom}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(result)
	return nil
}
func (dao *DataAccess) UpdateChatroom(ctx context.Context, chatroomData models.Chatroom) error {

	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.D{{"guid", chatroomData.OwnerGUID}, {"chatrooms.ChatroomId", chatroomData.ChatroomId}}

	update := bson.D{{"$set", bson.D{{"chatrooms.$.name", chatroomData.Name}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(result)
	return nil
}
func (dao *DataAccess) DeleteChatroom(ctx context.Context, chatroomData models.Chatroom) error {
	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.D{{"guid", chatroomData.OwnerGUID}, {"chatrooms.ChatroomId", chatroomData.ChatroomId}}

	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(result)
	return nil
}
