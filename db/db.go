package db

import (
	"context"
	// "encoding/json"
	// "fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var collection *mongo.collection

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
	mongoCredentials := options.Credential{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	// clientOpts := options.Client().SetAuth(mongoCredentials)
	// or
	// clientOPts := options.Client().ApplyURI(uri))
	// or
	// clientOPts := options.Client().ApplyURI(uri).SetAuth(mongoCredentials)
	// You shouldn't pass your username and password into the connection URI
	// but rather set them as options
	clientOptions := options.Client().ApplyURI(uri).SetAuth(mongoCredentials)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	log.Println("Connected to Mongodb...")
	return client, err
	// coll := client.Database("sample_mflix").Collection("movies")
	// title := "Back to the Future"

	// var result bson.M
	// err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).
	// 	Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the title %s\n", title)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }

	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)
}
