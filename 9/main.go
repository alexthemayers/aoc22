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
	NumberOfNodes   = 9
	PrintGridSize   = 32
	PrintGridOffset = PrintGridSize / 2
	LiveRender      = true
)

type node struct {
	x int // + is right, - is left
	y int // + is up, - is down
}

// recordPosition records each position of a node to positions
func recordPosition(k *[]*node, positions *[][]node) {
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

const input = "./9/example_input.txt"

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
		totalMoves += steps
		for i := 0; i < steps; i++ {
			//if totalMoves == 43 {
			//	fmt.Printf("breakhere")
			//}
			move(nodes[0], direction)
			for j := 0; j < NumberOfNodes-1; j++ {
				updateTail(nodes[j], nodes[j+1])
			}
			if LiveRender {
				printPositions(nodes)
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

func printPositions(nodes []*node) {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()

	for y := PrintGridSize; y > 0; y-- {
		for x := 0; x < PrintGridSize; x++ {
			//printing position set

			printed := false
			for i := 0; i < NumberOfNodes; i++ {
				if nodes[i].x+PrintGridOffset == x && nodes[i].y+PrintGridOffset == y {
					if !printed {
						fmt.Printf("%d", i+1)
						printed = true
					}
				}
			}

			fmt.Printf(" ")

		}
		fmt.Printf("\n")
	}
	fmt.Println("")
	if LiveRender {
		for i := 0; i < NumberOfNodes; i++ {
			fmt.Printf("NODE %d [x::%d y::%d]\n", i+1, nodes[i].x, nodes[i].y)
		}
		time.Sleep(time.Millisecond * 120)
	}
}
