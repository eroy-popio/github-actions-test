package models

import (
	"time"
)

type Message struct {
	Id        int64					`json:"id" bson:"_id"`
	Title     string    			`json:"title" bson:"title"`
	Body      string    			`json:"body" bson:"body"`
	CreatedAt time.Time 			`json:"created_at" bson:"createdAt"`
}