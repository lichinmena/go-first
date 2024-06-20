package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

//Desde la version 1.18 se puede usar ANY

// Forma antigua con interface
func PrintList(list ...interface{}) {
	for _, value := range list {
		fmt.Println(value)
	}
}

// Forma nueva con any
func PrintList2(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

// Restricciones en las funciones genericas
// Funcion variadica
func Sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

// Funcion genericas para evitar duplicidad de codigo
// Se usa parametros de tipos, son marcadores que se usan en la definicion de
// una funcion generica que representa un tipo desconocido
// Restriccion arbitrario.
// Union de elementos
// La funcion puede trabajar con int y con float64
// Restriccion de aproximacion se coloca el simbolo ~
//Tipos de restricciones
//1.- Restricciones arbitrarias    => T int
//2.- Restricciones de union       => T int | float64
//3.- Restriccion de aproximacion  => T ~int | ~float64
/*
func Sum2[T int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}
*/
func Sum2[T ~int | ~float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

type integer int

// Para crear restricciones
type Numbers interface {
	~int | ~float64 | ~float32
}

// Usa la restriccion creada Numbers
func Sum3[T Numbers](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// COn la utilizacion del paquete instalado
func Sum4[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Funcion con restriccion comparacion
func Includes[T comparable](list []T, value T) bool {
	for _, row := range list {
		if row == value {
			return true
		}
	}
	return false
}

// Funcion con restriccion filtrado
func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

// Estructuras genericas
type Product struct {
	Id    uint
	Desc  string
	Price float32
}

// Estructura generica
type ProductV2[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {
	PrintList("Alex", "Roel", "Juan", "Pedro")
	PrintList(1, 2, 3, 4, 5, 6, 7)
	PrintList(1, 2.2, "Hola", true)
	fmt.Println("")
	PrintList2("Alex", "Roel", "Juan", "Pedro")
	PrintList2(1, 2, 3, 4, 5, 6, 7)
	PrintList2(1, 2.2, "Hola", true)
	fmt.Println("")
	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println("")
	fmt.Println(Sum2(1, 2, 3, 4, 5))
	fmt.Println(Sum2(13.11, 21.11, 3.22, 4.33, 5.44))
	fmt.Println("")
	var num1 integer = 100
	var num2 integer = 300
	fmt.Println(Sum2(num1, num2)) //Usa la restriccion por aproximacion
	fmt.Println("")
	fmt.Println(Sum3(num1, num2))
	fmt.Println("")
	fmt.Println(Sum4[float32](3.4, .6))
	fmt.Println(Sum4[integer](2, 3))
	fmt.Println("")
	strings := []string{"a", "b", "c"}
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(Includes(strings, "a"))
	fmt.Println(Includes(strings, "f"))

	fmt.Println(Includes(numbers, 1))
	fmt.Println(Includes(numbers, 6))

	fmt.Println("")
	fmt.Println(Filter(numbers, func(value int) bool { return value > 3 }))
	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))

	//Creacion del producto
	product1 := Product{1, "Zapatos", 50}
	fmt.Println(product1)
	//Creacion del producto con estrutura generica
	product2 := ProductV2[uint]{1, "Zapatos", 50}
	fmt.Println(product2)
	product3 := ProductV2[string]{"JUSJS", "Zapatos", 50}
	fmt.Println(product3)

}
