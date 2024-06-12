package main

import (
	"fmt"
	"log"

	"github.com/lichinmena/greetings"
)

//go mod init uso_modulo (sin haber creado la main.go)
//Como aun esta en local se realiza:
//go mod edit -replace github.com/lichinmena/greetings=../../greetings
//Finalmente para ejecutar los cambios
//go mod tidy

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0) //Establecer de formato en cero. NO se muestra la fecha y la hora

	names := []string{"Alex", "Roel", "Juan"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	/*
		message, err := greetings.Hello("Luis")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(message)
	*/
}
