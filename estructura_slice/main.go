package main

import "fmt"

// Trabaja con partes de un arreglo
// Agrega y eliminar elementos de manera dinamica
// Mutable
func main() {
	var a []int
	fmt.Println(a)

	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles"}
	fmt.Println(diasSemana)

	//Crear una a partir de otra
	rebanada := diasSemana[0:2]
	fmt.Println(rebanada)

	//Capacidad
	fmt.Println(len(rebanada)) //Len
	fmt.Println(cap(rebanada)) //Capacidad Obtiene el tamaño del slice padre

	//Agregar elemento
	rebanada = append(rebanada, "Jueves", "Viernes", "Sabado")
	fmt.Println(rebanada)

	//Quitar
	//Eliminar el segundo elemento
	//Usa el indice
	rebanada = append(rebanada[:2], rebanada[3:]...)
	fmt.Println(rebanada)

	//Funcion make
	//Crea rebanada
	nombres := make([]string, 5, 10) //Tipo dato - Elementos a almacenar - Capcidad
	nombres[0] = "Luis"
	fmt.Println(nombres)

	//Copy, copiar elementos de un slice a otro slice
	r1 := []int{1, 2, 3, 4, 5}
	r2 := make([]int, 5) //Creas un slice con 5 pos - otra forma de crear

	fmt.Println(r1)
	fmt.Println(r2)

	copy(r2, r1)

	fmt.Println(r1)
	fmt.Println(r2)

}
