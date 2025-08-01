package main

import (
	"fmt"
	"net/http"
	"rpsweb/handlers"
)

func main() {

	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/newGame", handlers.NewGame)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	http.ListenAndServe(port, router)
}
