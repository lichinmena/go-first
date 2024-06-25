package main

import (
	"bufio"
	"db_mysql/controllers"
	"db_mysql/database"
	"db_mysql/models"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" //paquete que se usa de manera indirecta
)

func main() {
	/*
		dns := "VVVVVV:AAAAA@tcp(localhost:ss)/xx"

		db, err := sql.Open("mysql", dns)

		if err != nil {
			log.Fatal(err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}

		log.Println("---- Conexion exitosa ----")
	*/
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	/*
		controllers.ListContacts(db)
		fmt.Println("")
		controllers.GetById(db, 2)
		fmt.Println("")
		newContact := models.Contact{
			Name:  "Juan",
			Email: "juan@gmail.com",
			Phone: "9999907867",
		}
		controllers.Create(db, newContact)
	*/

	/*
		updateContact := models.Contact{
			Id:    4,
			Name:  "Juanelo",
			Email: "juanelo@gmail.com",
			Phone: "9999907867",
		}
		controllers.Update(db, updateContact)
	*/

	//controllers.Delete(db, 1)
	//fmt.Println("")

	//INteractividad
	for {
		fmt.Println("")
		fmt.Println("----> Menu <----")
		fmt.Println("1.- Listar contactos")
		fmt.Println("2.- Obtener contacto por ID")
		fmt.Println("3.- Crear nuevo contacto")
		fmt.Println("4.- Actualizar contact")
		fmt.Println("5.- Eliminar contacto")
		fmt.Println("6.- Salir")
		fmt.Println("Seleccione una opcion: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			controllers.ListContacts(db)
		case 2:
			fmt.Println("Ingrese el ID del contacto: ")
			var id int
			fmt.Scanln(&id)
			controllers.GetById(db, id)
		case 3:
			newContact := inputDetailsDetails(option)
			controllers.Create(db, newContact)
			controllers.ListContacts(db)
		case 4:
			updateContact := inputDetailsDetails(option)
			controllers.Update(db, updateContact)
			controllers.ListContacts(db)
		case 5:
			fmt.Println("Ingrese el ID del contacto que desea eliminar:")
			var id int
			fmt.Scanln(&id)
			controllers.Delete(db, id)
			controllers.ListContacts(db)
		case 6:
			fmt.Println("Saliendo del programa......")
		default:
			fmt.Println("--> Opcion incorrecta o no valida <--")
		}
	}
}

func inputDetailsDetails(option int) models.Contact {
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 4 {
		fmt.Println("Ingrese el ID del contacto que desea editar:")
		var id int
		fmt.Scanln(&id)
		contact.Id = id
	}

	fmt.Print("Ingrese el nombre del contacto:")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el email del contacto:")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el telefono del contacto:")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
