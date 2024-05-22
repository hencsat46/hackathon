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

func (dao *DataAccess) FetchUserChatrooms(ctx context.Context, userData models.User) ([]models.Chatroom, error) {
	var chatrooms []models.Chatroom

	coll := dao.mongoConnection.Database("ringo").Collection("chatrooms")

	filter := bson.D{{"guid", userData.GUID}}

	if err := coll.FindOne(context.TODO(), filter).Decode(&chatrooms); err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	return chatrooms, nil
}

func (dao *DataAccess) LoginUser(ctx context.Context, userData models.User) (*models.User, error) {
	slog.Debug(fmt.Sprintf("login user %v\n", userData))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"username", userData.Username}, {"password", userData.HashedPassword}}

	mongoData := new(migrations.MongoUser)

	err := coll.FindOne(context.TODO(), filter).Decode(&mongoData)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	returnValue := &models.User{
		GUID: mongoData.GUID,
	}

	return returnValue, nil
}

func (dao *DataAccess) CreateUser(ctx context.Context, userData models.User) (*models.User, error) {
	log.Println("hui")
	//slog.Debug(fmt.Sprintf("creating user %v\n", userData))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	data := migrations.MongoUser{
		GUID:           userData.GUID,
		Username:       userData.Username,
		HashedPassword: userData.HashedPassword,
		Email:          userData.Email,
	}

	_, err := coll.InsertOne(context.TODO(), data)
	log.Println("jopa")
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	user := bson.D{{"guid", userData.GUID}, {"chatrooms", primitive.A{}}}
	_, err = dao.mongoConnection.Database("ringo").Collection("chatrooms").InsertOne(context.TODO(), user)
	if err != nil {
		slog.Debug(err.Error())
		return nil, err
	}

	return &userData, nil
}

func (dao *DataAccess) UpdateUsername(ctx context.Context, userData models.User) error {
	slog.Debug(fmt.Sprintf("updating username %v\n", userData))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", userData.GUID}}

	update := bson.D{{"$set", bson.D{{"username", userData.Username}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (dao *DataAccess) UpdateEmail(ctx context.Context, userData models.User) error {
	slog.Debug(fmt.Sprintf("updating email %v\n", userData))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", userData.GUID}}

	update := bson.D{{"$set", bson.D{{"email", userData.Email}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (dao *DataAccess) UpdatePassword(ctx context.Context, userData models.User) error {
	slog.Debug(fmt.Sprintf("updating password %v\n", userData))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", userData.GUID}}

	update := bson.D{{"$set", bson.D{{"password", userData.HashedPassword}}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}

func (dao *DataAccess) DeleteUser(ctx context.Context, userData models.User) error {
	slog.Debug(fmt.Sprintf("deleting user %v\n", userData))
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", userData.GUID}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		slog.Debug(err.Error())
		return err
	}
	return nil
}
