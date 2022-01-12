package routers

import (
	"net/http"

	"github.com/Manolin89/GuiaGolang/bd"
	"github.com/Manolin89/GuiaGolang/models"
)

func ConsultarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametyro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID
	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultarRelacion(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
