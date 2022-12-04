package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const input = "./3/input.txt"

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
	fmt.Printf("Score for part 1: %d\n", score)

	// Part 2

	var badges = make([]rune, 0, len(rucksacks)/3)
	for i := 0; i < len(rucksacks)-2; i += 3 {
		badgesLen := len(badges)
		for _, j := range rucksacks[i] {
			for _, k := range rucksacks[i+1] {
				for _, l := range rucksacks[i+2] {
					if j == k && k == l {
						badges = append(badges, rune(l))
					}
					if len(badges) != badgesLen {
						break
					}
				}
				if len(badges) != badgesLen {
					break
				}
			}
			if len(badges) != badgesLen {
				break
			}
		}
	}

	score = 0
	for b := range badges {
		if badges[b] > 96 {
			score += int(badges[b]) - 96
			continue
		}
		score += int(badges[b]) - 38
	}
	fmt.Printf("Score for part 2: %d\n", score)
}
