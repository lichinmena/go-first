package controllers

import (
	"database/sql"
	"db_mysql/models"
	"fmt"
	"log"
)

func ListContacts(db *sql.DB) {
	query := "SELECT * FROM contact"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("===> Lista de contactos <===")

	for rows.Next() {
		contact := models.Contact{}
		var valueEmail sql.NullString
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Sin correo electronico"
		}

		fmt.Println("--------------------------------------")

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
	}
	fmt.Println("--------------------------------------")

}

func GetById(db *sql.DB, contactID int) {
	query := "SELECT * FROM contact WHERE id = ?"
	row := db.QueryRow(query, contactID)
	contact := models.Contact{}
	var valueEmail sql.NullString

	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No se encontro registro con el id: %d", contactID)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "Sin correo electronico"
	}

	fmt.Println("Lista de un contacto")

	fmt.Println("--------------------------------------")

	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)

	fmt.Println("--------------------------------------")

}

func Create(db *sql.DB, contact models.Contact) {
	query := "insert into contact (name, email, phone) values(?, ?, ?)"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto registrado con exito")
}

func Update(db *sql.DB, contact models.Contact) {
	query := "update contact SET name = ?, email = ?, phone = ? WHERE id = ?"
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto actualizado con exito")
}

func Delete(db *sql.DB, id int) {
	query := "DELETE FROM contact WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto eliminado con exito")
}
