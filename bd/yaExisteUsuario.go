package bd

import (
	"context"
	"time"

	"github.com/Manolin89/GuiaGolang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuaurio(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("usuarios")

	condicon := bson.M{"email": email}
	var result models.Usuario

	err := collection.FindOne(ctx, condicon).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
