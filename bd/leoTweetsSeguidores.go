package bd

import (
	"context"
	"time"

	"github.com/Manolin89/GuiaGolang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelveTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	collect := db.Collection("relacion")

	skip := (pagina - 1) * 20
	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{
		"$match": bson.M{
			"usuarioid": ID,
		},
	})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	condiciones = append(condiciones, bson.M{
		"$unwind": "$tweet",
	})
	condiciones = append(condiciones, bson.M{
		"$sort": bson.M{
			"fecha": -1,
		},
	})
	condiciones = append(condiciones, bson.M{
		"$skip": skip,
	})
	condiciones = append(condiciones, bson.M{
		"$limit": 20,
	})

	cursor, err := collect.Aggregate(ctx, condiciones)
	var result []models.DevuelveTweetsSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
