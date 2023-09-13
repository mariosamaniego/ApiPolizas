package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"main.go/services"
	"main.go/utilities"
)

func ConsultarArticulos(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inicio ArticulosHandler::ConsultarArticulos")

	w.Header().Set("Content-Type", "application/json")
	articulos := services.ConsultarArticulos()
	json.Marshal(articulos)

	if len(articulos) == 0 {
		log.Printf("Error al consultar los Articulos")
		response.Meta.Status = utilities.StatusFail
		response.Data = "Error al consultar los Articulos"
		response, _ := json.Marshal(response)

		fmt.Fprintf(w, string(response))

		return
	}
	response.Meta.Status = "OK"
	response.Data = articulos
	response, _ := json.Marshal(response)

	fmt.Fprintf(w, string(response))

	log.Printf("Fin ArticulosHandler::ConsultarArticulos")
	return
}
