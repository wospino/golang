package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*if time.Now().Hour() < 12 {
		fmt.Println("Mañana")
	} else if time.Now().Hour() < 17 {
		fmt.Println("Tarde")
	} else {
		fmt.Println("Noche")
	}*/

	/*switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Mañana")
	case t.Hour() < 17:
		fmt.Println("Tarde")
	default:
		fmt.Println("Noche")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run on -> Windows")
	case "linux":
		fmt.Println("Go run on -> Linux")
	case "darwin":
		fmt.Println("Go run on -> macOS")
	default:
		fmt.Println("Go run on -> Other Os")
	}*/
	//var i int

	texto, entero := hello("William", 5)
	fmt.Println("El nombre enviado fue  " + texto + " repeticiones " + strconv.Itoa(entero))
}

func hello(name string, repeticion int) (nombre string, veces int) {
	for i := 1; i <= repeticion; i++ {
		/*if i == 5 {
			continue // break para salir del ciclo
		}*/
		fmt.Print(strconv.Itoa(i) + " " + name + "\n")
	}
	nombre = name
	veces = repeticion
	return
}
