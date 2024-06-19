package main

import "fmt"

//Funciones variadicas
//Son funciones que pueden n argumentos

//Funciones recursivas
//Funciones que se ejecutan asi misma

//Funciones anonimas
//Funciones que no tienen un nombre o identificador

//Funciones de orden superior

//closures
//Una funcion anonima esta definida dentro de otra y puede recordar
//variables de una funcion exterior

// Funcion variadica
// De ultimo se pone el parametro variadico
func suma(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Hola %s, la suma es: %d\n", name, total)
	//fmt.Printf("%T - %v\n", nums, nums)
	return total
}

// Recibe n cantidad de datos y tipos (Variadica)
func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

// Funciones recursivas========>
// Una funcion asi misma
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Funciones anominas
func saludar(name string, f func(string)) {
	f(name)
}

// Normal
func duplicar(n int) int {
	return n * 2
}

// Normal
func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

//Orden superior
//Es aquella funcion que recibe una funcion como argunmento
//y retorna una funcion.

func double(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}

//Closures
//Una funcion anomina que se define dentro de una funcion
//y que puede acceder, modificar variables definidos en el ambito
//de la funcion externa. Recuerda el valor de la variable del ambito de la
//funcion, y utilizarlo en cualquier momento posterior,

func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(suma("Luis", 1, 2, 3, 4, 5, 6, 7, 8, 9))

	imprimirDatos("Hola", 28, true, 3.5)

	fmt.Println("=========>")

	fmt.Println(factorial(5))

	fmt.Println("=========>")
	//Funcion anomina
	func() {
		fmt.Println("Hola soy una funcion anomina")
	}() //Los parentesis permite que se ejecute

	//Se crea una funcion anomina, y se asigna a una variable
	//La variable recibe el parametro
	saludo := func(name string) {
		fmt.Printf("Hola, %s !\n", name)
	}

	saludo("Alex")

	//Metodo que recibe un string y una funcion
	saludar("Alex", saludo)

	//Ejcutar la funcion aplicar FUNCIONES COMO VALOR
	r1 := aplicar(duplicar, 5)
	r2 := aplicar(triplicar, 5)
	fmt.Println(r1, r2)

	fmt.Println("=========>")

	//Orden superior
	r := double(addOne, 3) //Primero se ejecuta la funcion double y luego addOne
	fmt.Println("Resultado: ", r)

	//clousures
	fmt.Println("=========>")
	nextInt := incrementer()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
