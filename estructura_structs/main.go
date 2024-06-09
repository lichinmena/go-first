package main

import "fmt"

//

// Creacion
type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	//Instanciar
	var p Persona
	p.nombre = "Luis"
	p.edad = 35
	p.correo = "test@outlook.com"
	fmt.Println(p)

	//Otra forma de instanciar
	p2 := Persona{"Fernando", 65, "asas@outlook.com"}
	fmt.Println(p2)

	fmt.Println(p.nombre)
	fmt.Println(p.edad)
	fmt.Println(p.correo)
}
