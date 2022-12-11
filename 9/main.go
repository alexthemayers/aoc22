package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	NumberOfNodes   = 10
	PrintGridSize   = 48
	PrintGridOffset = PrintGridSize / 2
	PrintDelay      = 100
	LiveRender      = false
	Debug           = false
)

type node struct {
	x int // + is right, - is left
	y int // + is up, - is down
}

func generateNodes(x int, y int) []*node {
	nodes := make([]*node, 0, NumberOfNodes)
	for i := 0; i < NumberOfNodes; i++ {
		nodes = append(nodes, &node{x: x, y: y})
	}
	return nodes
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

// updateTail moves the tail appropriately based on the relative movement of the head
func updateTail(head *node, tail *node) {
	// y axis movement
	if head.x == tail.x {
		if Debug {
			if head.y == tail.y {
				return
			}
			if head.y-tail.y == 1 || head.y-tail.y == -1 {
				return
			}
		}

		if head.y-tail.y == 2 {
			tail.y++
			return
		}
		if head.y-tail.y == -2 {
			tail.y--
			return
		}
	}
	// x axis movement
	if head.y == tail.y {
		if Debug {
			if head.x-tail.x == 1 || head.x-tail.x == -1 {
				return
			}
		}

		if head.x-tail.x == 2 {
			tail.x++
			return
		}
		if head.x-tail.x == -2 {
			tail.x--
			return
		}
	}

	if Debug {
		if (head.x-tail.x == -1 && head.y-tail.y == 1) ||
			(head.x-tail.x == 1 && head.y-tail.y == -1) {
			return
		}

		if (head.x-tail.x == -1 && head.y-tail.y == -1) ||
			(head.x-tail.x == 1 && head.y-tail.y == 1) {
			return
		}
	}

	// TOP LEFT
	if (head.x-tail.x == -2 && head.y-tail.y == 1) ||
		(head.x-tail.x == -1 && head.y-tail.y == 2) ||
		(head.x-tail.x == -2 && head.y-tail.y == 2) { // DIAGONAL CASE
		tail.x--
		tail.y++
		return
	}

	// TOP RIGHT
	if (head.x-tail.x == 1 && head.y-tail.y == 2) ||
		(head.x-tail.x == 2 && head.y-tail.y == 1) ||
		(head.x-tail.x == 2 && head.y-tail.y == 2) { // DIAGONAL CASE
		tail.x++
		tail.y++
		return
	}

	// BOTTOM RIGHT
	if (head.x-tail.x == 2 && head.y-tail.y == -1) ||
		(head.x-tail.x == 1 && head.y-tail.y == -2) ||
		(head.x-tail.x == 2 && head.y-tail.y == -2) { // DIAGONAL CASE
		tail.x++
		tail.y--
		return
	}

	// BOTTOM LEFT
	if (head.x-tail.x == -1 && head.y-tail.y == -2) ||
		(head.x-tail.x == -2 && head.y-tail.y == -1) ||
		(head.x-tail.x == -2 && head.y-tail.y == -2) { // DIAGONAL CASE
		tail.x--
		tail.y--
		return
	}

	if Debug {
		log.Printf("The shit has hit the fan\nHead:: [x: %d | y: %d]\nTail:: [x: %d | y: %d]\nDelta:: [x: %d | y: %d]\\n", head.x, head.y, tail.x, tail.y, head.x-tail.x, head.y-tail.y)
	}
}

// recordPosition records each position of a node to positions
func recordPosition(k *node, positions *[]*node) {
	for _, n := range *positions {
		if n.x == k.x && n.y == k.y {
			return // get out if it exists
		}
	}
	temp := node{x: k.x, y: k.y}
	*positions = append(*positions, &temp)
}

func printPositions(moves int, nodes []*node) {
	for y := PrintGridSize; y > 0; y-- {
		for x := 0; x < PrintGridSize; x++ {
			printed := false
			for i := 0; i < NumberOfNodes; i++ {
				if nodes[i].x+PrintGridOffset == x && nodes[i].y+PrintGridOffset == y {
					if !printed {
						s := strconv.Itoa(i + 1)
						if len(s) == 1 {
							fmt.Printf(" %s", s)
						} else {
							fmt.Printf("%s", s)
						}
						printed = true
					}
				}
			}
			if !printed {
				fmt.Printf(" .")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("")
	if Debug {
		fmt.Printf("move:: %d\n", moves)
		for i := 0; i < NumberOfNodes; i++ {
			fmt.Printf("NODE %d [x::%d y::%d]\n", i+1, nodes[i].x, nodes[i].y)
		}
	}
	time.Sleep(time.Millisecond * PrintDelay)
}

const input = "./9/input.txt"

func main() {

	data, err := os.ReadFile(input)
	if err != nil {
		panic("Could not open file: " + err.Error())
	}
	inputLines := strings.Split(bytes.NewBuffer(data).String(), "\n")

	nodes := generateNodes(0, 0)
	fmt.Printf("we have %d nodes\n", len(nodes))
	positions := make([]*node, 0, 1024)
	totalMoves := 0

	for _, line := range inputLines {
		direction := strings.Split(line, " ")[0]
		steps, _ := strconv.Atoi(strings.Split(line, " ")[1])
		totalMoves += steps
		for i := 0; i < steps; i++ {
			move(nodes[0], direction)
			for j := 0; j < NumberOfNodes-1; j++ {
				updateTail(nodes[j], nodes[j+1])
			}
			if LiveRender {
				printPositions(totalMoves, nodes)
			}
			recordPosition(nodes[len(nodes)-1], &positions)
		}
	}
	fmt.Printf("number of positions :: %d\n", len(positions))
	fmt.Printf("number of total moves :: %d\n", totalMoves)
}
