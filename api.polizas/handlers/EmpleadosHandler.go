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

func ConsultarEmpleados(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inicio EmpleadosHandler::ConsultarEmpleados")

	w.Header().Set("Content-Type", "application/json")
	empleados := services.ConsultarEmpleados()

	if len(empleados) == 0 {
		log.Printf("Error al consultar los Empleados")
		response.Meta.Status = utilities.StatusFail
		response.Data = "Error al consultar los empleados"
		response, _ := json.Marshal(response)

		fmt.Fprintf(w, string(response))
		log.Printf("Fin EmpleadosHandler::ConsultarEmpleados")
		return
	}
	response.Meta.Status = "OK"
	response.Data = empleados
	response, _ := json.Marshal(response)

	fmt.Fprintf(w, string(response))

	log.Printf("Fin EmpleadosHandler::ConsultarEmpleados")
	return
}

func AgregarEmpleado(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inicio EmpleadosHandler::AgregarEmpleado")
	var empleado entities.Empleado
	var empleadoresponse = entities.Mensaje{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &empleado); err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	empleadoresponse = services.AgregarEmpleado(int(empleado.IdEmpleado), empleado.Nombre, empleado.Apellido, empleado.Puesto)

	if empleadoresponse.StatusRespuesta == 0 {
		respuestaService = utilities.StatusFail
	} else {
		respuestaService = utilities.StatusOk
	}
	response.Meta.Status = respuestaService
	response.Data = empleadoresponse
	response, _ := json.Marshal(response)
	fmt.Fprintf(w, string(response))
	log.Printf("Fin EmpleadosHandler::AgregarEmpleado")
}

func EliminarGeneral(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inicio EmpleadosHandler::EliminarGeneral")

	var responseeliminar = entities.Mensaje{}

	queryParams := r.URL.Query()
	id := queryParams.Get("opcion")
	idelimina := queryParams.Get("eliminar")

	opcion, _ := strconv.Atoi(id)
	eliminar, _ := strconv.Atoi(idelimina)

	responseeliminar = services.EliminarGeneral(opcion, eliminar)

	response.Meta.Status = "OK"
	response.Data = responseeliminar
	response, _ := json.Marshal(response)
	fmt.Fprintf(w, string(response))
	log.Printf("Fin EmpleadosHandler::EliminarGeneral")
}
