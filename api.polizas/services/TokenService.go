package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	entities "main.go/entities"
)

var urlsso = os.Getenv("sso")

func ValidarToken(token string) entities.ResponseSSO {

	method := "POST"
	payload := strings.NewReader(``)

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
