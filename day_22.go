package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type tile struct {
	char   string
	top    *tile
	bottom *tile
	right  *tile
	left   *tile

	colIndex int
	rowIndex int
}

func day22SolutionA() {
	f, err := os.Open("day_22_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var scanStep bool
	var length int

	var tilesMap [][]*tile
	var start *tile

	var stepNumbers []int
	stepDirections := []string{"R"}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		if value == "" {
			scanStep = true
			continue
		}

		values := strings.Split(value, "")

		if !scanStep {
			if len(values) > length {
				for colIndex := range tilesMap {
					for i := length; i < len(values); i++ {
						tilesMap[colIndex] = append(tilesMap[colIndex], &tile{
							char: " ",
						})
					}
				}

				length = len(values)
			}

			var tiles []*tile
			for _, char := range values {
				newTile := &tile{
					char: char,
				}

				if start == nil && char != " " {
					start = newTile
				}

				tiles = append(tiles, newTile)
			}

			for i := len(values); i < length; i++ {
				tiles = append(tiles, &tile{
					char: " ",
				})
			}

			tilesMap = append(tilesMap, tiles)
		} else {
			var numberStr string
			for _, char := range values {
				if isStepNumber(char) {
					numberStr += char
				} else {
					if numberStr != "" {
						numberInt, _ := strconv.ParseInt(numberStr, 10, 64)
						stepNumbers = append(stepNumbers, int(numberInt))
						numberStr = ""
					}

					stepDirections = append(stepDirections, char)
				}
			}

			numberInt, _ := strconv.ParseInt(numberStr, 10, 64)
			stepNumbers = append(stepNumbers, int(numberInt))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tilesMap = buildTileGraph(tilesMap)

	finish, facing := traverseTile(stepNumbers, stepDirections, start)

	fmt.Println(((finish.colIndex + 1) * 1000) + ((finish.rowIndex + 1) * 4) + getFacingPoint(facing))
}

func traverseTile(stepNumbers []int, stepDirections []string, currTile *tile) (*tile, string) {
	currDirection := "T"
	for index, direction := range stepDirections {
		currDirection = getDirection(currDirection, direction)

		switch currDirection {
		case "T":
			for i := 0; i < stepNumbers[index]; i++ {
				if currTile.top.char != "#" {
					currTile = currTile.top
				}
			}
		case "B":
			for i := 0; i < stepNumbers[index]; i++ {
				if currTile.bottom.char != "#" {
					currTile = currTile.bottom
				}
			}
		case "R":
			for i := 0; i < stepNumbers[index]; i++ {
				if currTile.right.char != "#" {
					currTile = currTile.right
				}
			}
		case "L":
			for i := 0; i < stepNumbers[index]; i++ {
				if currTile.left.char != "#" {
					currTile = currTile.left
				}
			}
		}
	}

	return currTile, currDirection
}

func buildTileGraph(tilesMap [][]*tile) [][]*tile {
	for colIndex, row := range tilesMap {
		for rowIndex, item := range row {
			if item.char != " " {
				var top *tile
				if colIndex-1 >= 0 && tilesMap[colIndex-1][rowIndex].char != " " {
					top = tilesMap[colIndex-1][rowIndex]
				} else {
					top = tilesMap[findMostBottomTile(tilesMap, rowIndex)][rowIndex]
				}

				var bottom *tile
				if colIndex+1 < len(tilesMap) && tilesMap[colIndex+1][rowIndex].char != " " {
					bottom = tilesMap[colIndex+1][rowIndex]
				} else {
					bottom = tilesMap[findMostTopTile(tilesMap, rowIndex)][rowIndex]
				}

				var right *tile
				if rowIndex+1 < len(tilesMap[0]) && tilesMap[colIndex][rowIndex+1].char != " " {
					right = tilesMap[colIndex][rowIndex+1]
				} else {
					right = tilesMap[colIndex][findMostLeftTile(tilesMap, colIndex)]
				}

				var left *tile
				if rowIndex-1 >= 0 && tilesMap[colIndex][rowIndex-1].char != " " {
					left = tilesMap[colIndex][rowIndex-1]
				} else {
					left = tilesMap[colIndex][findMostRightTile(tilesMap, colIndex)]
				}

				item.top = top
				item.bottom = bottom
				item.right = right
				item.left = left
				item.colIndex = colIndex
				item.rowIndex = rowIndex

				tilesMap[colIndex][rowIndex] = item
			}
		}
	}

	return tilesMap
}

func findMostBottomTile(tilesMap [][]*tile, rowIndex int) int {
	for i := len(tilesMap) - 1; i >= 0; i-- {
		if tilesMap[i][rowIndex].char != " " {
			return i
		}
	}

	return 0
}

func findMostTopTile(tilesMap [][]*tile, rowIndex int) int {
	for i := 0; i < len(tilesMap); i++ {
		if tilesMap[i][rowIndex].char != " " {
			return i
		}
	}

	return 0
}

func findMostLeftTile(tilesMap [][]*tile, colIndex int) int {
	for i := 0; i < len(tilesMap[0]); i++ {
		if tilesMap[colIndex][i].char != " " {
			return i
		}
	}

	return 0
}

func findMostRightTile(tilesMap [][]*tile, colIndex int) int {
	for i := len(tilesMap[0]) - 1; i >= 0; i-- {
		if tilesMap[colIndex][i].char != " " {
			return i
		}
	}

	return 0
}

func day22SolutionB() {
	f, err := os.Open("day_22_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var scanStep bool
	var length int

	var tilesMap [][]*tile
	var start *tile

	var stepNumbers []int
	stepDirections := []string{"R"}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		if value == "" {
			scanStep = true
			continue
		}

		values := strings.Split(value, "")

		if !scanStep {
			if len(values) > length {
				for colIndex := range tilesMap {
					for i := length; i < len(values); i++ {
						tilesMap[colIndex] = append(tilesMap[colIndex], &tile{
							char: " ",
						})
					}
				}

				length = len(values)
			}

			var tiles []*tile
			for _, char := range values {
				newTile := &tile{
					char: char,
				}

				if start == nil && char != " " {
					start = newTile
				}

				tiles = append(tiles, newTile)
			}

			for i := len(values); i < length; i++ {
				tiles = append(tiles, &tile{
					char: " ",
				})
			}

			tilesMap = append(tilesMap, tiles)
		} else {
			var numberStr string
			for _, char := range values {
				if isStepNumber(char) {
					numberStr += char
				} else {
					if numberStr != "" {
						numberInt, _ := strconv.ParseInt(numberStr, 10, 64)
						stepNumbers = append(stepNumbers, int(numberInt))
						numberStr = ""
					}

					stepDirections = append(stepDirections, char)
				}
			}

			numberInt, _ := strconv.ParseInt(numberStr, 10, 64)
			stepNumbers = append(stepNumbers, int(numberInt))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tilesMap = buildTileGraphType2(tilesMap)

	finish, facing := traverseTileType2(tilesMap, stepNumbers, stepDirections, start)

	fmt.Println(((finish.colIndex + 1) * 1000) + ((finish.rowIndex + 1) * 4) + getFacingPoint(facing))
}

func buildTileGraphType2(tilesMap [][]*tile) [][]*tile {
	for colIndex, row := range tilesMap {
		for rowIndex, item := range row {
			if item.char != " " {
				item.colIndex = colIndex
				item.rowIndex = rowIndex

				tilesMap[colIndex][rowIndex] = item
			}
		}
	}

	return tilesMap
}

func traverseTileType2(tilesMap [][]*tile, stepNumbers []int, stepDirections []string, currTile *tile) (*tile, string) {
	currDirection := "T"
	for index, direction := range stepDirections {
		currDirection = getDirection(currDirection, direction)

		for i := 0; i < stepNumbers[index]; i++ {
			switch currDirection {
			case "T":
				if currTile.colIndex-1 >= 0 &&
					tilesMap[currTile.colIndex-1][currTile.rowIndex].char != " " {

					if tilesMap[currTile.colIndex-1][currTile.rowIndex].char != "#" {
						currTile = tilesMap[currTile.colIndex-1][currTile.rowIndex]
					}
					continue
				}
			case "B":
				if currTile.colIndex+1 < len(tilesMap) &&
					tilesMap[currTile.colIndex+1][currTile.rowIndex].char != " " {

					if tilesMap[currTile.colIndex+1][currTile.rowIndex].char != "#" {
						currTile = tilesMap[currTile.colIndex+1][currTile.rowIndex]
					}
					continue
				}
			case "R":
				if currTile.rowIndex+1 < len(tilesMap[0]) &&
					tilesMap[currTile.colIndex][currTile.rowIndex+1].char != " " {

					if tilesMap[currTile.colIndex][currTile.rowIndex+1].char != "#" {
						currTile = tilesMap[currTile.colIndex][currTile.rowIndex+1]
					}
					continue
				}
			case "L":
				if currTile.rowIndex-1 >= 0 &&
					tilesMap[currTile.colIndex][currTile.rowIndex-1].char != " " {

					if tilesMap[currTile.colIndex][currTile.rowIndex-1].char != "#" {
						currTile = tilesMap[currTile.colIndex][currTile.rowIndex-1]
					}
					continue
				}
			}

			nextDirection, nextCol, nextRow := findWrap(currTile.colIndex, currTile.rowIndex, currDirection)
			if tilesMap[nextCol][nextRow].char != "#" {
				currTile = tilesMap[nextCol][nextRow]
				currDirection = nextDirection
			}
		}
	}

	return currTile, currDirection
}

func findWrap(colIndex, rowIndex int, direction string) (string, int, int) {
	_, nextFacing, nextCol, nextRow := getConnectedSide(colIndex, rowIndex, direction)

	return nextFacing, nextCol, nextRow
}

//func main() {
//	day22SolutionB()
//}
