package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day4SolutionA() {
	f, err := os.Open("day_4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var count int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, ",")

		firstVal := strings.Split(values[0], "-")
		secondVal := strings.Split(values[1], "-")

		firstLow, _ := strconv.ParseInt(firstVal[0], 10, 64)
		firstHigh, _ := strconv.ParseInt(firstVal[1], 10, 64)
		secondLow, _ := strconv.ParseInt(secondVal[0], 10, 64)
		secondHigh, _ := strconv.ParseInt(secondVal[1], 10, 64)

		if firstLow >= secondLow && firstHigh <= secondHigh {
			count++
		} else if secondLow >= firstLow && secondHigh <= firstHigh {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

func day4SolutionB() {
	f, err := os.Open("day_4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var count int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, ",")

		firstVal := strings.Split(values[0], "-")
		secondVal := strings.Split(values[1], "-")

		firstLow, _ := strconv.ParseInt(firstVal[0], 10, 64)
		firstHigh, _ := strconv.ParseInt(firstVal[1], 10, 64)
		secondLow, _ := strconv.ParseInt(secondVal[0], 10, 64)
		secondHigh, _ := strconv.ParseInt(secondVal[1], 10, 64)

		if (firstLow >= secondLow && firstLow <= secondHigh) || (firstHigh >= secondLow && firstHigh <= secondHigh) {
			count++
		} else if (secondLow >= firstLow && secondLow <= firstHigh) || (secondHigh >= firstLow && secondHigh <= firstHigh) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

//func main() {
//	day4SolutionB()
//}
