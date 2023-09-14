package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"main.go/entities"
	"main.go/handlers"
)

func main() {
	file, err := os.OpenFile("logPolizas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%s Fatal: Fatal Error Signal")
	}
	defer file.Close()
	log.SetOutput(file)

	log.Printf("Inicio del api.polizas")

	mux := mux.NewRouter()
	mux.Use(CORS)
	mux.Use(Token)
	path := "/api/v1"
	mux.HandleFunc(path+"/Empleados", handlers.ConsultarEmpleados).Methods("GET", "OPTIONS")
	mux.HandleFunc(path+"/Articulos", handlers.ConsultarArticulos).Methods("GET", "OPTIONS")
	mux.HandleFunc(path+"/AgregarEmpleado", handlers.AgregarEmpleado).Methods("POST", "OPTIONS")
	mux.HandleFunc(path+"/ConsultarPoliza", handlers.ConsultarPoliza).Methods("GET", "OPTIONS")
	mux.HandleFunc(path+"/PolizaEmpelado", handlers.ConsultarPolizasEmpleado).Methods("GET", "OPTIONS")
	mux.HandleFunc(path+"/AgregarPolizas", handlers.AgregarPoliza).Methods("POST", "OPTIONS")
	mux.HandleFunc(path+"/Eliminar", handlers.EliminarGeneral).Methods("POST", "OPTIONS")
	mux.HandleFunc(path+"/ActualizarPoliza", handlers.ActualizarPoliza).Methods("POST", "OPTIONS")

	log.Printf("%s", http.ListenAndServe("10.59.21.109:3000", mux))
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Inicio Main::CORS")
		(w).Header().Set("Content-Type", "application/json")
		(w).Header().Set("Content-Type", "text/html; charset=utf-8")
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "GET, HEAD, OPTIONS, POST, PUT, DELETE")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		fmt.Println("cors")

		next.ServeHTTP(w, r)
		log.Printf("Finaliza Main::CORS")
	})
}

func Token(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Inicio Main::Token")

		var responseToken entities.Response

		var urlsso = os.Getenv("sso")
		method := "POST"
		payload := strings.NewReader(``)

		req, err := http.NewRequest(method, urlsso, payload)
		if err != nil {
			fmt.Println(err)
			log.Printf("%s", err)

		}

		req.Header.Add("Authorization", r.Header.Get("Authorization"))

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			log.Printf("%s", err)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error al leer el cuerpo de la respuesta:")
			log.Printf("%s", err)
		}

		var response entities.ResponseSSO
		json.Unmarshal(body, &response)

		responseToken.Meta.Status = response.Meta.Status
		responseToken.Data = response.Meta.Error

		status, _ := json.Marshal(responseToken)

		if responseToken.Meta.Status == "FAIL" {
			fmt.Fprintf(w, string(status))
			log.Printf("%s", err)
			return

		}
		next.ServeHTTP(w, r)
		log.Printf("Termina Main::Token")
		return
	})
}
