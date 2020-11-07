package main

import (
	"fmt"

	"github.com/Cesar/prueba_1/Matematicas/mate"
)

type operacion struct {
	nombre string
	acceso bool
}

func main() {

	var numerooperaciones, eleccionoperaciones int

	operaciones := []operacion{
		operacion{
			nombre: "suma",
			acceso: false,
		},
		operacion{
			nombre: "resta",
			acceso: false,
		},
		operacion{
			nombre: "multiplicación",
			acceso: false,
		},
		operacion{
			nombre: "división",
			acceso: false,
		},
	}

	elecciones := []int{}

	//Introducción...
	fmt.Println("Bienvenido")

	for numerooperaciones < 1 {

		fmt.Println("¿Cuantas operaciones deseas realizar?")
		fmt.Scan(&numerooperaciones)
		if numerooperaciones < 1 {
			fmt.Println("Hey ese no es un numero valido >:v")
		}
	}

	fmt.Println("Muy bien")

	fmt.Println("¿Que ejercicios deseas realizar?")

	for i := 0; i < len(operaciones); i++ {
		fmt.Printf("%d.%s\t", i+1, operaciones[i].nombre)
	}
	fmt.Println("0.Listo")

	for {
		fmt.Scan(&eleccionoperaciones)

		if eleccionoperaciones == 0 {
			break
		}

		if eleccionoperaciones <= len(operaciones) {

			operaciones[eleccionoperaciones-1].acceso = true

			elecciones = append(elecciones, eleccionoperaciones)

			if len(elecciones) != 1 {

				if revisar(elecciones) == true {

					elecciones = append(elecciones[:len(elecciones)-1])

				}

			}

			fmt.Println(elecciones)

			if len(elecciones) == len(operaciones) {
				break
			}

		} else {
			fmt.Println("Hey, esa no era una opción")
		}

	}

	fmt.Println(operaciones)

	fmt.Println(mate.Suma(1, 2))
}

func revisar(vector []int) (aviso bool) {

	for i := 0; i < len(vector)-1; i++ {

		for i := 0; i < len(vector)-1; i++ {

			if vector[i] == vector[i+1] {

				aviso = true

				return
			}

		}

	}

	return

}
