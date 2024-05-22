package main

import (
	"context"
	"hackathon/internal/business"
	dataaccess "hackathon/internal/dataAccess"
	handlers "hackathon/internal/presentation"
	"hackathon/pkg/config"
	"hackathon/pkg/jwt"
	"hackathon/pkg/logger"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
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
	business := business.NewService(dataaccess, dataaccess, dataaccess)
	pres := handlers.NewHandler(cfg, fiber.New(), business, business, business, business, jwt.New(cfg))

	if err := pres.Start(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
