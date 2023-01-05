package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day3SolutionA() {
	f, err := os.Open("day_3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var result []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		compartmentLen := len(value) / 2

		first := value[:compartmentLen]
		second := value[compartmentLen:]

		firstMap := make(map[string]struct{})
		for _, item := range strings.Split(first, "") {
			firstMap[item] = struct{}{}
		}

		secondMap := make(map[string]struct{})
		for _, item := range strings.Split(second, "") {
			secondMap[item] = struct{}{}
		}

		for item := range secondMap {
			if _, ok := firstMap[item]; ok {
				result = append(result, item)
				continue
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var priority int
	for _, item := range result {
		priority += getPriority(item)
	}
	fmt.Println(priority)
}

func day3SolutionB() {
	f, err := os.Open("day_3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var result []string
	var compartments []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		compartments = append(compartments, value)
		if len(compartments) == 3 {
			first := compartments[0]
			second := compartments[1]
			third := compartments[2]

			firstMap := make(map[string]struct{})
			for _, item := range strings.Split(first, "") {
				firstMap[item] = struct{}{}
			}

			secondMap := make(map[string]struct{})
			for _, item := range strings.Split(second, "") {
				secondMap[item] = struct{}{}
			}

			thirdMap := make(map[string]struct{})
			for _, item := range strings.Split(third, "") {
				thirdMap[item] = struct{}{}
			}

			for item := range firstMap {
				if _, ok := secondMap[item]; ok {
					if _, ok2 := thirdMap[item]; ok2 {
						result = append(result, item)
						continue
					}
				}
			}

			compartments = make([]string, 0)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var priority int
	for _, item := range result {
		priority += getPriority(item)
	}
	fmt.Println(priority)
}

func getPriority(str string) int {
	strRune := []rune(str)[0]
	if int(strRune) < 97 {
		return int(strRune) - 38
	} else {
		return int(strRune) - 96
	}
}

//func main() {
//	day3SolutionB()
//}
