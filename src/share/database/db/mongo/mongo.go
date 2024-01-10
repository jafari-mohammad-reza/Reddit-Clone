package db

import (
	"Reddit-Clone/src/share/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var mongoClient *mongo.Client
var mongoDataBase *mongo.Database

func InitMongo(cfg *config.Config) (context.CancelFunc, error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.Mongo.User, cfg.Mongo.Password, cfg.Mongo.Host, cfg.Mongo.Port)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return cancel, err
	}
	mongoClient = client
	mongoDataBase = client.Database(cfg.Mongo.DbName)
	return cancel, nil
}

func GetClient() *mongo.Client {
	return mongoClient
}
func GetDb() *mongo.Database {
	return mongoDataBase
}
func CloseClient() error {
	return mongoClient.Disconnect(context.TODO())
}
