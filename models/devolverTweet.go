package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevolverTweet struct {
	ID      primitive.ObjectID `bson: "_id" json: "_id, omitepmty" `
	UserID  string             `bson: "userid" json: "userid, omitepmty" `
	Mensaje string             `bson: "mensaje" json: "mensaje, omitepmty" `
	Fecha   time.Time          `bson: "fecha" json: "fecha, omitepmty" `
}
