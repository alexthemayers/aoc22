package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input = "./1/input.txt"

func main() {
	data, err := os.ReadFile(input)
	if err != nil {
		panic("Failure reading file: " + input)
	}

	strData := strings.Split(bytes.NewBuffer(data).String(), "\n")
	perElf := make(map[int][]int)
	elfCount := 1
	var elf []int
	for _, line := range strData {
		if line != "" {
			i, err := strconv.Atoi(line)
			if err != nil {
				panic("Couldn't convert line " + line + " to int")
			}
			elf = append(elf, i)
			continue
		} else {
			perElf[elfCount] = elf
			elfCount++
			elf = nil
		}
	}

	elfTotals := make(map[int]int)
	mostCalories := 0

	for i := range perElf {
		sum := 0
		for _, contents := range perElf[i] {
			sum += contents
		}
		elfTotals[i] = sum

		// track largest
		if sum > elfTotals[mostCalories] {
			mostCalories = i
		}
	}

	//fmt.Printf("There are %d elves\n", len(perElf))
	//for elf, contents := range perElf {
	//	fmt.Printf("Elf %d has contents of %#v\n", elf, contents)
	//}
	//for elf, total := range elfTotals {
	//	fmt.Printf("Elf %d has calorie total of %#v\n", elf, total)
	//}

	fmt.Printf("Elf nr %d has the most calories at %d calories\n", mostCalories, elfTotals[mostCalories])

}
