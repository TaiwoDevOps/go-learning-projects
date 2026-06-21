package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd numbers: ", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// rows := 5
	// for i := 1; i <= rows; i++ {
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {

	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	guessGame()

}

func guessGame() {

	target := rand.Intn(100) + 1

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")

	var guessCounter int = 5

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Take your pick")

	for scanner.Scan() {

		guess, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Enter only digits")
			continue
		}

		if guess == target {
			fmt.Printf("Good guess, your lucky number (%d) is correct\n", guess)
			break
		} else if guess < target {
			if (target - guess) <= 10 {
				fmt.Println("Your guessed value is pretty close but not correct")
			} else {

				fmt.Println("You are way too far from it, bump it up better")
			}
		} else {
			if (guess - target) <= 10 {
				fmt.Println("Your guessed value is pretty close but not correct")
			} else {

				fmt.Println("You are way too far from it, tone it down ")
			}
		}
		guessCounter--
		if guessCounter != 0 {
			fmt.Printf("You have (%d) tries left\n", guessCounter)
			fmt.Println("Take your pick")
		} else if guessCounter == 0 {
			fmt.Printf("You lost, number was %d.\nGame Over!!!\n", target)
			break
		}
	}

}
