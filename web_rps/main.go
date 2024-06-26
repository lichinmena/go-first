package main

import (
	"log"
	"net/http"
	"web_rps/handlers"
)

//Handler -> Funcion que se encarga de manejar una solicitud http

//Roter -> Dirige las peticiones http al handler indicado

func main() {
	//Crear enrutador
	router := http.NewServeMux()

	//Configurar las rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.AboutMe)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, router))
}
