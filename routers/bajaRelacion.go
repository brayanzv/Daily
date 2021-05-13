package routers

import (
	"fmt"
	"net/http"

	"github.com/brayanzv/Daily/bd"
	"github.com/brayanzv/Daily/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID
	fmt.Println(t.UsuarioID)
	fmt.Println(t.UsuarioRelacionID)
	status, err := bd.BorroRelacion(t)

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
