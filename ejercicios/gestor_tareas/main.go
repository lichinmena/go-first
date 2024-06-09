package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

// Agregar tareas
func (l *ListaTareas) agregar(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// Marcar completado
func (l *ListaTareas) completado(index int) {
	l.tareas[index].completado = true
}

// Metodo para editar tarea
func (l *ListaTareas) editar(index int, t Tarea) {
	l.tareas[index] = t
}

// Eliminar
func (l *ListaTareas) eliminar(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[+1:]...)
}

func main() {
	lista := ListaTareas{}

	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println("Seleccione una opcion: \n",
			"1.- Agregar tarea\n",
			"2.- Marcar tarea como completada \n",
			"3.- Editar tarea \n",
			"4.- Eliminar tarea \n",
			"5.- Salir")

		fmt.Print("Ingrese la opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')

			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')

			lista.agregar(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)
			lista.completado(index)
			fmt.Println("Tarea marcada como completada correctamente.")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea que desea actualizar: ")
			fmt.Scanln(&index)

			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')

			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')

			lista.editar(index, t)
			fmt.Print("Tarea actualizada correctamente.")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			lista.eliminar(index)
			fmt.Print("Tarea eliminada correctamente.")
		case 5:
			fmt.Println("Saliendo del programa ...")
			return
		default:
			fmt.Println("Opcion invalida")

		}

		//Listar todas las tareas
		fmt.Println("Lista de tareas")
		fmt.Println("=================================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("=================================")
	}
}
