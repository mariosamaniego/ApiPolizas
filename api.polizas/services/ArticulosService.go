package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"main.go/entities"
)

func ConsultarArticulos() entities.Articulos {
	log.Printf("Inicio ArticulosService::ConsultarArticulos")
	articulos := entities.Articulos{}

	db, err := sql.Open("postgres", urlPostgress)
	if err != nil {
		log.Printf("%s", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT sku,nombre,cantidad FROM fun_consultaarticulos()")
	if err != nil {
		fmt.Println("Error en la fun_consultaarticulos")
		log.Printf("Error en la fun_consultaarticulos")
		return entities.Articulos{}

	}
	defer rows.Close()

	for rows.Next() {
		articulo := entities.Articulo{}
		err := rows.Scan(&articulo.Sku, &articulo.Nombre, &articulo.Cantidad)
		articulos = append(articulos, articulo)

		if err != nil {
			log.Printf("%s", err)
			return articulos
		}
	}

	if err = rows.Err(); err != nil {
		log.Printf("%s", err)
		return articulos
	}

	log.Printf("Fin ArticulosService::ConsultarEmpleados")

	return articulos
}
