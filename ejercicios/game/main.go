package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(100))
	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 3

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingrese el numero (Intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("Felicitaciones, adivinaste el numero !!")
			jugarDeNuevo()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El numero a divinar es mayor!!")
		} else if numIngresado > numAleatorio {
			fmt.Println("El numero a divinar es menor!!")

		}
	}
	fmt.Println("Se acabaron los intentos. El numero era: ", numAleatorio)
	jugarDeNuevo()
}

func jugarDeNuevo() {
	var eleccion string
	fmt.Println("Quieres jugar nuevamente (s/n): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("!!! Gracias por jugar")
	default:
		fmt.Println("Eleccion incorrecta. Intentalo nuevamente!!!!!")
		jugarDeNuevo()
	}
}
