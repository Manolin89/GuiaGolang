package middlew

import (
	"net/http"

	"github.com/Manolin89/GuiaGolang/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConection() == 0 {
			http.Error(rw, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
