package main

import (
	"cryptopals-go/challenges"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a challenge number")
		os.Exit(1)
	}

	challengeNumber := os.Args[1]

	switch challengeNumber {
	case "1":
		challenges.Challenge1()
	case "2":
		challenges.Challenge2()
	case "3":
		challenges.Challenge3()
	case "4":
		challenges.Challenge4()
	default:
		fmt.Println("Invalid challenge number")
	}

}
