package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	inspectedItem      []int
	operation          func(worryLevel int) int
	test               func(worryLevel int) bool
	trueTo             int
	falseTo            int
	numberOfInspection int
}

type byInspection []monkey

func (m byInspection) Len() int           { return len(m) }
func (m byInspection) Less(i, j int) bool { return m[i].numberOfInspection > m[j].numberOfInspection }
func (m byInspection) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func day11SolutionA() {
	f, err := os.Open("day_11_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	monkeys := make(map[int]monkey)
	var id int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(strings.Trim(value, " "), " ")

		switch values[0] {
		case "Monkey":
			id64, _ := strconv.ParseInt(strings.TrimSuffix(values[1], ":"), 10, 64)
			id = int(id64)
		case "Starting":
			index := strings.Index(value, ":")
			itemStr := value[index+1:]
			items := strings.Split(itemStr, ",")
			var inspectedItems []int
			for _, item := range items {
				itemInt, _ := strconv.ParseInt(strings.Trim(item, " "), 10, 64)
				inspectedItems = append(inspectedItems, int(itemInt))
			}
			currentMonkey := monkeys[id]
			currentMonkey.inspectedItem = inspectedItems
			monkeys[id] = currentMonkey
		case "Operation:":
			currentMonkey := monkeys[id]
			currentMonkey.operation = getOperation(values[4], values[5])
			monkeys[id] = currentMonkey
		case "Test:":
			currentMonkey := monkeys[id]
			currentMonkey.test = getTest(values[3])
			monkeys[id] = currentMonkey
		case "If":
			currentMonkey := monkeys[id]
			operandInt, _ := strconv.ParseInt(strings.Trim(values[5], " "), 10, 64)

			switch values[1] {
			case "true:":
				currentMonkey.trueTo = int(operandInt)
			case "false:":
				currentMonkey.falseTo = int(operandInt)
			}
			monkeys[id] = currentMonkey
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			doType1(j, monkeys)
		}
	}

	var sortedMonkeys byInspection
	for j := 0; j < len(monkeys); j++ {
		sortedMonkeys = append(sortedMonkeys, monkeys[j])
	}

	sort.Sort(sortedMonkeys)

	fmt.Println(sortedMonkeys[0].numberOfInspection * sortedMonkeys[1].numberOfInspection)
}

func day11SolutionB() {
	f, err := os.Open("day_11_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	monkeys := make(map[int]monkey)
	var id int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(strings.Trim(value, " "), " ")

		switch values[0] {
		case "Monkey":
			id64, _ := strconv.ParseInt(strings.TrimSuffix(values[1], ":"), 10, 64)
			id = int(id64)
		case "Starting":
			index := strings.Index(value, ":")
			itemStr := value[index+1:]
			items := strings.Split(itemStr, ",")
			var inspectedItems []int
			for _, item := range items {
				itemInt, _ := strconv.ParseInt(strings.Trim(item, " "), 10, 64)
				inspectedItems = append(inspectedItems, int(itemInt))
			}
			currentMonkey := monkeys[id]
			currentMonkey.inspectedItem = inspectedItems
			monkeys[id] = currentMonkey
		case "Operation:":
			currentMonkey := monkeys[id]
			currentMonkey.operation = getOperation(values[4], values[5])
			monkeys[id] = currentMonkey
		case "Test:":
			currentMonkey := monkeys[id]
			currentMonkey.test = getTest(values[3])
			monkeys[id] = currentMonkey
		case "If":
			currentMonkey := monkeys[id]
			operandInt, _ := strconv.ParseInt(strings.Trim(values[5], " "), 10, 64)

			switch values[1] {
			case "true:":
				currentMonkey.trueTo = int(operandInt)
			case "false:":
				currentMonkey.falseTo = int(operandInt)
			}
			monkeys[id] = currentMonkey
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			doType2(j, monkeys)
		}
	}

	var sortedMonkeys byInspection
	for j := 0; j < len(monkeys); j++ {
		sortedMonkeys = append(sortedMonkeys, monkeys[j])
	}

	sort.Sort(sortedMonkeys)

	fmt.Println(sortedMonkeys[0].numberOfInspection * sortedMonkeys[1].numberOfInspection)
}

func getOperation(operator, operand string) func(i int) int {
	var operandInt int64
	if strings.Trim(operand, " ") == "old" {
		operandInt = -1
	} else {
		operandInt, _ = strconv.ParseInt(strings.Trim(operand, " "), 10, 64)
	}

	if operandInt == -1 {
		switch strings.Trim(operator, " ") {
		case "*":
			return func(i int) int {
				return i * i
			}
		case "/":
			return func(i int) int {
				return i / i
			}
		case "+":
			return func(i int) int {
				return i + i
			}
		case "-":
			return func(i int) int {
				return i - i
			}
		default:
			return func(i int) int {
				return i
			}
		}
	}

	switch strings.Trim(operator, " ") {
	case "*":
		return func(i int) int {
			return i * int(operandInt)
		}
	case "/":
		return func(i int) int {
			return i / int(operandInt)
		}
	case "+":
		return func(i int) int {
			return i + int(operandInt)
		}
	case "-":
		return func(i int) int {
			return i - int(operandInt)
		}
	default:
		return func(i int) int {
			return i
		}
	}
}

func getTest(operand string) func(int) bool {
	operandInt, _ := strconv.ParseInt(strings.Trim(operand, " "), 10, 64)

	return func(i int) bool {
		return i%int(operandInt) == 0
	}
}

func doType1(currentId int, monkeys map[int]monkey) {
	currMonkey := monkeys[currentId]

	for _, item := range currMonkey.inspectedItem {
		worryLevel := currMonkey.operation(item)
		worryLevel /= 3

		test(currMonkey, worryLevel, monkeys)

		currMonkey.numberOfInspection++
	}

	currMonkey.inspectedItem = []int{}
	monkeys[currentId] = currMonkey
}

func doType2(currentId int, monkeys map[int]monkey) {
	currMonkey := monkeys[currentId]

	for _, item := range currMonkey.inspectedItem {
		worryLevel := currMonkey.operation(item)
		worryLevel = worryLevel % 9699690

		test(currMonkey, worryLevel, monkeys)

		currMonkey.numberOfInspection++
	}

	currMonkey.inspectedItem = []int{}
	monkeys[currentId] = currMonkey
}

func test(currMonkey monkey, worryLevel int, monkeys map[int]monkey) {
	if currMonkey.test(worryLevel) {
		destMonkey := monkeys[currMonkey.trueTo]

		destMonkey.inspectedItem = append(destMonkey.inspectedItem, worryLevel)
		monkeys[currMonkey.trueTo] = destMonkey
	} else {
		destMonkey := monkeys[currMonkey.falseTo]

		destMonkey.inspectedItem = append(destMonkey.inspectedItem, worryLevel)
		monkeys[currMonkey.falseTo] = destMonkey
	}
}

//func main() {
//	day11SolutionB()
//}
