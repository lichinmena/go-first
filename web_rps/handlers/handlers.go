package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		http.Error(w, "Error al analizar plantillas", http.StatusInternalServerError)
		return
	}

	/*
		data := struct {
			Title   string
			Message string
		}{
			Title:   "Pagina de inicio",
			Message: "Bienvenido a piedra, papel y tijera",
		}
	*/
	err = tpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear nuevo juego")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Juego")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")
}

func AboutMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Acerca de")
}
