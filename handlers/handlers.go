package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/GuillermoEchague/backend/middlew"
	"github.com/GuillermoEchague/backend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores seteo mi puerto, el Handler ypongo a escuchar al Servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
