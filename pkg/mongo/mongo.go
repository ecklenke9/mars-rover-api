package mongo

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// Create constant Collection to hold collection name
const Collection = "test"

// Create Mongo struct to be populated on the controller
// the controller will use this package to pull Rover image arrays
// from mongo by the date given
type Mongo struct {
	client *mongo.Client
}

func Connect() (*Mongo, error) {
	// Load in mongo uri from .env file
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	// Set uri from .env to var mongoUri
	mongoUri := os.Getenv("MONGO_URI")

	// Set httpclient options
	clientOptions := options.Client().ApplyURI(mongoUri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Println("Connect(): ", err)
		return nil, err
	}

	// Return Mongo client
	return &Mongo{client: client}, nil
}
