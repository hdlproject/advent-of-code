package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day2SolutionA() {
	f, err := os.Open("day_2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	codeMap := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	var score int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, " ")

		enemy := values[0]
		me := codeMap[values[1]]

		score += baseScore(me)

		if enemy == me {
			score += 3
			continue
		}

		if me == "A" && enemy == "C" {
			score += 6
		} else if me == "B" && enemy == "A" {
			score += 6
		} else if me == "C" && enemy == "B" {
			score += 6
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(score)
}

func day2SolutionB() {
	f, err := os.Open("day_2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	codeMap := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	var score int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, " ")

		enemy := values[0]
		me := codeMap[values[1]]

		score = score + baseScore(getChosen(me, enemy)) + getScore(me)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(score)
}

func baseScore(chosen string) int {
	switch chosen {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	default:
		return 0
	}
}

func getScore(chosen string) int {
	switch chosen {
	case "A":
		return 0
	case "B":
		return 3
	case "C":
		return 6
	default:
		return 0
	}
}

func getChosen(me, enemy string) string {
	switch me {
	case "A":
		switch enemy {
		case "A":
			return "C"
		case "B":
			return "A"
		case "C":
			return "B"
		}
	case "B":
		return enemy
	case "C":
		switch enemy {
		case "A":
			return "B"
		case "B":
			return "C"
		case "C":
			return "A"
		}
	}

	return me
}

//func main() {
//	day2SolutionB()
//}
