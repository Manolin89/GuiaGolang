package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Manolin89/GuiaGolang/bd"
)

func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	openfile, err := os.Open("uploads/avatar/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openfile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
		return
	}
}
