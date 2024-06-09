package main

import (
	"fmt"
	"os"
)

//Se utiliza para posponer la ejecucion
//de una funcion, hasta la funcion que la contiene
//Halla finalido

//¿Para que se usa, que beneficio?

func main() {
	/*
		defer fmt.Println(3) //Se ejecutara de ultimo por defer
		fmt.Println(1)
		fmt.Println(2)
	*/

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close() //Se va a ejecutar antes de que finalize la funcion main

	_, err = file.Write([]byte("Hola, Luis Mena"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
