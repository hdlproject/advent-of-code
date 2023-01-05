package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day10SolutionA() {
	f, err := os.Open("day_10_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	registries := make(map[int]int)
	cycle := 0
	registry := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, " ")

		var addition int64
		switch values[0] {
		case "noop":
			cycle++
		case "addx":
			cycle += 2

			addition, _ = strconv.ParseInt(values[1], 10, 64)
		}

		recordSignalStrength(cycle, registry, registries)
		registry += int(addition)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var total int
	for _, strength := range registries {
		total += strength
	}

	fmt.Println(total)
}

func day10SolutionB() {
	f, err := os.Open("day_10_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cycle := 0
	registry := 1
	minSprite := 0
	maxSprite := 2
	var buffer []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, " ")

		var addition int64
		switch values[0] {
		case "noop":
			buffer = append(buffer, determineDrawing(cycle, minSprite, maxSprite))
			cycle++
		case "addx":
			buffer = append(buffer, determineDrawing(cycle, minSprite, maxSprite))
			cycle++
			buffer = append(buffer, determineDrawing(cycle, minSprite, maxSprite))
			cycle++
			addition, _ = strconv.ParseInt(values[1], 10, 64)
		}

		registry += int(addition)
		minSprite, maxSprite = determineSprite(registry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			fmt.Print(buffer[(i*40)+j])
		}
		fmt.Println()
	}
}

func recordSignalStrength(cycle, registry int, registries map[int]int) {
	checkPoints := []int{20, 60, 100, 140, 180, 220}
	for _, point := range checkPoints {
		if _, ok := registries[point]; !ok && cycle >= point {
			registries[point] = registry * point
		}
	}
}

func determineSprite(registry int) (min, max int) {
	return registry - 1, registry + 1
}

func determineDrawing(cycle, minSprite, maxSprite int) string {
	cycle = cycle % 40
	if cycle >= minSprite && cycle <= maxSprite {
		return "#"
	} else {
		return "."
	}
}

//func main() {
//	day10SolutionB()
//}
