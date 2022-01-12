package bd

import (
	"context"
	"time"

	"github.com/Manolin89/GuiaGolang/models"
)

func BorrarRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	collect := db.Collection("relacion")

	_, err := collect.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
