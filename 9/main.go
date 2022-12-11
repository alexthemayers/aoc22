package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	NumberOfNodes   = 10
	PrintGridSize   = 128
	PrintGridOffset = PrintGridSize / 2
	PrintDelay      = 50
	LiveRender      = false
)

type node struct {
	x int // + is right, - is left
	y int // + is up, - is down
}

// recordPosition records each position of a node to positions
func recordPosition(k *[]*node, positions *[][]node) {
	// TODO find out how to optimize this. 99% of processing time is spent here
	temp := make([]node, 0, 1024)
	for _, n := range *k {
		temp = append(temp, node{x: n.x, y: n.y})
	}
	for _, p := range *positions {
		if reflect.DeepEqual(p, temp) {
			return // get out if it exists
		}
	}

	*positions = append(*positions, temp)
}

// updateTail moves the tail appropriately based on the relevant movement of the head
func updateTail(head *node, tail *node) {
	// y axis movement
	if head.x == tail.x {

		// HAPPY CASE
		if head.y == tail.y {
			return
		}
		if head.y-tail.y == 1 || head.y-tail.y == -1 {
			return
		}
		//

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

		// HAPPY CASE
		if head.x-tail.x == 1 || head.x-tail.x == -1 {
			return
		}
		//

		if head.x-tail.x == 2 {
			tail.x++
			return
		}
		if head.x-tail.x == -2 {
			tail.x--
			return
		}
	}

	// HAPPY CASE
	if (head.x-tail.x == -1 && head.y-tail.y == 1) ||
		(head.x-tail.x == 1 && head.y-tail.y == -1) {
		return
	}

	if (head.x-tail.x == -1 && head.y-tail.y == -1) ||
		(head.x-tail.x == 1 && head.y-tail.y == 1) {
		return
	}
	//

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
	log.Printf("The shit has hit the fan\nHead:: [x: %d | y: %d]\nTail:: [x: %d | y: %d]\nDelta:: [x: %d | y: %d]\\n", head.x, head.y, tail.x, tail.y, head.x-tail.x, head.y-tail.y)
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

	nodes := generateNodes(0, 0)
	fmt.Printf("we have %d nodes\n", len(nodes))
	positions := make([][]node, 0, 1024)
	totalMoves := 0

	for _, line := range inputLines {
		direction := strings.Split(line, " ")[0]
		steps, _ := strconv.Atoi(strings.Split(line, " ")[1])
		for i := 0; i < steps; i++ {
			totalMoves++
			if totalMoves == 57 {
				fmt.Printf("breakhere")
			}

			move(nodes[0], direction)
			for j := 0; j < NumberOfNodes-1; j++ {
				updateTail(nodes[j], nodes[j+1])
			}
			if LiveRender {
				printPositions(totalMoves, nodes)
			}
			recordPosition(&nodes, &positions)

		}
	}

	fmt.Printf("number of positions :: %d\n", len(positions))
	fmt.Printf("number of total moves :: %d\n", totalMoves)
}

func generateNodes(x int, y int) []*node {
	nodes := make([]*node, 0, NumberOfNodes)
	for i := 0; i < NumberOfNodes; i++ {
		nodes = append(nodes, &node{x: x, y: y})
	}
	return nodes
}

func printPositions(moves int, nodes []*node) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for y := PrintGridSize; y > 0; y-- {
		for x := 0; x < PrintGridSize; x++ {
			//printing position set

			printed := false
			for i := 0; i < NumberOfNodes; i++ {
				if nodes[i].x+PrintGridOffset == x && nodes[i].y+PrintGridOffset == y {
					if !printed {
						fmt.Printf(" %d", i+1)
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
	fmt.Printf("move:: %d\n", moves)
	for i := 0; i < NumberOfNodes; i++ {
		fmt.Printf("NODE %d [x::%d y::%d]\n", i+1, nodes[i].x, nodes[i].y)
	}
	time.Sleep(time.Millisecond * PrintDelay)
}
