package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoClient is exported Mongo Database client
var MongoClient *mongo.Client

// ConnectDatabaseMongo is used to connect the MongoDB database
func ConnectDatabaseMongo() {
	log.Println("Database connecting...")
	// Set client options

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_ATLAS_URI")))
	MongoClient = client
	if err != nil {
		log.Fatal(err)
	}
	// make sure read preferces primary & check connection
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})

	log.Println("list of databases Available", databases)

	log.Println("..... Connected to Database  ..........")
}
