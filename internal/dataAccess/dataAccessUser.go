package dataaccess

import (
	"context"
	"hackathon/migrations"
	"hackathon/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dao *DataAccess) CreateUser(ctx context.Context, userData models.User) (*models.User, error) {
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	data := migrations.MongoUser{
		GUID:           userData.GUID,
		Username:       userData.Username,
		HashedPassword: userData.HashedPassword,
		Email:          userData.Email,
	}

	_, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	user := bson.D{{"guid", userData.GUID}, {"chatrooms", primitive.A{}}}
	_, err = dao.mongoConnection.Database("ringo").Collection("chatrooms").InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &userData, nil
}
func (dao *DataAccess) UpdateUsername(ctx context.Context, userData models.User) error {
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", userData.GUID}}

	update := bson.D{{"$set", bson.D{{"username", userData.Username}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(result)
	return nil
}
func (dao *DataAccess) UpdateEmail(ctx context.Context, userData models.User) error {
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", userData.GUID}}

	update := bson.D{{"$set", bson.D{{"email", userData.Email}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(result)
	return nil
}
func (dao *DataAccess) UpdatePassword(ctx context.Context, userData models.User) error {
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", userData.GUID}}

	update := bson.D{{"$set", bson.D{{"password", userData.HashedPassword}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(result)
	return nil
}
func (dao *DataAccess) DeleteUser(ctx context.Context, userData models.User) error {
	coll := dao.mongoConnection.Database("ringo").Collection("users")

	filter := bson.D{{"guid", userData.GUID}}

	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(result)
	return nil
}
