package domain

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"helloworld/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoClient  *mongo.Client
)

func connect(uri string)(*mongo.Client, context.Context, context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error{
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB successfully")
	return nil
}

func InitialiseMongoDB() {
	client, ctx, _, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	MongoClient = client
	//defer close(client, ctx, cancel)
	err = ping(client, ctx)
	if err != nil {
		panic(err)
	}
}

func Create(msg *models.Message) error {
	coll := MongoClient.Database("test").Collection("messages")
	result, err := coll.InsertOne(context.TODO(), msg)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return  nil
}

func Update(msg *models.Message) error {
	coll := MongoClient.Database("test").Collection("messages")
	var res models.Message
	err := coll.FindOne(context.TODO(),bson.M{"_id":msg.Id}).Decode(&res)
	if err != nil {
		return err
	}
	msg.CreatedAt =  res.CreatedAt
	result, err := coll.UpdateOne(context.TODO(), bson.M{"_id":msg.Id},msg)
	if err != nil {
		return err
	}
	fmt.Printf("No. of documents updated: %v\n", result.ModifiedCount)
	return  nil
}