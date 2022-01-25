package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/sethvargo/go-diceware/diceware"
)

var guess string
var guesses int
var wholeWord string
var word []string
var correctGuesses []string
var wrongGuesses []string

func main() {
	clearScreen()
	generateVars()
	getStatus()
}

// clearScreen clears the terminal window
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS != "windows" {
		cmd = exec.Command("clear")
	} else {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

// generateWordVars initializes word vars
func generateVars() {
	tempSlice, err := diceware.Generate(1)
	if err != nil {
		log.Fatal(err)
	}
	wholeWord = tempSlice[0]
	word = strings.Split(tempSlice[0], "")
	for range word {
		correctGuesses = append(correctGuesses, "_")
	}
}

// getGuess reads from stdin and prints current game status
func getGuess() {
	var correctInput bool = false
	reader := bufio.NewReader(os.Stdin)

	printMan()
	fmt.Printf("✅ Correct guesses: ")
	for i := range correctGuesses {
		fmt.Printf("%v ", correctGuesses[i])
	}
	fmt.Println()

	fmt.Printf("❌ Wrong guesses (%v/11): ", guesses)
	for i := range wrongGuesses {
		fmt.Printf("%v ", wrongGuesses[i])
	}
	fmt.Println()

	for correctInput == false {
		fmt.Printf("📝 Enter a character: ")
		guess, _ = reader.ReadString('\n')
		checkInput(&correctInput)
	}
	clearScreen()
	checkGuess()
}

// checkInput takes a pointer to a state and determines if a guess is valid input
func checkInput(state *bool) {
	guess = strings.ToLower(strings.TrimSpace(guess))
	if len(guess) > 1 {
		fmt.Printf("🤔 You can only choose one character!\n")
	} else if len(guess) < 1 {
		fmt.Printf("🤔 Did you really choose a character?\n")
	} else if strings.Contains(strings.Join(wrongGuesses[:], ","), guess) || strings.Contains(strings.Join(correctGuesses[:], ","), guess) {
		fmt.Printf("🤔 You already guessed that character, try another one!\n")
	} else {
		*state = true
	}
}

// checkGuess determines if a guess is correct or not and then calls getGuess()
func checkGuess() {
	if strings.Contains(wholeWord, guess) == true {
		for i := range word {
			if guess == word[i] {
				correctGuesses[i] = guess
			}
		}
	} else {
		wrongGuesses = append(wrongGuesses, guess)
		guesses++
	}
	getStatus()
}

// getStatus determines if the game should continue or not.
func getStatus() {
	if strings.Compare(strings.Join(word[:], ","), strings.Join(correctGuesses[:], ",")) == 0 && guesses <= 11 {
		printMan()
		fmt.Printf("🥳 You won! The word was: %v\n", wholeWord)
		os.Exit(0)
	} else if guesses >= 11 {
		printMan()
		fmt.Printf("😭 Sorry, you lost, the word was: %v\n", wholeWord)
		os.Exit(0)
	} else {
		getGuess()
	}
}

//  printMan is a shitty way of painting the hangman.
func printMan() {
	if guesses == 0 {
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
	} else if guesses == 1 {
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("___")
	} else if guesses == 2 {
		fmt.Println(" ")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println("_|_")
	} else if guesses == 3 {
		fmt.Println(" _______")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println("_|_")
	} else if guesses == 4 {
		fmt.Println(" _______")
		fmt.Println(" |/")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println("_|_")
	} else if guesses == 5 {
		fmt.Println(" _______")
		fmt.Println(" |/     |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println("_|_")
	} else if guesses == 6 {
		fmt.Println(" _______")
		fmt.Println(" |/     |")
		fmt.Println(" |     (_)")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println("_|_")
	} else if guesses == 7 {
		fmt.Println(" _______")
		fmt.Println(" |/     |")
		fmt.Println(" |     (_)")
		fmt.Println(" |      |")
		fmt.Println(" |      |")
		fmt.Println(" |")
		fmt.Println(" |")
		fmt.Println("_|_")
	} else if guesses == 8 {
		fmt.Println(" _______")
		fmt.Println(" |/     |")
		fmt.Println(" |     (_)")
		fmt.Println(" |      |")
		fmt.Println(" |      |")
		fmt.Println(" |     /")
		fmt.Println(" |")
		fmt.Println("_|_")
	} else if guesses == 9 {
		fmt.Println(" _______")
		fmt.Println(" |/     |")
		fmt.Println(" |     (_)")
		fmt.Println(" |      |")
		fmt.Println(" |      |")
		fmt.Println(" |     / \\")
		fmt.Println(" |")
		fmt.Println("_|_")
	} else if guesses == 10 {
		fmt.Println(" _______")
		fmt.Println(" |/     |")
		fmt.Println(" |     (_)")
		fmt.Println(" |     \\|")
		fmt.Println(" |      |")
		fmt.Println(" |     / \\")
		fmt.Println(" |")
		fmt.Println("_|_")
	} else if guesses == 11 {
		fmt.Println(" _______")
		fmt.Println(" |/     |")
		fmt.Println(" |     (_)")
		fmt.Println(" |     \\|/")
		fmt.Println(" |      |")
		fmt.Println(" |     / \\")
		fmt.Println(" |")
		fmt.Println("_|_")
	}
}
