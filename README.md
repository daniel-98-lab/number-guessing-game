# number-guessing-game

Build a simple number guessing game to test your luck.  Challenge from [number-guessing-game](https://roadmap.sh/projects/number-guessing-game)

## Features

- The game offers up to three difficulty levels.
- At the end of the game, if you win, you can see the number of attempts it took to win.
- You can view how much time it took you to guess the correct number.
- If you want to play again, you just need to confirm to start over.
- For each incorrect attempt, you will receive a series of hints to help you guess the number.

## Installation

To get started, you will need to have Go installed on your machine. Follow the instructions [here](https://golang.org/doc/install) if you don't have Go set up.

1. Clone this repository:
```bash
git clone https://github.com/daniel-98-lab/number-guessing-game.git
```

2. Navigate to the project Repository
```bash
cd number-guessing-game
```

3. Install dependencies:
```bash
go mod tidy
```

4. Build the project:
```bash
go build -o number-guessing-game ./cmd
```

## Usage
To use the game, run the following command from the terminal
```bash
./number-guessing-game
```

## Example:
```bash
*************************************************************
Welcome to The Number Guessing Game, Are you ready to lost???
*************************************************************
Please select the difficulty level: 
Easy: 10
Medium: 5
Hard: 2
10

Invalid selection. Please choose between easy, medium, or hard.
easy

Please write a number: 
12
Wrong answer!!Too low!

Please write a number: 
45
Wrong answer!!Too high!

Please write a number: 
33
Wrong answer!!Too high!

Please write a number: 
23

Congratulations you won!!!!, Number of attempts: 4
Time to guess the number:
0 days, 0 hours, 0 minutes, and 19 seconds

Do you want to play more: true or false?
false
```
