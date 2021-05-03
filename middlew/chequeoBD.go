package middlew

import (
	"net/http"

	"github.com/brayanzv/Daily/bd"
)

//ChequeoBD el middlew que me permites saber ek estado de la bd
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
		}
		next.ServeHTTP(w, r)
	}
}
