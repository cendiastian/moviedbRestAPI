package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigDb struct {
	DbHost string
	DbPort string
}

func (config *ConfigDb) InitialDb() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%v:%v",
		config.DbHost,
		config.DbPort,
	)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
