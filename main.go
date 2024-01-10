package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	// Se especifica la url de la que se va a extraer datos
	url := "https://blog.desafiolatam.com/15-proyectos-para-desarrollar-en-go-y-su-importancia-ti/"

	// Se realiza una solicitud Http a la url
	resp, err := http.Get(url)
	if err != nil {
		//Si da error en la respuesta va a tirar mensaje de eeror y parar
		fmt.Println("Error al realizar la solicitud: ", err)
		return
	}

	// el defer asegura que se cierre la conexion con el servidor
	// no cerrarlo puede ocasionar problemas de recursos para el servidor
	defer resp.Body.Close()

	//Analiza el cuerpo de la respuesta y lo convierte a un arbol DOM
	doc, err := html.Parse(resp.Body)
	if err != nil {
		//Si da error en el analisis tira error y se detiene
		fmt.Println("Error al analizar HTML: ", err)
		return
	}

	//llama a ExtraerLinks con el documento HTML para extraer los enlaces
	links := ExtraerLinks(doc)
	for _, link := range links {
		// Se imprime los links dependiendo de la longitud de la lista de links
		fmt.Println(link)
	}

}
