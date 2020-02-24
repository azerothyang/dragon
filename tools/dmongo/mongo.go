package dmongo

import (
	"context"
	"dragon/core/dragon/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// mongodb Client
var Client *mongo.Client

func InitDB() {
	var err error
	uri := "mongodb://" + conf.Conf.Database.Mongodb.Username + ":" + conf.Conf.Database.Mongodb.Password + "@" + conf.Conf.Database.Mongodb.Host + ":" + conf.Conf.Database.Mongodb.Port + "/" + conf.Conf.Database.Mongodb.Database
	// all connect or select/query timeout
	timeout := time.Duration(conf.Conf.Database.Mongodb.Timeout) * time.Second
	clientOptions := options.Client().ApplyURI(uri).SetServerSelectionTimeout(timeout)
	Client, err = mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// return Default config mongodb database
func DefaultDB() *mongo.Database {
	return Client.Database(conf.Conf.Database.Mongodb.Database)
}
