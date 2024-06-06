package main

import "fmt"

//Longitud fija
//Inmutable

func main() {
	//Declarar matriz
	var matriz [5]int

	//modificar elemento de matriz
	matriz[0] = 10

	fmt.Println(matriz)

	//Declarar e inializar
	var a = [5]int{10, 20, 30, 40, 50}
	fmt.Println(a)

	//Declarar e inializar cuando no sabemos cuantos elementos tendra
	var b = [...]int{10, 20, 30, 40, 50}
	fmt.Println(b)

	//Acceder al elemento
	fmt.Println(b[3])

	//Recorrer
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	//Otra forma de iterar
	for index, value := range b {
		fmt.Printf("Indice: %d, Valor: %d\n", index, value)
	}

	//Matriz bidimensional
	var c = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(c)

}
