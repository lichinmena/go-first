package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

//En go no existen constructures, pero se simulan

//Un atributo declarado con minuscula es PRIVADO
//Un attributo declarado con mayuscula es PUBLICO

//La compocision en GO permite crear nuevas estructuras
//que contienen campos y metodos de otras estructuras.

//Polimorfismo
//Capacidad de los objetos de diferentes clases
//Para responder el mismo mensaje o metodo
//Se implementa por un interface

//Define un contrato de que metodos deben ser implementados

func main() {
	var myBook = book.Book{
		Title:  "Mody Dick",
		Author: "Herman Melville",
		Pages:  300,
	}

	myBook2 := book.NewBook("1", "2", 300)

	myTextBook := book.NewTextBook("1", "2", 1, "4", "5")

	myBook.PrintInfo()
	//myBook2.PrintInfo()
	//composicion.PrintInfo()

	//Probar la interface
	fmt.Println("===> Probando interface")
	//myBook2.PrintInfo()
	//myTextBook.PrintInfo()

	book.Print(myBook2)
	book.Print(myTextBook)

	//Animal
	miPerro := animal.Perro{Nombre: "Max"}
	miGato := animal.Gato{Nombre: "Tom"}

	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

	fmt.Println("")

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Tom"},
		&animal.Perro{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Luna"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}

}
