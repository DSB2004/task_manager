package lib

import (
	"context"
	"log"
	"server/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client = nil

func ConnectDB() error {
	if MongoDBClient == nil {
		MONGO_URI := config.GetEnv("MONGO_URI", "")
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(MONGO_URI).SetServerAPIOptions(serverAPI)

		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			log.Fatal("Error happened while connecting database")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		pingErr := client.Ping(ctx, nil)

		if pingErr != nil {
			return pingErr
		}

		MongoDBClient = client
		log.Println("MongoDB connected successfully")
	}
	return nil
}

func DisconnectDB() {
	if MongoDBClient != nil {
		log.Println("MongoDB disconnected")
		MongoDBClient.Disconnect(context.TODO())
	}
}
