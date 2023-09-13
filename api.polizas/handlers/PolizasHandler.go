package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"main.go/entities"
	"main.go/services"
	"main.go/utilities"
)

var response entities.Response
var respuestaService string

func ConsultarPoliza(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inicio PolizasHandler::ConsultarPoliza")

	var polizasResponse = entities.PolizasResponse{}

	queryParams := r.URL.Query()
	id := queryParams.Get("idpoliza")

	idpoliza, _ := strconv.Atoi(id)
	polizasResponse = services.ConsultarPoliza(idpoliza)

	if polizasResponse.Poliza.IdPoliza == 0 {
		log.Printf("Error al consultar la poliza")
		response.Meta.Status = utilities.StatusFail
		response.Data = utilities.ErrorPoliza
		response, _ := json.Marshal(response)

		fmt.Fprintf(w, string(response))
		return
	}

	response.Meta.Status = "OK"
	response.Data = polizasResponse
	response, _ := json.Marshal(response)
	fmt.Fprintf(w, string(response))
	log.Printf("Fin PolizasHandler::ConsultarPoliza")
}

func ConsultarPolizasEmpleado(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inicio PolizasHandler::ConsultarPolizasEmpleado")

	var polizasResponse2 = entities.PolizasResponse2{}

	queryParams := r.URL.Query()
	id := queryParams.Get("idempleado")
	idempleado, _ := strconv.Atoi(id)

	polizasResponse2 = services.ConsultarPolizasEmpleado(idempleado)

	response.Meta.Status = "OK"
	response.Data = polizasResponse2
	response, _ := json.Marshal(response)
	fmt.Fprintf(w, string(response))
	log.Printf("Fin PolizasHandler::ConsultarPolizasEmpleado")
}

func AgregarPoliza(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inicio PolizasHandler::AgregarPoliza")

	var poliza entities.Poliza
	var polizaresponse = entities.Mensaje{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &poliza); err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	polizaresponse = services.AgregarPoliza(int(poliza.IdPoliza),
		int(poliza.EmpleadoGenero), int(poliza.Sku), int(poliza.Cantidad))

	fmt.Println(polizaresponse.StatusRespuesta)
	if polizaresponse.StatusRespuesta == 0 {
		respuestaService = utilities.StatusFail
	} else {
		respuestaService = utilities.StatusOk
	}
	response.Meta.Status = respuestaService
	response.Data = polizaresponse
	response, _ := json.Marshal(response)
	fmt.Fprintf(w, string(response))
	log.Printf("Fin PolizasHandler::AgregarPoliza")
}

func ActualizarPoliza(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inicio PolizasHandler::ActualizarPoliza")

	var poliza entities.Poliza
	var empleado entities.Empleado
	var actualizaresponse = entities.Mensaje{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &poliza); err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &empleado); err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}
	if body != nil {
		log.Printf("Body recibido con Exitoso")
	}

	actualizaresponse = services.ActualizarPoliza(int(poliza.IdPoliza), int(empleado.IdEmpleado),
		int(poliza.Sku), int(poliza.Cantidad), empleado.Nombre, empleado.Apellido, empleado.Puesto)

	if actualizaresponse.StatusRespuesta == 0 {
		respuestaService = utilities.StatusFail
	} else {
		respuestaService = utilities.StatusOk
	}

	response.Meta.Status = respuestaService
	response.Data = actualizaresponse
	response, _ := json.Marshal(response)
	fmt.Fprintf(w, string(response))

	log.Printf("Fin PolizasHandler::ActualizarPoliza")
}
