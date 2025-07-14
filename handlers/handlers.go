package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir  = "./templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "index.html", nil)

	/*	tpl := template.Must(template.ParseFiles("./templates/base.html", "./templates/index.html"))
		/*	data := struct {
				Title   string
				Message string
			}{
				Title:   "Pagina de Inicio",
				Message: "¡Bienvenido! al juego de juna daniel",
			}*/

	/*
	   err := tpl.ExecuteTemplate(w, "base", nil)

	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}
	*/
}
func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "new-game.html", nil)
}
func Game(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "game.html", nil)
}
func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jugar el juego")
}
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))
	err := tpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
