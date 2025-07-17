package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // Piedra. Vence a las tijeras. (tijeras + 1) % 3 = 0
	PAPER    = 1 // papel. vence a piedra. (piedra + 1) % 3 = 1
	SCISSORS = 2 // tijeras. vence al pepel. (papel + 1) % 3 = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChioce    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"Bien echo!",
	"Buen trabajo!",
	"Deberias comprar un boleto de loteria",
}

var loseMessages = []string{
	"Que lastima!",
	"Intentalo de nuevo!",
	"Hoy simplemente no es tu dia.",
}

var drawMessages = []string{
	"Las grandes mentes piensan iguales",
	"Oh Oh. Intentalo de nuevo.",
	"Nadie gana, pero puedes intentarlo de nuevo.",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligío PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligío PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligío SCISSORS"
	}

	messageInt := rand.Intn(3)
	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana!"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora gana!"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChioce:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}

}
