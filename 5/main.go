package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const input = "./5/input.txt"

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func getStackNames(data []string) []string {
	ret := make([]string, 0, len(data[1]))
	for _, s := range data {
		if strings.HasPrefix(s, "[") {
			continue
		}
		ret = append(ret, strings.Split(s, "")...)
		break
	}
	return ret
}

func getStackContents(data []string) Stack {
	var ret Stack
	for _, s := range data {
		if !strings.HasPrefix(s, "[") {
			break
		}
		ret = append(ret, strings.Split(s, ""))
	}
	return ret
}

func zipStackContents(stackNames []string, stack Stack) Stack {
	s := make(map[string][]string)
	for _, lineData := range stack {
		for i := 0; i < len(lineData) && i < len(stackNames); i++ {
			if stackNames[i] == " " {
				continue
			}
			s[stackNames[i]] = append(s[stackNames[i]], lineData[i])
		}
	}
	stack = make([][]string, len(s))
	for k, v := range s {
		if ks, err := strconv.Atoi(k); err == nil {
			i := ks - 1
			stack[i] = make([]string, 0, len(v))
			stack[i] = reverse(v)
		}
	}
	return stack
}

type Stack [][]string

func cleanStack(s Stack) (stack Stack) {
	stack = make(Stack, len(s)+1)
	stack[0] = []string{}
	for i, c := range s {
		var temp []string
		for _, char := range c {
			if char != " " {
				temp = append(temp, char)
			}
		}
		stack[i+1] = make([]string, 0, len(temp))
		stack[i+1] = temp
	}
	return stack
}

func (s *Stack) makeAMove(move string, sameOrder bool) error {
	// move 3 from 9 to 1
	times, err := strconv.Atoi(strings.Split(move, " ")[1])
	if err != nil {
		return err
	}
	from, err := strconv.Atoi(strings.Split(move, " ")[3])
	if err != nil {
		return err
	}
	to, err := strconv.Atoi(strings.Split(move, " ")[5])
	if err != nil {
		return err
	}
	s.move(from, to, times, sameOrder)
	return nil
}

func (s *Stack) move(src, dest, times int, sameOrder bool) {
	var temp Stack = *s

	newSrc := make([]string, 0, len(temp[src])-times)
	newSrc = append(newSrc, temp[src][:len(temp[src])-times]...)
	newDest := temp[dest]
	if sameOrder {
		newDest = append(newDest, temp[src][len(temp[src])-times:]...)
	} else {
		newDest = append(newDest, reverse(temp[src][len(temp[src])-times:])...)
	}

	//fmt.Printf("        %#v\n", temp[src])
	temp[src] = newSrc
	//fmt.Printf("is now: %#v\n\n", temp[src])

	//fmt.Printf("        %#v\n", temp[dest])
	temp[dest] = newDest
	//fmt.Printf("is now: %#v\n\n", temp[dest])

	s = &temp
}

func main() {
	sameOrder := false

	//
	// This line enables output for part 2
	sameOrder = true
	//

	// Open Input file
	data, err := os.ReadFile(input)
	if err != nil {
		panic("Could not open file: " + err.Error())
	}
	inputLines := strings.Split(bytes.NewBuffer(data).String(), "\n")

	// populate stacks
	stackNames := getStackNames(inputLines)
	contents := getStackContents(inputLines)
	unsortedDirtyStack := zipStackContents(stackNames, contents)
	stack := cleanStack(unsortedDirtyStack)

	for _, line := range inputLines {
		if !strings.HasPrefix(line, "move") {
			continue
		}
		err := stack.makeAMove(line, sameOrder)
		if err != nil {
			log.Fatal(err)
		}
	}

	for k, v := range stack {
		fmt.Printf("%s [len: %.2d ] ::: %s \n", k, len(v), v)
	}
	for _, i := range stack {
		if len(i) == 0 {
			continue
		}
		fmt.Printf("%s", i[len(i)-1])
	}
	fmt.Printf("\n\n")
}
