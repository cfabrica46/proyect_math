package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/cfabrica46/prueba_1/Matematicas/mate"
)

type operacion struct {
	nombre string
	acceso bool
	id     int
}

func main() {

	var numerooperaciones, eleccionoperaciones, scifras, rcifras, m1cifras, m2cifras, d1cifras, d2cifras, respuesta, puntosbuenos, puntosmalos int

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

					fmt.Println("¿De Cuantas cifras sera el primer factor?")
					fmt.Scan(&m1cifras)
					fmt.Println("¿Hasta que tabla seran las multiplicaciones?")
					fmt.Scan(&m2cifras)
				case 4:
					fmt.Println("¿De Cuantas cifras sera el primer factor?")
					fmt.Scan(&d1cifras)
					fmt.Println("¿Hasta que numero se dividirá?")
					fmt.Scan(&d2cifras)
				}

				fmt.Println("¿Desea agregar otra operacion?")
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

	}

	for i := range ejercicios {

		fmt.Printf("Ejercicio %d:\n", i+1)

		for index := range operaciones {

			if ejercicios[i] == operaciones[index].id {

				switch operaciones[index].id {
				case 0:

					sumando := convertircifras(scifras)

					operador1 := rand.Intn(sumando)
					operador2 := rand.Intn(sumando)

					fmt.Printf("%d + %d \n", operador1, operador2)
					fmt.Scan(&respuesta)

					if respuesta == mate.Sumar(operador1, operador2) {

						fmt.Println("Muy bien, la respuesta es correcta")

						puntosbuenos++

					} else {
						fmt.Println("Que mala suerte, esa no era la respuesta")
						puntosmalos++
					}

				case 1:

					sustraendo := convertircifras(rcifras)

					operador1 := rand.Intn(sustraendo)
					operador2 := rand.Intn(sustraendo)

					fmt.Printf("%d - %d \n", operador1, operador2)
					fmt.Scan(&respuesta)

					if respuesta == mate.Restar(operador1, operador2) {

						fmt.Println("Muy bien, la respuesta es correcta")

						puntosbuenos++

					} else {
						fmt.Println("Que mala suerte, esa no era la respuesta")
						puntosmalos++
					}

				case 2:

					mfactor1 := convertircifras(m1cifras)

					operador1 := rand.Intn(mfactor1)
					operador2 := rand.Intn(m2cifras)

					fmt.Printf("%d * %d \n", operador1, operador2)

					fmt.Scan(&respuesta)

					if respuesta == mate.Multiplicar(operador1, operador2) {

						fmt.Println("Muy bien, la respuesta es correcta")

						puntosbuenos++

					} else {
						fmt.Println("Que mala suerte, esa no era la respuesta")
						puntosmalos++
					}

				case 3:

					dfactor1 := convertircifras(d1cifras)

					operador1 := rand.Intn(dfactor1)
					operador2 := rand.Intn(d2cifras)

					fmt.Printf("%d ÷ %d \n", operador1, operador2)
					fmt.Scan(&respuesta)

					if respuesta == mate.Dividir(operador1, operador2) {

						fmt.Println("Muy bien, la respuesta es correcta")

						puntosbuenos++

					} else {
						fmt.Println("Que mala suerte, esa no era la respuesta")
						puntosmalos++
					}
				}

			}

		}

	}

	fmt.Println("Estos son los resultados:")

	fmt.Printf("%d puntos a favor\n", puntosbuenos)
	fmt.Printf("%d puntos en contra\n", puntosmalos)

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

func convertircifras(cifras int) (r int) {

	var buffer bytes.Buffer

	for i := 0; i < cifras; i++ {

		buffer.WriteString("9")

	}

	r, err := strconv.Atoi(buffer.String())

	if err != nil {
		log.Fatal(err)
	}
	return
}
