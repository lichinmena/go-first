package main

import "fmt"

//Son usados como receptores
//Almacena la direccion de memoria
//Referencia y acceder a la variable original a traves
//de la direccion de memoria

type Persona struct {
	nombre string
	edad   int
	correo string
}

// Crea un metodo del struct persona
// Recibe parametro como puntero
func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es: ", p.nombre)
}

func main() {
	var x int = 10
	var p *int = &x // Mandar la referencia de la memoria de x a p

	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(&x)

	editar(&x) //Envia un puntero

	fmt.Println(x) //Imprime el valor de la variable

	pp := Persona{"Alex", 28, "asas@outlook.com"}
	pp.diHola()
}

// Recibe un puntero
func editar(x *int) {
	*x = 20
}
