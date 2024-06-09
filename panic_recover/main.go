package main

import "fmt"

//Panic y recover se utiliza para manejar errores graves
//Controlar panico se hace con recover

//Defer se usa para que cuando se provoque un panic
//El sistema no caiga, y pueda continuar.
// Ya que defer lo que hace es ejecutarse al final

func divide(dividendo, divisor int) {

	//Funciones anonimasssssss
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	//validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

/*
func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir entre cero")
	}
}
*/

func main() {
	divide(100, 10)
	divide(200, 25)
	divide(34, 0)
	divide(100, 4)
}
