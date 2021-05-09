package main

import (
	"log"

	"github.com/GuillermoEchague/backend/bd"
	"github.com/GuillermoEchague/backend/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexión a la BD")
		return
	}
	handlers.Manejadores()
}
