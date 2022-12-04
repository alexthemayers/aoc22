package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const input = "./3/input.txt"

/*
Multiple rucksacks
Each with two compartments
*/

func main() {
	// Open Input file
	data, err := os.ReadFile(input)
	if err != nil {
		panic("Could not open file: " + err.Error())
	}
	rucksacks := strings.Split(bytes.NewBuffer(data).String(), "\n")
	items := make([]rune, 0, len(rucksacks))
	for i := range rucksacks {
		r := rucksacks[i]
		firstC := r[:len(r)/2]
		secondC := r[len(r)/2:]
		for _, j := range firstC {
			l := len(items)
			for _, k := range secondC {
				if j == k {
					items = append(items, rune(j))
					break
				}
			}
			if len(items) != l {
				break
			}
		}
	}
	score := 0
	for s := range items {
		if items[s] > 96 {
			score += int(items[s]) - 96
			continue
		}
		score += int(items[s]) - 38
	}
	fmt.Println(score)
}
