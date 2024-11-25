package services

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/daniel-98-lab/number-guessing-game/internal/enums"
)

// Types of Difficulty with Chances to play
var difficultyWithChancesMap = map[enums.Difficulty]int{
	enums.Easy:   10,
	enums.Medium: 5,
	enums.Hard:   2,
}

type NumberGuessingGame struct{}

func (s *NumberGuessingGame) Play() {
	playAgain := false
	var timeGuessNumber, timeStart time.Time
	for {
		timeStart = time.Now()
		fmt.Println("*************************************************************")
		fmt.Println("Welcome to The Number Guessing Game, Are you ready to lost???")
		fmt.Println("*************************************************************")
		chances := difficultyWithChancesMap[enums.Difficulty(getDifficulty())]
		randomNumber := getRandomNumber()

		isWin, attempts := guessGame(chances, randomNumber)
		if isWin {
			timeGuessNumber = time.Now()
			fmt.Printf("Congratulations you won!!!!, Number of attempts: %d\n", attempts)
			compareDate(timeStart, timeGuessNumber)
		} else {
			fmt.Println("You lost :(!!")
		}
		fmt.Println("Do you want to play more: true or false?")

		fmt.Scanln(&playAgain)
		if !playAgain {
			break
		}
	}
}

// Function to obtain difficulty, if the user chooses wrong, it is repeated until the correct difficulty is obtained.
func getDifficulty() string {
	fmt.Println("Please select the difficulty level: ")
	fmt.Println("Easy: 10")
	fmt.Println("Medium: 5")
	fmt.Println("Hard: 2")

	var diff string
	for {

		fmt.Scanln(&diff)

		// Trim Spaces and To lower diff
		diff = strings.ToLower(strings.TrimSpace(diff))

		if _, exists := difficultyWithChancesMap[enums.Difficulty(diff)]; exists {
			return diff
		}
		fmt.Println("Invalid selection. Please choose between easy, medium, or hard.")
	}
}

func getRandomNumber() int {
	return rand.Intn(100) + 1
}

// GuessGame handles the guessing logic, returning success status and attempts used.
func guessGame(chances, randomNumber int) (bool, int) {
	var number int
	for i := 0; i < chances; i++ {
		fmt.Println("Please write a number: ")

		_, err := fmt.Scanln(&number)

		// Check for scanning errors
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			clearInputBuffer()
			continue
		}

		if number == randomNumber {
			return true, i + 1
		} else {
			fmt.Printf("Wrong answer!!")
			hintProvide(randomNumber, number)
		}
	}
	return false, chances
}

// HELPERS

// function to clear input buffer
func clearInputBuffer() {
	var discard string
	fmt.Scanln(&discard) // Consume the rest of the input
}

// Function to compare two dates and show the difference between them
func compareDate(dateStart, dateEnd time.Time) {
	duration := dateEnd.Sub(dateStart)

	// Extract time components
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	fmt.Printf("Time to guess the number:\n")
	fmt.Printf("%d days, %d hours, %d minutes, and %d seconds\n", days, hours, minutes, seconds)
}

// HintProvide gives feedback on whether the user's guess is too low, too high, or close to the random number.
func hintProvide(randomNumber, numberUser int) {
	if numberUser < randomNumber {
		fmt.Println("Too low!")
	} else if numberUser > randomNumber {
		fmt.Println("Too high!")
	}

	if abs(numberUser-randomNumber) <= 5 {
		fmt.Println("You're close!")
	}
}

// Abs calculates the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
