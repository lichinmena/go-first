package main

import (
	"fmt"

	"github.com/lichinmena/greetings"
)

//go mod init uso_modulo (sin haber creado la main.go)
//Como aun esta en local se realiza:
//go mod edit -replace github.com/lichinmena/greetings=../../greetings
//Finalmente para ejecutar los cambios
//go mod tidy

func main() {
	message := greetings.Hello("Luis")
	fmt.Println(message)
}
