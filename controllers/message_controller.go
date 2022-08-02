package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"helloworld/domain"
)

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{})(*mongo.InsertOneResult, error) {

// select database and collection ith Client.Database method
// and Database.Collection method
collection := client.Database(dataBase).Collection(col)

// InsertOne accept two argument of type Context
// and of empty interface
result, err := collection.InsertOne(ctx, doc)
return result, err
}

func CreateMessage(c *gin.Context) {
	coll := domain.MongoClient.Database("test").Collection("messages")

	var doc interface{}

	doc = bson.D{
		{"rollNo", 175},
		{"maths", 80},
		{"science", 90},
		{"computer", 95},
	}

	result, err := coll.InsertOne(context.TODO(), doc)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

}

func UpdateMessage(c *gin.Context) {

}
