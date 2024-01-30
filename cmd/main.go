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
	case "5":
		challenges.Challenge5()
	case "6":
		challenges.Challenge6()
	case "7":
		challenges.Challenge7()
	case "8":
		challenges.Challenge8()
	case "9":
		challenges.Challenge9()
	case "10":
		challenges.Challenge10()
	case "11":
		challenges.Challenge11()
	case "12":
		challenges.Challenge12()
	case "13":
		challenges.Challenge13()
	case "14":
		challenges.Challenge14()
	case "15":
		challenges.Challenge15()
	case "16":
		challenges.Challenge16()
	default:
		fmt.Println("Invalid challenge number")
	}

}
