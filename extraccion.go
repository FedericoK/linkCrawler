package main

import (
	"golang.org/x/net/html"
)

// La funcion recibe un nodo de html y devuelve una lista de string
func ExtraerLinks(n *html.Node) []string {

	//Se crea la variable para almacenar la lista(Slice)
	var links []string

	//Verifica si el nodo actual es un elemento de tipo "<a>" que es la etiqueta de links
	if n.Type == html.ElementNode && n.Data == "a" {

		//si es <a>, recorre sus atributos para encontrar 'href' que es la etiqueta con la direccion
		//.Attr es un atributo de la Estructura Node el mismo es un tipo []Attribute que tiene Key(nombre del atributo html) y Val(valor del atributo)
		for _, a := range n.Attr {
			if a.Key == "href" {
				// si a.key(atributo) en un href agrega su valor a la lista
				links = append(links, a.Val)
				//ya agrego el link de este nodo entonces para ya el for
				break
			}
		}
	}

	//el bucle comienza por el primer hijo del nodo n.FirstChild
	//el bucle seguira mientras tenga hijos todavia el nodo c != nil, cuando no tenga mas hijos sera 0 y se detendra
	//cuando termine con un bucle asumira el valor delsiguiente hijo c = c.NextSibling
	for c := n.FirstChild; c != nil; c = c.NextSibling {

		//recursividad de ExtraerLinks, para poder ver si los hijos no tienen mas hijos con links
		//el operador '...' se utiliza para agregar datos de un slice a otro
		links = append(links, ExtraerLinks(c)...)
	}

	//retorna la lista creada
	return links
}
