package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Cesar/prueba_1/Matematicas/mate"
)

type operacion struct {
	nombre string
	acceso bool
	id     int
}

func main() {

	var numerooperaciones, eleccionoperaciones, scifras, rcifras, mcifras, dcifras int

	var buffer bytes.Buffer

	operaciones := []operacion{
		operacion{
			nombre: "suma",
			acceso: false,
			id:     0,
		},
		operacion{
			nombre: "resta",
			acceso: false,
			id:     1,
		},
		operacion{
			nombre: "multiplicación",
			acceso: false,
			id:     2,
		},
		operacion{
			nombre: "división",
			acceso: false,
			id:     3,
		},
	}

	elecciones := []int{}

	ejercicios := []int{}

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

			if len(elecciones) == 0 {
				fmt.Println("Elige alguna opción")
			} else {
				break
			}

		}

		if eleccionoperaciones <= len(operaciones) && eleccionoperaciones > 0 {

			operaciones[eleccionoperaciones-1].acceso = true

			if revisar(elecciones, eleccionoperaciones) == true {

				fmt.Println("Esa opción ya fue seleccionada")

			} else {

				elecciones = append(elecciones, eleccionoperaciones)

				switch eleccionoperaciones {
				case 1:
					fmt.Println("¿De Cuantas cifras sera la suma?")
					fmt.Scan(&scifras)
				case 2:
					fmt.Println("¿De Cuantas cifras sera la Resta?")
					fmt.Scan(&rcifras)
				case 3:
					fmt.Println("¿Hasta que tabla seran las multiplicaciones?")
					fmt.Scan(&mcifras)
				case 4:
					fmt.Println("¿Hasta que numero se dividirá?")
					fmt.Scan(&dcifras)
				}

			}

			if len(elecciones) == len(operaciones) {
				break
			}

		} else {

			if eleccionoperaciones == 0 {

			} else {
				fmt.Println("Hey, esa no era una opción")
			}

		}

	}

	fmt.Println("Seleccionaste:")

	for i, v := range elecciones {

		fmt.Printf("%d.%s\t", i+1, operaciones[v-1].nombre)

	}
	for i := 0; i < 2; i++ {
		fmt.Println("")

	}

	//Cuerpo del codigo...

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numerooperaciones; i++ {

		a := rand.Intn(len(elecciones))

		ejercicios = append(ejercicios, a)

		fmt.Println(ejercicios)

	}

	for i := range ejercicios {

		for index := range operaciones {

			if ejercicios[i] == operaciones[index].id {

				switch operaciones[index].id {
				case 0:

					for i := 0; i < scifras; i++ {

						buffer.WriteString("9")

					}

					sumando, err := strconv.Atoi(buffer.String())

					if err != nil {
						log.Fatal(err)
					}

					operador1 = rand.Intn(sumando)
					operador2 := rand.Intn(sumando)

				case 1:
					operador1 := rand.Intn(rcifras)
					operador2 := rand.Intn(rcifras)
				case 2:
					operador1 := rand.Intn(scifras)
					operador2 := rand.Intn(mcifras)
				case 3:
					operador1 := rand.Intn(scifras)
					operador2 := rand.Intn(dcifras)
				}

			}

		}

	}

	fmt.Println(mate.Suma(1, 2))

}

func revisar(v []int, n int) (aviso bool) {

	for i := 0; i < len(v); i++ {

		for i := 0; i < len(v); i++ {

			if v[i] == n {

				aviso = true

				return
			}

		}

	}

	return

}
