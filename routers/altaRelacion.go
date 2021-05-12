package routers

import (
	"net/http"

	"github.com/brayanzv/Daily/bd"
	"github.com/brayanzv/Daily/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviarse el id", http.StatusBadRequest)
		return
	}

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertarRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar la relacion", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "no se pudo subir la relacion", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
