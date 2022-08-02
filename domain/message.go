package domain

import (
	"context"
	"fmt"
	"helloworld/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoClient  *mongo.Client
)

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func(){

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil{
			panic(err)
		}
	}()
}


func connect(uri string)(*mongo.Client, context.Context, context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30 * time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error{

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func InitialiseMongoDB() {
	client, _, _, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	MongoClient = client

	//defer close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	//err = ping(client, ctx)
	//if err != nil {
	//	panic(err)
	//}
}

func Create(msg *models.Message) (*models.Message, error) {
	coll := MongoClient.Database("test").Collection("messages")
	result, err := coll.InsertOne(context.TODO(), msg)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	msg.Id = result.InsertedID
	return msg, nil
}