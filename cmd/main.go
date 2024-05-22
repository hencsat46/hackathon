package main

import (
	"context"

	chatroomservice "hackathon/internal/business/chatroomService"
	hubloaderservice "hackathon/internal/business/hubLoaderService"
	messageservice "hackathon/internal/business/messageService"
	userservice "hackathon/internal/business/userService"
	wsservice "hackathon/internal/business/wsService"
	dataaccess "hackathon/internal/dataAccess"
	"hackathon/pkg/config"
	"hackathon/pkg/logger"
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.New()
	logger := logger.New(cfg)
	logger.SetAsDefault()

	mng, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.Mongo))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dataaccess := dataaccess.NewDataAccess(mng)

	chatroomSvc := chatroomservice.New(dataaccess)
	hubLoaderSvc := hubloaderservice.New(dataaccess)
	messageSvc := messageservice.New(dataaccess)
	userSvc := userservice.New(dataaccess)
	wsSvc := wsservice.New(dataaccess)

	if err := pres.Start(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
