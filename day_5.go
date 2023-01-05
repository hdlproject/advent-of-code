package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack []string

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(str string) {
	*s = append(*s, str)
}

func (s *stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *stack) MultiplePush(str []string) {
	*s = append(*s, str...)
}

func (s *stack) MultiplePop(count int) ([]string, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		var index int
		if len(*s) < count {
			index = len(*s)

		} else {
			index = len(*s) - count
		}

		elements := (*s)[index:]
		*s = (*s)[:index]
		return elements, true
	}
}

func day5SolutionA() {
	f, err := os.Open("day_5_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stacks := getStartStacks()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, " ")

		amount, _ := strconv.ParseInt(values[1], 10, 64)
		from, _ := strconv.ParseInt(values[3], 10, 64)
		to, _ := strconv.ParseInt(values[5], 10, 64)

		for i := 0; i < int(amount); i++ {
			if element, ok := stacks[from-1].Pop(); ok {
				stacks[to-1].Push(element)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, stack := range stacks {
		if element, ok := stack.Pop(); ok {
			fmt.Print(element)
		}
	}
}

func day5SolutionB() {
	f, err := os.Open("day_5_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stacks := getStartStacks()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, " ")

		amount, _ := strconv.ParseInt(values[1], 10, 64)
		from, _ := strconv.ParseInt(values[3], 10, 64)
		to, _ := strconv.ParseInt(values[5], 10, 64)

		if elements, ok := stacks[from-1].MultiplePop(int(amount)); ok {
			stacks[to-1].MultiplePush(elements)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, stack := range stacks {
		if element, ok := stack.Pop(); ok {
			fmt.Print(element)
		}
	}
}

func getStartStacks() []stack {
	return []stack{
		{"J", "H", "G", "M", "Z", "N", "T", "F"},
		{"V", "W", "J"},
		{"G", "V", "L", "J", "B", "T", "H"},
		{"B", "P", "J", "N", "C", "D", "V", "L"},
		{"F", "W", "S", "M", "P", "R", "G"},
		{"G", "H", "C", "F", "B", "N", "V", "M"},
		{"D", "H", "G", "M", "R"},
		{"H", "N", "M", "V", "Z", "D"},
		{"G", "N", "F", "H"},
	}

	//return []stack{
	//	{"Z", "N"},
	//	{"M", "C", "D"},
	//	{"P"},
	//}
}

//func main() {
//	day5SolutionB()
//}
