package main

import (
	"fmt"
	/*con el manejador de modulos, en la raiz del proyecto con la consola
	ejecutamos go mod init holamundo lo que genera un modulo en el
	proyecto y dos archivos, go.mod y go.sum */
	"rsc.io/quote"
)

func main() {
	fmt.Println("hola mundo")
	//Utilizando un paquete de un tercero que se adiciona por medio del
	//manejador de modulos de go
	fmt.Println(quote.Hello())
}
