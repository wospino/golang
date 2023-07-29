package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100)

	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("ingrese un numero (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scan(&numIngresado)
		if numIngresado == numAleatorio {
			fmt.Println("!feclicitaciones adivinaste el numero con %d intentos ", intentos)
			break
		} else if numIngresado > numAleatorio {
			fmt.Println("el numero a adivinar es menor ")
			continue
		} else if numIngresado < numAleatorio {
			fmt.Println("el numero ingresado es mayor ")
			continue
		}

	}
	fmt.Println("se acabaron los intentos, el numero aleatorio era ", numAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var eleccion string
	fmt.Printf("quieres jugar nuevamente? ")
	fmt.Scan(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("gracias por jugar")
	default:
		fmt.Println("Eleccion incorrecta")
		jugarNuevamente()
	}

}
