package main

import (
	"log"
	"os"
)

func main() {
	log.Print("Este es un mensaje de registro")
	log.Println("Este es otro mensaje de registro")

	//Detiene el programa
	//log.Fatal("Oye, soy un registro de errores")

	//log.Panic("Oye soy un panic")

	//Configurar la clase
	log.SetPrefix("main(): ")
	log.Print("Soy un log")
	//log.Fatal("Soy un log fatal")

	//Enviar registros en un archivo
	//Registrar errores en un archivo

	//Manejo de archivos
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	//Escribir log en el archivo
	log.SetOutput(file)
	log.Print("!Oye soy un log!!")
}
