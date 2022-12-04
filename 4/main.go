package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input = "./4/input.txt"

/*
Every section has an id number
each elf is assigned a range of sections

some assignments overlap
*/

func main() {
	// Open Input file
	data, err := os.ReadFile(input)
	if err != nil {
		panic("Could not open file: " + err.Error())
	}
	assignments := strings.Split(bytes.NewBuffer(data).String(), "\n")
	assignmentOverlaps := 0
	for _, assignment := range assignments {
		first := strings.Split(strings.Split(assignment, ",")[0], "-")
		second := strings.Split(strings.Split(assignment, ",")[1], "-")
		f0, _ := strconv.Atoi(first[0])
		f1, _ := strconv.Atoi(first[1])
		s0, _ := strconv.Atoi(second[0])
		s1, _ := strconv.Atoi(second[1])
		if (f0 <= s0 && f1 >= s1) || (s0 <= f0 && s1 >= f1) {
			assignmentOverlaps++
		}
	}
	fmt.Printf("Part 1: %d assignment overlaps\n", assignmentOverlaps)

}
