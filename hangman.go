// In this version of hangman I have tried to create a better structure for the game.
// Arguments are instead turned in to global variables.
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

// getGuess reads from stdin and checks the input for errors.
func getGuess() {
	var correctInput bool = false
	reader := bufio.NewReader(os.Stdin)

	printMan()
	fmt.Printf("✅ Correct guesses: %v\n", correctGuesses)
	fmt.Printf("❌ Wrong guesses (%v/11): %v\n", guesses, wrongGuesses)

	for correctInput == false {
		fmt.Printf("📝 Enter a character: ")
		guess, _ = reader.ReadString('\n')
		guess = strings.ToLower(strings.TrimSpace(guess))
		if len(guess) > 1 {
			fmt.Printf("🤔 You can only choose one character!\n")
		} else if len(guess) < 1 {
			fmt.Printf("🤔 Did you really choose a character?\n")
		} else {
			correctInput = true
		}
	}
	clearScreen()
	checkGuess()
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
