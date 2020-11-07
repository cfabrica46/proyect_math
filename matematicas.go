package main

import (
	"fmt"

	"github.com/Cesar/prueba_1/Matematicas/mate"
)

func main() {

	var numerooperaciones int

	vector1 := []string{"suma", "0"}
	vector2 := []string{"resta", "0"}
	vector3 := []string{"multiplicación", "0"}
	vector4 := []string{"división", "0"}

	operaciones := [][]string{vector1, vector2, vector3, vector4}

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
		fmt.Printf("%d.%s\t", i+1, operaciones[i][0])
	}
	fmt.Println("")

	fmt.Println(mate.Suma(1, 2))
}
