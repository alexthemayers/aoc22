package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	itemsByWorryLevel []int
	operation         []string
	testNumber        int
	test              map[bool]int // eg: if true throw to monkey nr 1
}

const input = "./11/input.txt"

func main() {
	data, err := os.ReadFile(input)
	if err != nil {
		panic("Could not open file: " + err.Error())
	}
	monkeyData := strings.Split(bytes.NewBuffer(data).String(), "\n\n")
	monkeys := make(map[int]Monkey, len(monkeyData))
	for _, m := range monkeyData {
		i, monkey := makeMonkey(m)
		monkeys[i] = monkey
	}
	for i := 0; i < len(monkeys); i++ {
		for j := 0; j < len(monkeys[i].itemsByWorryLevel); j++ {

		}
	}
}

func makeMonkey(mData string) (int, Monkey) {
	id := 0

	monkey := Monkey{
		test: make(map[bool]int, 2),
	}
	for _, line := range strings.Split(mData, "\n") {
		tokens := strings.Split(line, " ")
		if strings.HasPrefix(line, "Monkey") {
			i := strings.Split(tokens[1], ":")[0]
			var err error
			id, err = strconv.Atoi(i)
			if err != nil {
				log.Fatalf("PARSE ERROR - Monkey #: Couldn't parse string '%s' to int\n", i)
			}
		}
		if strings.HasPrefix(line, "  Starting") {
			for _, t := range tokens[4:] {
				if strings.Contains(t, ",") {
					t = strings.Split(t, ",")[0]
				}
				i, err := strconv.Atoi(t)
				if err != nil {
					log.Fatalf("PARSE ERROR - Starting: Couldn't parse string '%s' to int\n", t)
				}
				monkey.itemsByWorryLevel = append(monkey.itemsByWorryLevel, i)
			}
		}
		if strings.HasPrefix(line, "  Operation:") {
			for _, t := range tokens[5:] {
				monkey.operation = append(monkey.operation, t)
			}
		}
		if strings.HasPrefix(line, "  Test:") {
			i, err := strconv.Atoi(tokens[5])
			if err != nil {
				log.Fatalf("PARSE ERROR - Operation: Couldn't parse '%s' to int\n", tokens[5])
			}
			monkey.testNumber = i
		}
		if strings.HasPrefix(line, "    If true:") {
			i, err := strconv.Atoi(tokens[9])
			if err != nil {
				log.Fatalf("PARSE ERROR - If true: Couldn't parse '%s' to int\n", tokens[9])
			}
			monkey.test[true] = i
		}
		if strings.HasPrefix(line, "    If false:") {
			i, err := strconv.Atoi(tokens[9])
			if err != nil {
				log.Fatalf("PARSE ERROR - If false: Couldn't parse '%s' to int\n", tokens[9])
			}
			monkey.test[false] = i
		}
	}
	return id, monkey
}
