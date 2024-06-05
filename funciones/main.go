package main

import "fmt"

func main() {
	hello()
	helloV2("Luis")
	saludo := helloV3("Fernando")
	fmt.Println(saludo)
	sum1 := calcV1(4, 6)
	sum2, mul2 := calcV2(4, 6)
	sum3, mul3 := calcV3(4, 6)
	println(sum1)
	println(sum2)
	println(mul2)
	println(sum3)
	println(mul3)
}

func hello() {
	fmt.Println("Hola desde la funcion")
}

func helloV2(name string) {
	fmt.Println("Hola", name)
}

func helloV3(name string) string {
	return "Hola, " + name
}

// retorna un dato
func calcV1(a, b int) int {
	return a + b
}

// retorna dos datos
func calcV2(a, b int) (int, int) {
	suma := a + b
	mul := a * b
	return suma, mul
}

// retorna dos datos, con declaracion en el tipo de dato retornado
func calcV3(a, b int) (suma, mul int) {
	//Se elimina los :
	suma = a + b
	mul = a * b
	return
}
