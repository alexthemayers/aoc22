package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const input = "./2/input.txt"

func main() {
	// Open Input file
	data, err := os.ReadFile(input)
	if err != nil {
		panic("Could not open file: " + err.Error())
	}
	stringData := strings.Split(bytes.NewBuffer(data).String(), "\n")

	totalScore := 0

	for i := range stringData {
		opponent := strings.Split(stringData[i], " ")[0]
		response := strings.Split(stringData[i], " ")[1]
		var opponentScore, responseScore int

		switch opponent {
		case "A":
			opponentScore += 1
		case "B":
			opponentScore += 2
		case "C":
			opponentScore += 3
		}

		switch response {
		case "X":
			responseScore += 1
		case "Y":
			responseScore += 2
		case "Z":
			responseScore += 3
		}

		// Draw Case
		if responseScore == opponentScore {
			fmt.Println("it's a draw")
			responseScore += 3
			totalScore += responseScore
			continue
		}

		// Win Case
		switch opponent {
		case "A":
			if response == "Y" {
				fmt.Println("we've won!")
				responseScore += 6
				totalScore += responseScore
				continue
			}
			if response == "Z" {
				fmt.Println("we've lost!")
				totalScore += responseScore
				continue
			}
		case "B":
			if response == "Z" {
				fmt.Println("we've won!")
				responseScore += 6
				totalScore += responseScore
				continue
			}
			if response == "X" {
				fmt.Println("we've lost!")
				totalScore += responseScore
				continue
			}
		case "C":
			if response == "X" {
				fmt.Println("we've won!")
				responseScore += 6
				totalScore += responseScore
				continue
			}
			if response == "Y" {
				fmt.Println("we've lost!")
				totalScore += responseScore
				continue
			}
		}
	}
	fmt.Printf("%d is the total score for part 1!\n", totalScore)
}
