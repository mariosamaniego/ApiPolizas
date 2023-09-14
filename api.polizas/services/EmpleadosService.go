package services

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"main.go/entities"
)

func ConsultarEmpleados() entities.Empleados {
	log.Printf("Inicio EmpleadosService::ConsultarEmpleados")

	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("Error conexion", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT idempleado,nombre,apellido,puesto FROM fun_consultaempleados()")
	if err != nil {
		log.Printf("Error en la fun_consultaempleados")
		return entities.Empleados{}

	}
	defer rows.Close()
	empleados := entities.Empleados{}

	for rows.Next() {
		empleado := entities.Empleado{}
		err := rows.Scan(&empleado.IdEmpleado, &empleado.Nombre, &empleado.Apellido, &empleado.Puesto)
		empleados = append(empleados, empleado)

		if err != nil {
			log.Printf("", err)
			return empleados
		}

	}

	if err = rows.Err(); err != nil {
		log.Printf("", err)
		return empleados
	}
	log.Printf("Fin EmpleadosService::ConsultarEmpleados")

	return empleados

}

func AgregarEmpleado(idempleado int, nombre string, apellido string, puesto string) entities.Mensaje {
	log.Printf("Inicio EmpleadosService::AgregarEmpleado")
	var irespuesta int32
	mensaje := entities.Mensaje{}
	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT fun_registraempleado FROM fun_registraempleado($1::INTEGER,$2,$3,$4)", idempleado, nombre, apellido, puesto)
	if err != nil {
		log.Printf("Error al ejecutar la funcion fun_registraempleado()")
		log.Printf("", err)

	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&irespuesta)
		if err != nil {
			log.Printf("Error al escanear los resultados de la funcion fun_registraempleado()")
			log.Printf("%s", err)

		}

		if irespuesta == 0 {
			log.Printf("Error al registrar al empleado,Empleado ya existe")
			mensaje.Respuesta = "No se ha podido registrar Empleado"
			mensaje.StatusRespuesta = 0
		} else {
			mensaje.Respuesta = "Registro Exitoso"
			mensaje.StatusRespuesta = 1
		}

		log.Printf("Termina EmpleadosService::AgregarEmpleado")

	}
	return mensaje

}

func EliminarGeneral(opcion int, eliminar int) entities.Mensaje {
	log.Printf("Inicio EmpleadosService::EliminarGeneral")

	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT fun_eliminargeneral FROM fun_eliminargeneral($1::INTEGER,$2::INTEGER)", opcion, eliminar)
	if err != nil {
		log.Printf("Error al ejecutar la funcion fun_eliminargeneral()")
		log.Printf("", err)

	}
	defer rows.Close()
	var irespuesta int32
	mensaje := entities.Mensaje{}

	if rows.Next() {

		err := rows.Scan(&irespuesta)

		if err != nil {
			log.Printf("Error al escanear los resultados de la funcion fun_eliminargeneral()")
			log.Printf("", err)

		}

		if irespuesta == 1 {
			mensaje.Respuesta = "Eliminacion Exitosa"
		} else if irespuesta == 2 {
			mensaje.Respuesta = "El empleado tiene polizas activas, Imposible eliminar"
		} else {
			mensaje.Respuesta = "Error al Eliminar Registro"
		}

		log.Printf("Termina EmpleadosService::EliminarGeneral")

	}
	return mensaje

}
