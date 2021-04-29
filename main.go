package main

import (
	"log"

	"github.com/brayanzv/Daily/bd"
	"github.com/brayanzv/Daily/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin Conexion a la bd")
		return
	}
	handlers.Manejadores()
}
