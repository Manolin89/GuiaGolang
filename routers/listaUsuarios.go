package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Manolin89/GuiaGolang/bd"
)

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)
	result, status := bd.LeerUsuauriosTodos(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
