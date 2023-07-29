package main

import (
	"fmt"
	"math"
)

func main() {
	var base float64
	var altura float64
	var area float64
	var perimetro float64
	var hipotenusa float64
	fmt.Println("Digite la base del triangulo Rectangulo")
	fmt.Scan(&base)
	fmt.Println("Digite la altura del triangulo Rectangulo")
	fmt.Scan(&altura)

	area = (base * altura) / 2
	hipotenusa = math.Sqrt(math.Pow(base, 2) + math.Pow(altura, 2))
	perimetro = area + base + hipotenusa

	fmt.Println("El area del triangulo Rectangulo es -> ", area)
	fmt.Println("El perimetro del triangulo rectangulo es -> ", perimetro)

}
