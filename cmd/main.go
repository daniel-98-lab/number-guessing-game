package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/daniel-98-lab/number-guessing-game/internal/enums"
)

func main() {
	fmt.Println("*************************************************************")
	fmt.Println("Welcome to The Number Guessing Game, Are you ready to lost???")
	fmt.Println("Please select the difficult level: ")
	fmt.Println("Easy: 10")
	fmt.Println("Medium: 5")
	fmt.Println("Hard: 2")
	fmt.Println("*************************************************************")

	var diff string
	var difficultyWithChancesMap = map[enums.Difficulty]int{
		enums.Easy:   10,
		enums.Medium: 5,
		enums.Hard:   2,
	}
	var number int
	wasWin := false

	for {
		fmt.Println("Please select some difficulty level (easy, medium, hard)")

		//Input difficulty
		fmt.Scanln(&diff)

		//Trim Spaces and To lower diff
		diff = strings.ToLower(strings.TrimSpace(diff))

		if _, exists := difficultyWithChancesMap[enums.Difficulty(diff)]; exists {
			break
		}
		fmt.Println("Invalid selection. Please choose between easy, medium, or hard.")
	}

	chances := difficultyWithChancesMap[enums.Difficulty(diff)]

	fmt.Printf("Chances: %d\n", chances)

	//Random Number
	randomNumber := rand.Intn(100) + 1

	fmt.Printf("Random Number: %d\n", randomNumber)

	for i := 0; i < chances; i++ {
		fmt.Println("Please write a number: ")

		//Input difficulty
		fmt.Scanln(&number)

		if number == randomNumber {
			wasWin = true
			fmt.Printf("Congratulations you won!!!!, Number of attempts: %d\n", i)
			break
		} else {
			fmt.Printf("Wrong answer!!")
		}
	}

	if !wasWin {
		fmt.Println("You lost, you don't luck!!")
	}
}
