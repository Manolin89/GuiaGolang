package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/Manolin89/GuiaGolang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultarRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	collect := db.Collection("usuarios")

	condicion := bson.M{
		"usuario":           t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := collect.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
