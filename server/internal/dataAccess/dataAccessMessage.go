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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoArray struct {
	Id    primitive.ObjectID
	Array []models.Message `bson:"messages"`
}

func (dao *DataAccess) FetchMessagesForChatroom(ctx context.Context, chatroomID string) ([]models.Message, error) {
	//var messages migrations.MongoChatroom
	messages := make([]models.Message, 0)
	var array migrations.MongoChatroom
	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.M{"chatroom_id": chatroomID}
	projection := bson.M{"messages": true, "_id": false}
	log.Println("chatroom id: ", chatroomID, projection)

	options := options.FindOne().SetProjection(bson.M{"messages": true})

	if err := coll.FindOne(context.TODO(), filter, options).Decode(&array); err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	for _, v := range array.Messages {
		messages = append(messages, models.Message{
			MessageId:  v.(primitive.D)[0].Value.(string),
			ChatroomId: v.(primitive.D)[1].Value.(string),
			SenderGUID: v.(primitive.D)[2].Value.(string),
			SenderName: v.(primitive.D)[3].Value.(string),
			Content:    v.(primitive.D)[4].Value.(string),
			Image:      v.(primitive.D)[5].Value.(bool),
		})
		log.Println()
	}
	// for _,  := range messages.Messages {
	// 	//log.Printf("%T", v)
	// }

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

	update := bson.M{"$push": bson.M{"messages": data}}

	if _, err := coll.UpdateOne(context.TODO(), filter, update); err != nil {
		slog.Debug(err.Error())
		return err
	}

	return nil
}

func (dao *DataAccess) UpdateMessage(ctx context.Context, newContent, messageID, chatroomId string) error {
	slog.Debug(fmt.Sprintf("updating message: message id: %v, chatroom id: %v, new content: %v", messageID, chatroomId, newContent))

	coll := dao.mongoConnection.Database("ringo").Collection("messages")

	filter := bson.M{"chatroom_id": chatroomId, "chatrooms.message_id": messageID}

	update := bson.M{"$set": bson.M{"messages.content": newContent}}

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
