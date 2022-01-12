package routers

import (
	"net/http"

	"github.com/Manolin89/GuiaGolang/bd"
	"github.com/Manolin89/GuiaGolang/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametyro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertarRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
