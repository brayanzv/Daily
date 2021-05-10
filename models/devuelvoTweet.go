package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoTweet struct {
	ID      primitive.ObjectID `bson: "_id" json: "_id,omitempty"`
	UserId  string             `bson: "userID" json: "userId,omitempty"`
	Mensaje string             `bson: "mensaje" json: "mensaje,omitempty"`
	fecha   time.Time          `bson: "fecha" json: "fecha,omitempty"`
}
