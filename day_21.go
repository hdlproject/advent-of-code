package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type monkeyNode struct {
	isNumber       bool
	number         int64
	operation      string
	anotherMonkeys []string
}

func day21SolutionA() {
	f, err := os.Open("day_21_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	monkeyMap := make(map[string]*monkeyNode)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		values := strings.Split(value, " ")
		name := strings.TrimSuffix(values[0], ":")

		switch len(values) {
		case 2:
			number, _ := strconv.ParseInt(values[1], 10, 64)

			monkeyMap[name] = &monkeyNode{
				isNumber: true,
				number:   number,
			}
		case 4:
			monkeyMap[name] = &monkeyNode{
				operation:      values[2],
				anotherMonkeys: []string{values[1], values[3]},
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := execute(monkeyMap, "root")

	fmt.Println(result)
}

func execute(monkeyMap map[string]*monkeyNode, name string) int64 {
	currMonkey := monkeyMap[name]

	if currMonkey.isNumber {
		return currMonkey.number
	} else {
		switch currMonkey.operation {
		case "+":
			return execute(monkeyMap, currMonkey.anotherMonkeys[0]) + execute(monkeyMap, currMonkey.anotherMonkeys[1])
		case "-":
			return execute(monkeyMap, currMonkey.anotherMonkeys[0]) - execute(monkeyMap, currMonkey.anotherMonkeys[1])
		case "*":
			return execute(monkeyMap, currMonkey.anotherMonkeys[0]) * execute(monkeyMap, currMonkey.anotherMonkeys[1])
		case "/":
			return execute(monkeyMap, currMonkey.anotherMonkeys[0]) / execute(monkeyMap, currMonkey.anotherMonkeys[1])
		}
	}

	return 0
}

func day21SolutionB() {
	f, err := os.Open("day_21_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	monkeyMap := make(map[string]*monkeyNode)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		values := strings.Split(value, " ")
		name := strings.TrimSuffix(values[0], ":")

		switch len(values) {
		case 2:
			number, _ := strconv.ParseInt(values[1], 10, 64)

			monkeyMap[name] = &monkeyNode{
				isNumber: true,
				number:   number,
			}
		case 4:
			monkeyMap[name] = &monkeyNode{
				operation:      values[2],
				anotherMonkeys: []string{values[1], values[3]},
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rootMonkey := monkeyMap["root"]
	left := findHuman(monkeyMap, rootMonkey.anotherMonkeys[0])

	var result int64
	if !left {
		result = execute(monkeyMap, rootMonkey.anotherMonkeys[0])
		result = executeType2(monkeyMap, rootMonkey.anotherMonkeys[1], result)
	} else {
		result = execute(monkeyMap, rootMonkey.anotherMonkeys[1])
		result = executeType2(monkeyMap, rootMonkey.anotherMonkeys[0], result)
	}

	fmt.Println(result)
}

func findHuman(monkeyMap map[string]*monkeyNode, name string) bool {
	if name == "humn" {
		return true
	} else {
		currMonkey := monkeyMap[name]
		if currMonkey.isNumber {
			return false
		}

		return findHuman(monkeyMap, currMonkey.anotherMonkeys[0]) || findHuman(monkeyMap, currMonkey.anotherMonkeys[1])
	}
}

func executeType2(monkeyMap map[string]*monkeyNode, name string, expectedResult int64) int64 {
	if name == "humn" {
		return expectedResult
	}

	currMonkey := monkeyMap[name]
	left := findHuman(monkeyMap, currMonkey.anotherMonkeys[0])

	var result int64
	var calculatedMonkey string
	if !left {
		result = execute(monkeyMap, currMonkey.anotherMonkeys[0])
		calculatedMonkey = currMonkey.anotherMonkeys[1]
	} else {
		result = execute(monkeyMap, currMonkey.anotherMonkeys[1])
		calculatedMonkey = currMonkey.anotherMonkeys[0]
	}

	switch currMonkey.operation {
	case "+":
		expectedResult = expectedResult - result
	case "-":
		if left {
			expectedResult = expectedResult + result
		} else {
			expectedResult = result - expectedResult
		}
	case "*":
		expectedResult = expectedResult / result
	case "/":
		if left {
			expectedResult = expectedResult * result
		} else {
			expectedResult = result / expectedResult
		}
	}

	return executeType2(monkeyMap, calculatedMonkey, expectedResult)
}

//func main() {
//	day21SolutionB()
//}
