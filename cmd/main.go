package main

import (
	"github.com/daniel-98-lab/number-guessing-game/internal/services"
)

func main() {
	service := services.NumberGuessingGame{}
	service.Play()
}
