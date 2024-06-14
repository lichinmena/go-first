package main

import "library/book"

//En go no existen constructures, pero se simulan

//Un atributo declarado con minuscula es PRIVADO
//Un attributo declarado con mayuscula es PUBLICO

//La compocision en GO permite crear nuevas estructuras
//que contienen campos y metodos de otras estructuras.

func main() {
	var myBook = book.Book{
		Title:  "Mody Dick",
		Author: "Herman Melville",
		Pages:  300,
	}

	myBook2 := book.NewBook("1", "2", 300)

	composicion := book.NewTexBook("1", "2", 1, "4", "5")

	myBook.PrintInfo()
	myBook2.PrintInfo()
	composicion.PrintInfo()

}
