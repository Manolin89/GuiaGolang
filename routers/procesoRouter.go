package routers

import (
	"errors"
	"strings"

	"github.com/Manolin89/GuiaGolang/bd"
	"github.com/Manolin89/GuiaGolang/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterDelDesarrollo-GuiaGolang")
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("Formato de Token Invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuaurio(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.Id
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("Token Invalido")
	}
	return claims, false, "", err
}
