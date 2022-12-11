package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	cycleNumber     = 0
	xRegister       = 1
	signalStrengths = map[int]int{20: 0, 60: 0, 100: 0, 140: 0, 180: 0, 220: 0}
)

const input = "./10/input.txt"

func cycle() {
	cycleNumber++
	for k := range signalStrengths {
		if cycleNumber == k {
			strength := signalStrength()
			signalStrengths[k] = strength
		}
	}
}

func noop() {
	cycle()
}
func addx(amt int) {
	cycle()
	cycle()
	xRegister += amt

}

func signalStrength() int {
	return cycleNumber * xRegister
}

func main() {

	data, err := os.ReadFile(input)
	if err != nil {
		panic("Could not open file: " + err.Error())
	}
	inputLines := strings.Split(bytes.NewBuffer(data).String(), "\n")
	for _, line := range inputLines {
		tokens := strings.Split(line, " ")
		if len(tokens) > 1 {
			amt, err := strconv.Atoi(tokens[1])
			if err != nil {
				log.Fatal(err.Error())
			}
			addx(amt)
			continue
		}
		noop()
	}
	fmt.Printf("Cycles:: %d\t\tRegister:: %d\n", cycleNumber, xRegister)
	signalStrengthSum := 0
	for k := range signalStrengths {
		signalStrengthSum += signalStrengths[k]
	}
	fmt.Printf("Signal strength sum:: %d\n", signalStrengthSum)

}
