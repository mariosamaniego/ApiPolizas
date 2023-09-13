package services

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"main.go/entities"
)

var urlPostgress = os.Getenv("urlPostgress")
var response entities.Response

func ConsultarPoliza(idpoliza int) entities.PolizasResponse {
	log.Printf("Inicio PolizasService::ConsultarPoliza")

	// Conectarse a la base de datos
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err)
		return entities.PolizasResponse{}
	}
	defer db.Close()

	// Ejecutar la consulta a la función de PostgreSQL
	rows, err := db.Query("SELECT id_polizas, empleado_genero, sku, cantidad, nombre,  apellido, puesto from fun_consultapoliza($1::INTEGER)", idpoliza)
	if err != nil {
		log.Printf("Error al ejecutar la funcion fun_consultapoliza()")
		log.Printf("", err)
		return entities.PolizasResponse{}
	}
	defer rows.Close()

	//polizas := entities.Polizas{}
	poliza := entities.Poliza{}
	empleado := entities.Empleado{}

	// Recorrer los resultados
	if rows.Next() {

		err := rows.Scan(&poliza.IdPoliza, &poliza.EmpleadoGenero, &poliza.Sku, &poliza.Cantidad, &empleado.Nombre, &empleado.Apellido, &empleado.Puesto)

		if err != nil {
			log.Printf("Error al escanear los resultados de la funcion fun_consultapolizas()")
			log.Printf("", err)
			return entities.PolizasResponse{}

		}

		var polizasResponse = entities.PolizasResponse{}
		if poliza.IdPoliza == 0 {
			return entities.PolizasResponse{}
		}
		polizasResponse.Poliza.IdPoliza = poliza.IdPoliza
		polizasResponse.Poliza.EmpleadoGenero = poliza.EmpleadoGenero
		polizasResponse.Poliza.Sku = poliza.Sku
		polizasResponse.Poliza.Cantidad = poliza.Cantidad
		polizasResponse.Empleado.IdEmpleado = poliza.EmpleadoGenero
		polizasResponse.Empleado.Nombre = empleado.Nombre
		polizasResponse.Empleado.Apellido = empleado.Apellido
		polizasResponse.Empleado.Puesto = empleado.Puesto

		log.Printf("Termina PolizasService::ConsultarPoliza")
		return polizasResponse
	}
	return entities.PolizasResponse{}

}

func ConsultarPolizasEmpleado(idempleado int) entities.PolizasResponse2 {
	log.Printf("Inicio PolizasService::ConsultarPolizasEmpleado")

	// Conectarse a la base de datos
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error al conectarse a la base de datos")
		log.Printf("", err)
		return entities.PolizasResponse2{}
	}
	defer db.Close()

	// Ejecutar la consulta a la función de PostgreSQL
	rows, err := db.Query("SELECT id_polizas,empleado_genero, sku, cantidad, nombre,apellido, puesto FROM fun_consultapolizasempleado($1::INTEGER)", idempleado)
	if err != nil {
		log.Printf("Error al ejecutar la funcion fun_consultapolizasempleado()")
		log.Printf("", err)
		return entities.PolizasResponse2{}
	}
	defer rows.Close()

	polizasresponse2 := entities.PolizasResponse2{}
	// Recorrer los resultados
	for rows.Next() {
		polizasresponse := entities.PolizasResponse{}
		err := rows.Scan(&polizasresponse.Poliza.IdPoliza,
			&polizasresponse.Poliza.EmpleadoGenero,
			&polizasresponse.Poliza.Sku,
			&polizasresponse.Poliza.Cantidad,
			&polizasresponse.Empleado.Nombre,
			&polizasresponse.Empleado.Apellido,
			&polizasresponse.Empleado.Puesto)
		polizasresponse.Empleado.IdEmpleado = polizasresponse.Poliza.EmpleadoGenero
		polizasresponse2 = append(polizasresponse2, polizasresponse)
		if err != nil {
			log.Fatal(err)
			return polizasresponse2
		}

	}

	return polizasresponse2

}

func AgregarPoliza(idpoliza int, idempleado int, sku int, cantidad int) entities.Mensaje {
	log.Printf("Inicio PolizasService::AgregarPoliza")

	var irespuesta int32
	mensaje := entities.Mensaje{}

	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT fun_registrarpoliza FROM fun_registrarpoliza($1::INTEGER,$2::INTEGER,$3::INTEGER,$4::INTEGER)", idpoliza, idempleado, sku, cantidad)
	if err != nil {
		log.Printf("Error al ejecutar la funcion fun_registrarpoliza()")
		log.Printf("", err)

	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&irespuesta)
		if err != nil {
			log.Printf("Error al escanear los resultados de la funcion fun_registraempleado()")
			log.Printf("", err)

		}

		switch irespuesta {
		case 1:
			mensaje.Respuesta = "Registro Exitoso"
		case 2:
			mensaje.Respuesta = "No existe empleado para registrar"
		case 3:
			mensaje.Respuesta = "No se encontro SKU con ese codigo"
		case 4:
			mensaje.Respuesta = "No Existe cantidad existente del articulo"
		case 5:
			mensaje.Respuesta = "Poliza ya ha sido registrada, intente con otra nuevamente"
		default:
			mensaje.Respuesta = "No se pudo realizar el registro de la poliza"
		}

		if irespuesta == 1 {
			mensaje.StatusRespuesta = 1
		} else {
			mensaje.StatusRespuesta = 0
		}

		log.Printf("Termina PolizasService::AgregarPoliza")
	}
	return mensaje

}

func ActualizarPoliza(idpoliza int, IdEmpleado int, sku int, cantidad int, nombre string, apellido string, puesto string) entities.Mensaje {
	log.Printf("Inicio PolizasService::ActualizarPoliza")

	var irespuesta int32
	mensaje := entities.Mensaje{}

	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT fun_actualizarpoliza FROM fun_actualizarpoliza($1::INTEGER,$2::INTEGER,$3::INTEGER,$4::INTEGER,$5,$6,$7)",
		idpoliza, IdEmpleado, sku, cantidad, nombre, apellido, puesto)
	if err != nil {
		log.Printf("Error al ejecutar la funcion fun_actualizarpoliza()")
		log.Printf("", err)

	}
	defer rows.Close()

	if rows.Next() {

		err := rows.Scan(&irespuesta)

		if err != nil {
			log.Printf("Error al escanear los resultados de la funcion fun_actualizarpoliza()")
			log.Printf("", err)

		}
		switch irespuesta {
		case 1:
			mensaje.Respuesta = "Registro Exitoso"
		case 2:
			mensaje.Respuesta = "No existe empleado para registrar"
		case 3:
			mensaje.Respuesta = "No se encontro SKU con ese codigo"
		case 4:
			mensaje.Respuesta = "No Existe cantidad existente del articulo"
		case 5:
			mensaje.Respuesta = "Poliza ya ha sido registrada, intente con otra nuevamente"
		default:
			mensaje.Respuesta = "No se pudo realizar el registro de la poliza"
		}

		if irespuesta == 1 {
			mensaje.StatusRespuesta = 1
		} else {
			mensaje.StatusRespuesta = 0
		}

		log.Printf("Termina PolizasService::ActualizarPoliza")
	}
	return mensaje

}
