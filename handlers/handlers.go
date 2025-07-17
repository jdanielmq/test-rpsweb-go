package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

type Player struct {
	Name string
}

var player Player

const (
	templateDir  = "./templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	restartValue()
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
	restartValue()
	renderTemplate(w, "new-game.html", nil)
}
func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("name")

	}
	fmt.Println("el nombre es: ", player.Name)
	if player.Name == "" {
		fmt.Println("redireccion a /newGame")
		http.Redirect(w, r, "newGame", http.StatusFound)
	}
	renderTemplate(w, "game.html", player)
}
func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)
	fmt.Println(result)

	out, err := json.MarshalIndent(result, "", "     ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	//fmt.Fprintf(w, "Jugar el juego")
}
func About(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func restartValue() {
	rps.ComputerScore = 0
	rps.PlayerScore = 0
	player.Name = ""

}
