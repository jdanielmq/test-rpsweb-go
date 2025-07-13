package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo")
}
func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "creando un nuevo juego...!!")
}
func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Juego")
}
func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jugar el juego")
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Acerca del Juego")
}
