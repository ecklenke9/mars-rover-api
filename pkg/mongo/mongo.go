package mongo

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Dao interface {
	// TODO: add method signatures
}

// RoverCollection statically declared
const RoverCollection = "rover"

// Mongo - struct - expected to be populated on the Controller
// The Controller will use this package to pull Rover Image Arrays
// by the date given
type Mongo struct {
	client *mongo.Client
}

func Connect() (*Mongo, error) {
	// load in mongo uri from .env file
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
		//log.Fatalf("error loading .env file")
	}
	// set uri from .env to var mongoUri
	mongoUri := os.Getenv("MONGO_URI")

	// set httpclient options
	clientOptions := options.Client().ApplyURI(mongoUri)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	// return Mongo client
	return &Mongo{client: client}, nil
}
