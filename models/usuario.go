package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID              primitive.ObjectID `bson: "_id, omitempty" json: "id" `
	Nombre          string             `bson: "nombre, omitempty" json: "nombre, omitempty" `
	Apellidos       string             `bson: "apellidos" json: "apellidos, omitempty" `
	FechaNacimiento time.Time          `bson: "fechaNacimiento" json: "fechaNacimiento, omitempty" `
	Email           string             `bson: "email" json: "email" `
	Password        string             `bson: "pasword" json: "password, omitempty" `
	Avatar          string             `bson: "avatar" json: "avatar, omitempty" `
	Banner          string             `bson: "banner" json: "banner, omitempty" `
	Biografia       string             `bson: "biografia" json: "biografia, omitempty" `
	Ubicacion       string             `bson: "ubicacion" json: "ubicacion, omitempty" `
	SitioWeb        string             `bson: "sitioWeb" json: "sitioWeb, omitempty" `
}
