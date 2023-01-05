package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	top    *tree
	bottom *tree
	right  *tree
	left   *tree
}

func newTree(top, bottom, right, left *tree) *tree {
	return &tree{
		top:    top,
		bottom: bottom,
		right:  right,
		left:   left,
	}
}

func day8SolutionA() {
	f, err := os.Open("day_8_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var graph [][]int
	var count int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, "")

		var row []int
		for _, item := range values {
			height, _ := strconv.ParseInt(item, 10, 64)
			row = append(row, int(height))
		}
		graph = append(graph, row)
	}

	for colIndex, row := range graph {
		for rowIndex := range row {
			if isVisibleFromTop(graph, rowIndex, colIndex) {
				//fmt.Println("top", rowIndex, colIndex, item)
				count++
			} else if isVisibleFromBottom(graph, rowIndex, colIndex) {
				//fmt.Println("bottom", rowIndex, colIndex, item)
				count++
			} else if isVisibleFromRight(graph, rowIndex, colIndex) {
				//fmt.Println("right", rowIndex, colIndex, item)
				count++
			} else if isVisibleFromLeft(graph, rowIndex, colIndex) {
				//fmt.Println("left", rowIndex, colIndex, item)
				count++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

func day8SolutionB() {
	f, err := os.Open("day_8_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var graph [][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, "")

		var row []int
		for _, item := range values {
			height, _ := strconv.ParseInt(item, 10, 64)
			row = append(row, int(height))
		}
		graph = append(graph, row)
	}

	var maxScore int64
	for colIndex, row := range graph {
		for rowIndex := range row {
			score := countVisibleFromTop(graph, rowIndex, colIndex) *
				countVisibleFromBottom(graph, rowIndex, colIndex) *
				countVisibleFromRight(graph, rowIndex, colIndex) *
				countVisibleFromLeft(graph, rowIndex, colIndex)

			maxScore = int64(math.Max(float64(maxScore), float64(score)))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(maxScore)
}

func isVisibleFromTop(graph [][]int, rowIndex, colIndex int) bool {
	for i := colIndex - 1; i >= 0; i-- {
		if graph[i][rowIndex] >= graph[colIndex][rowIndex] {
			return false
		}
	}

	return true
}

func countVisibleFromTop(graph [][]int, rowIndex, colIndex int) int {
	var count int
	for i := colIndex - 1; i >= 0; i-- {
		count++
		if graph[i][rowIndex] >= graph[colIndex][rowIndex] {
			break
		}
	}

	return count
}

func isVisibleFromBottom(graph [][]int, rowIndex, colIndex int) bool {
	for i := colIndex + 1; i < len(graph); i++ {
		if graph[i][rowIndex] >= graph[colIndex][rowIndex] {
			return false
		}
	}

	return true
}

func countVisibleFromBottom(graph [][]int, rowIndex, colIndex int) int {
	var count int
	for i := colIndex + 1; i < len(graph); i++ {
		count++
		if graph[i][rowIndex] >= graph[colIndex][rowIndex] {
			break
		}
	}

	return count
}

func isVisibleFromRight(graph [][]int, rowIndex, colIndex int) bool {
	for i := rowIndex + 1; i < len(graph[0]); i++ {
		if graph[colIndex][i] >= graph[colIndex][rowIndex] {
			return false
		}
	}

	return true
}

func countVisibleFromRight(graph [][]int, rowIndex, colIndex int) int {
	var count int
	for i := rowIndex + 1; i < len(graph[0]); i++ {
		count++
		if graph[colIndex][i] >= graph[colIndex][rowIndex] {
			break
		}
	}

	return count
}

func isVisibleFromLeft(graph [][]int, rowIndex, colIndex int) bool {
	for i := rowIndex - 1; i >= 0; i-- {
		if graph[colIndex][i] >= graph[colIndex][rowIndex] {
			return false
		}
	}

	return true
}

func countVisibleFromLeft(graph [][]int, rowIndex, colIndex int) int {
	var count int
	for i := rowIndex - 1; i >= 0; i-- {
		count++
		if graph[colIndex][i] >= graph[colIndex][rowIndex] {
			break
		}
	}

	return count
}

//func main() {
//	day8SolutionB()
//}
