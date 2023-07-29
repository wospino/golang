package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

//var firtsName, lastName, age = "William", "Ospino", 41

// firtsName string = "William"
// lastName  string = "Ospino"
// age       int    = 41
// declaracion de constantes
const Pi float32 = 3.14

const (
	X = 100
	Y = 0b1010
	Z = 0o12
	W = 0xFF
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {

	fmt.Println("hola mundo")
	fmt.Println(quote.Hello())

	//declaracion de variable
	firtsName, lastName, age := "William", "Ospino", 41
	fmt.Println("primer nombre ", firtsName, " Segundo nombre ", lastName, " edad ", age, "valor de Pi ", Pi)

	fmt.Println(X, Y, Z, W)

	fmt.Println(Viernes)

	fmt.Println(math.MinInt64, math.MaxInt64)

	var (
		defaultInt     int
		defaultUint    uint
		defaultFloat32 float32
		defaultBool    bool
		defaultString  string
	)

	fmt.Println(defaultInt, defaultUint, defaultFloat32, defaultBool, defaultString)

	//conversiones de tipos de dato:

	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(integer16 + int16(integer32))

	//string to int

	s := "100"
	i, _ := strconv.Atoi(s)
	fmt.Println(i + i)

	//int to string

	n := 42
	s = strconv.Itoa(n)

	fmt.Println(s)

	fmt.Print("otro mensaje \n")
	var name, lastName1 string
	var edad int

	//ingresando datos por consola, con fmt.Scan, escaneando una o mas variables.
	fmt.Println("ingrese su nombre y apellido")
	fmt.Scan(&name, &lastName1)
	fmt.Println("ingrese su edad")
	fmt.Scan(&edad)

	//se puede usar %v para cuando no se conoce el tipo de dato.
	// con %T se puede saber el tipo de dato que viene.
	fmt.Printf("hola, me llamo %s %s y tengo %d años.\n", name, lastName1, edad)

	greeting := fmt.Sprintf("hola, me llamo %s y tengo %d años.", name, edad)

	fmt.Println(greeting)
}
