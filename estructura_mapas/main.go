package main

import "fmt"

func main() {

	//No respeta el orden
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	//Acceder
	fmt.Println(colors["rojo"])

	//Agregar
	colors["negro"] = "#000000"

	fmt.Println(colors)

	valor := colors["rojo"]
	fmt.Println(valor)

	//Verificacion .- Existe?
	v, ok := colors["rojo"]
	fmt.Println(v, ok)

	//Verificacion .- Existe?
	z, ok := colors["blanco"]
	fmt.Println(z, ok)
	//Si no existe, devuelve una cadena vaciay false

	if ok {
		fmt.Println("Si existe la clave")
	} else {
		fmt.Println("No existe la clave")
	}

	//Otra forma de hacerlo
	if v, ok := colors["rojo"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("No existe esta clave")
	}

	//Eliminar
	delete(colors, "azul")
	fmt.Println(colors)

	//Iterar / recorrer
	for k, v := range colors {
		fmt.Printf("Clave: %s, Valor: %s\n", k, v)
	}

}
