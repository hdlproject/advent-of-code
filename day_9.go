package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day9SolutionA() {
	f, err := os.Open("day_9_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var mapTrack = make(map[int]map[int]struct{})
	var tailRow, tailCol int
	var headRow, headCol int

	mapTrack[tailCol] = make(map[int]struct{})
	mapTrack[tailCol][tailRow] = struct{}{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, " ")

		stepCount, _ := strconv.ParseInt(values[1], 10, 64)

		switch values[0] {
		case "U":
			for i := 0; i < int(stepCount); i++ {
				headCol++
				tailCol, tailRow = process2Knots(headCol, headRow, tailCol, tailRow, mapTrack)
			}
		case "D":
			for i := 0; i < int(stepCount); i++ {
				headCol--
				tailCol, tailRow = process2Knots(headCol, headRow, tailCol, tailRow, mapTrack)
			}
		case "R":
			for i := 0; i < int(stepCount); i++ {
				headRow++
				tailCol, tailRow = process2Knots(headCol, headRow, tailCol, tailRow, mapTrack)
			}
		case "L":
			for i := 0; i < int(stepCount); i++ {
				headRow--
				tailCol, tailRow = process2Knots(headCol, headRow, tailCol, tailRow, mapTrack)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var count int
	for _, row := range mapTrack {
		for range row {
			count++
		}
	}

	fmt.Println(count)
}

type knot struct {
	col int
	row int
}

func day9SolutionB() {
	f, err := os.Open("day_9_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var mapTrack = make(map[int]map[int]struct{})

	var knots []knot
	for i := 0; i < 10; i++ {
		knots = append(knots, knot{col: 0, row: 0})
	}

	mapTrack[knots[9].col] = make(map[int]struct{})
	mapTrack[knots[9].col][knots[9].row] = struct{}{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		values := strings.Split(value, " ")

		stepCount, _ := strconv.ParseInt(values[1], 10, 64)

		switch values[0] {
		case "U":
			for i := 0; i < int(stepCount); i++ {
				knots[0].col++
				process10Knots(knots, mapTrack)
			}
		case "D":
			for i := 0; i < int(stepCount); i++ {
				knots[0].col--
				process10Knots(knots, mapTrack)
			}
		case "R":
			for i := 0; i < int(stepCount); i++ {
				knots[0].row++
				process10Knots(knots, mapTrack)
			}
		case "L":
			for i := 0; i < int(stepCount); i++ {
				knots[0].row--
				process10Knots(knots, mapTrack)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var count int
	for _, row := range mapTrack {
		for range row {
			count++
		}
	}

	fmt.Println(count)
}

func moveTo(headCol, headRow, tailCol, tailRow int) (nextTailCol, nextTailRow int, isMoved bool) {
	nextTailCol = tailCol
	nextTailRow = tailRow

	if headCol-nextTailCol > 1 {
		isMoved = true
		nextTailCol++

		if headRow-nextTailRow > 1 {
			nextTailRow++
		} else if headRow-nextTailRow < -1 {
			nextTailRow--
		} else {
			nextTailRow = headRow
		}
	} else if headCol-nextTailCol < -1 {
		isMoved = true
		nextTailCol--

		if headRow-nextTailRow > 1 {
			nextTailRow++
		} else if headRow-nextTailRow < -1 {
			nextTailRow--
		} else {
			nextTailRow = headRow
		}
	} else if headRow-nextTailRow > 1 {
		isMoved = true
		nextTailRow++

		if headCol-nextTailCol > 1 {
			nextTailCol++
		} else if headCol-nextTailCol < -1 {
			nextTailCol--
		} else {
			nextTailCol = headCol
		}
	} else if headRow-nextTailRow < -1 {
		isMoved = true
		nextTailRow--

		if headCol-nextTailCol > 1 {
			nextTailCol++
		} else if headCol-nextTailCol < -1 {
			nextTailCol--
		} else {
			nextTailCol = headCol
		}
	}

	return nextTailCol, nextTailRow, isMoved
}

func process2Knots(headCol, headRow, tailCol, tailRow int, mapTrack map[int]map[int]struct{}) (newTailCol, newTailRow int) {
	var isMoved bool
	newTailCol, newTailRow, isMoved = moveTo(headCol, headRow, tailCol, tailRow)
	if isMoved {
		if _, ok := mapTrack[newTailCol]; !ok {
			mapTrack[newTailCol] = make(map[int]struct{})
		}
		mapTrack[newTailCol][newTailRow] = struct{}{}
	}

	return
}

func process10Knots(knots []knot, mapTrack map[int]map[int]struct{}) {
	var isMoved bool
	for j := 0; j < 9; j++ {
		knots[j+1].col, knots[j+1].row, isMoved = moveTo(knots[j].col, knots[j].row, knots[j+1].col, knots[j+1].row)

	}

	if isMoved {
		if _, ok := mapTrack[knots[9].col]; !ok {
			mapTrack[knots[9].col] = make(map[int]struct{})
		}
		mapTrack[knots[9].col][knots[9].row] = struct{}{}
	}
}

//func main() {
//	day9SolutionB()
//}
