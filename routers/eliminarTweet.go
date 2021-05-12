package routers

import (
	"net/http"

	"github.com/brayanzv/Daily/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	err := bd.BorrarTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar borrar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
