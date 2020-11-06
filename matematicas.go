package main

import (
	"fmt"

	"github.com/Cesar/prueba_1/Matematicas/mate"
)

func main() {

	var numerooperaciones int

	operaciones := make(map[string]int, 4)
	operaciones["suma"] = 0
	operaciones["resta"] = 0
	operaciones["multiplicación"] = 0
	operaciones["división"] = 0

	fmt.Println(operaciones)
	return

	//Introducción...
	fmt.Println("Bienvenido")

	for numerooperaciones < 1 {

		fmt.Println("¿Cuantas operaciones deseas realizar")
		fmt.Scan(&numerooperaciones)
		if numerooperaciones < 1 {
			fmt.Println("Hey ese no es un numero valido >:v")
		}
	}

	fmt.Println("Muy bien")

	fmt.Println("¿Que ejercicios deseas realizar?")

	for i := 0; i < len(operaciones); i++ {
		fmt.Printf("%d.%s")
	}

	fmt.Println(mate.Suma(1, 2))
}
