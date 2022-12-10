package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type node struct {
	x int // + is right, - is left
	y int // + is up, - is down
}

// recordPosition records each position of a node to positions
func recordPosition(k *node, positions *[]node) {
	temp := node{x: k.x, y: k.y}
	for _, p := range *positions {
		if reflect.DeepEqual(p, temp) {
			return // get out if it exists
		}
	}
	*positions = append(*positions, temp)
}

// updateTail moves the tail appropriately based on the relevant movement of the head
func updateTail(head *node, tail *node) {
	if head.y-tail.y > 2 {
		log.Printf("something is fishy. head.y:: %d\ttail.y:: %d\n", head.y, tail.y)
	}
	switch head.y - tail.y {
	case 2: // head 2 above
		tail.y++
		if head.x != tail.x {
			tail.x = head.x
		}
	case -2: // head 2 below
		tail.y--
		if head.x != tail.x {
			tail.x = head.x
		}
	}

	if head.x-tail.x > 2 {
		log.Printf("something is fishy. head.x:: %d\ttail.x:: %d\n", head.x, tail.x)
	}
	switch head.x - tail.x {
	case 2: // head 2 above
		tail.x++
		if head.y != tail.y {
			tail.y = head.y
		}
	case -2: // head 2 below
		tail.x--
		if head.y != tail.y {
			tail.y = head.y
		}
	}
}

// move updates the position of the head
func move(k *node, direction string) {
	switch direction {
	case "R":
		k.x++
	case "L":
		k.x--
	case "U":
		k.y++
	case "D":
		k.y--
	default:
		log.Fatal("should not get here!")
	}
}

const input = "./9/input.txt"

func main() {
	data, err := os.ReadFile(input)
	if err != nil {
		panic("Could not open file: " + err.Error())
	}
	inputLines := strings.Split(bytes.NewBuffer(data).String(), "\n")

	head := node{x: 0, y: 0}
	tail := node{x: 0, y: 0}
	positions := make([]node, 0, 1024)
	totalMoves := 0

	for _, line := range inputLines {
		direction := strings.Split(line, " ")[0]
		steps, _ := strconv.Atoi(strings.Split(line, " ")[1])
		totalMoves += steps
		for i := 0; i < steps; i++ {
			move(&head, direction)
			updateTail(&head, &tail)
			//printPositions(&head, &tail)
			recordPosition(&tail, &positions)
		}
	}

	fmt.Printf("number of positions :: %d\n", len(positions))
	fmt.Printf("number of total moves :: %d\n", totalMoves)
}

func printPositions(head *node, tail *node) {
	for y := 16; y > 0; y-- {
		for x := 0; x < 16; x++ {
			if head.x+8 == x && head.y+8 == y {
				fmt.Printf("H")
				continue
			}
			if tail.x+8 == x && tail.y+8 == y {
				fmt.Printf("T")
				continue
			}
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("HEAD [x::%d y::%d]\n", head.x, head.y)
	fmt.Printf("TAIL [x::%d y::%d]\n", tail.x, tail.y)
	fmt.Println("")
}
