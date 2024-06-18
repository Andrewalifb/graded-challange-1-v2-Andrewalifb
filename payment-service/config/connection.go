package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	// viper.SetConfigFile("config/.env")
	// viper.SetConfigFile("/app/config/.env")
	// viper.AutomaticEnv()

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatalf("Error while reading config file %s", err)
	// }

	// mongoUri := viper.GetString("MONGO_DB_URI")
	mongoUri := "mongodb://host.docker.internal:27017"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// databaseName := viper.GetString("DATABASE_NAME")
	databaseName := "payment-service"
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}
