package models

type Relacion struct {
	UsuarioID         string `bson: "usuairoid" json: "usuarioId" `
	UsuarioRelacionID string `bson: "usuairorelacionid" json: "usuarioRelacionId" `
}
