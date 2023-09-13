package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	entities "main.go/entities"
	"net/http"
	"os"
	"strings"
)

var urlsso = os.Getenv("sso")

func ValidarToken(token string) entities.ResponseSSO {

	method := "POST"
	payload := strings.NewReader(``)

	// Crear una nueva solicitud HTTP POST
	req, err := http.NewRequest(method, urlsso, payload)
	if err != nil {
		fmt.Println(err)

	}

	req.Header.Add("Authorization", token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body2, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error al leer el cuerpo de la respuesta:", err)
	}

	var response entities.ResponseSSO
	json.Unmarshal(body2, &response)

	status, _ := json.Marshal(response.Meta)

	fmt.Printf(string(status))

	return response
}
