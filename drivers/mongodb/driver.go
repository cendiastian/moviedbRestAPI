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
	DbPort    string
	DbCluster string
	DbUser    string
	DbPass    string
}

func (config *ConfigDb) InitialDb() *mongo.Client {
	fmt.Println(config.DbCluster + config.DbUser + config.DbPass)
	uri := fmt.Sprintf("mongodb+srv://%v:%v%v",
		config.DbUser,
		config.DbPass,
		config.DbCluster,
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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connected")

	return client
}
