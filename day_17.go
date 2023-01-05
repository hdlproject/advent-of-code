package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day17SolutionA() {
	f, err := os.Open("day_17_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	chamber := &chamber{
		highestRock: -1,
		space:       make([][]bool, 0, 7),
		maxLength:   7,
		coverMap:    make(map[int]int),
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		chamber.gasPushes = strings.Split(value, "")
		break
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	chamber.execute()

	chamber.printColoredSpace()

	fmt.Println(chamber.highestRock + chamber.base + 1)
}

func day17SolutionB() {
	f, err := os.Open("day_17_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	chamber := &chamber{
		highestRock: -1,
		space:       make([][]bool, 0, 7),
		maxLength:   7,
		coverMap:    make(map[int]int),
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		chamber.gasPushes = strings.Split(value, "")
		break
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	chamber.execute()

	chamber.printColoredSpace()

	fmt.Println(chamber.highestRock + chamber.base + 1)
}

//func main() {
//	day17SolutionB()
//}
