package main

import (
	"fmt"
)

func main() {
	/*var a [5]int
	a[0] = 10
	a[1] = 20
	a[2] = 30

	//inicializando una matriz
	//var b = [...]int{10, 20, 30, 40, 50}

	//fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	//recorriendo de otra forma
	for index, value := range a {
		fmt.Printf("indice : %d, Valor %d\n", index, value)
	}

	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz)*/
	rebanadas()

}

func rebanadas() {
	//var a [] int
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	diaRebanada := diasSemana[0:5]

	//Agregando elementos
	diaRebanada = append(diaRebanada, "Viernes", "Sabado", "Domingo")
	fmt.Println(diaRebanada)
	//quitando el segundo elemento
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	//funcion make

	nombres := make([]string, 5, 10)
	nombres[0] = "William"
	fmt.Println(nombres)

	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)
	//funcion copy
	println(copy(rebanada2, rebanada1))
	fmt.Println(rebanada1, rebanada2)
}
