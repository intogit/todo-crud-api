package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var collection *mongo.Collection

func ConnectToMongo() (*mongo.Client, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	clientOptions := options.Client().ApplyURI(uri)
	// TODO 4:
	// check why the authSource and options.Credentials is not working
	// throwing : Error: connection() error occurred during connection handshake:
	// auth error: sasl conversation error: unable to authenticate using mechanism "SCRAM-SHA-1": (AtlasError) bad auth : Authentication failed.
	// mongoCredentials := options.Credential{
	// 	AuthSource: "ninja",
	// 	Username:   os.Getenv("DB_USERNAME"),
	// 	Password:   os.Getenv("DB_PASSWORD"),
	// }
	// clientOptions := options.Client().ApplyURI(uri).SetAuth(mongoCredentials)
	// You shouldn't pass your username and password into the connection URI
	// but rather set them as options
	// clientOptions.SetAuth(mongoCredentials)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Connected to Mongodb...")
	return client, nil
}
