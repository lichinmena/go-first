package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json: "name"`
	Email string `json: "email"`
	Phone string `json: "phone"`
}

func saveContactToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func loadContacstFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)

	if err != nil {
		return err
	}
	return nil
}

func main() {
	//Slice de contactos
	var contacts []Contact

	//Cargar archivos existentes
	err := loadContacstFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}

	//CRear instancia de fubio
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("====== Gestor de contactos ======\n",
			"1.- Agregar un contacto\n",
			"2.- Mostrar todos los contactos\n",
			"3.- Salir\n",
			"Elige una opcion: ")
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion:", err)
			return
		}

		//Manejar la opcion del usuario
		switch option {
		case 1:
			var c Contact

			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')

			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')

			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			//Agregar un contacto a slice
			contacts = append(contacts, c)

			//Guardar en un archivo json
			if err := saveContactToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto:", err)
			}

		case 2:
			fmt.Println("====================")

			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("====================")
		case 3:
			//salir del programa
			return
		default:
			fmt.Println("Opcion no valida")
		}
	}

}
