package main

import (
	"fmt"

	"github.com/Cesar/prueba_1/Matematicas/mate"
)

func main() {

	var numerooperaciones int

	fmt.Println("Bienvenido")

	for numerooperaciones < 1 {

		fmt.Println("¿Cuantas operaciones deseas realizar")
		fmt.Scan(&numerooperaciones)
		if numerooperaciones < 1 {
			fmt.Println("Hey ese no es un numero valido >:v")
		}
	}

	fmt.Println(mate.Suma(1, 2))

}
