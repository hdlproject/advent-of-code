package main

import "fmt"

func getConnectedSide(colIndex, rowIndex int, direction string) (string, string, int, int) {
	sideMap := map[string]map[string]struct {
		side      string
		facing    string
		wrapIndex func(colIndex, rowIndex int) (int, int)
	}{
		"A": {
			"T": {
				side:   "B",
				facing: "T",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					return colIndex - 1, rowIndex
				},
			},
			"B": {
				side:   "D",
				facing: "B",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					return colIndex + 1, rowIndex
				},
			},
			"R": {
				side:   "C",
				facing: "T",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					currStartCol, _, _, _ := getSideRange("A")
					_, startRow, finishCol, _ := getSideRange("C")
					return finishCol, startRow + (colIndex - currStartCol)
				},
			},
			"L": {
				side:   "E",
				facing: "B",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					currStartCol, _, _, _ := getSideRange("A")
					startCol, startRow, _, _ := getSideRange("E")
					return startCol, startRow + (colIndex - currStartCol)
				},
			},
		},
		"B": {
			"T": {
				side:   "F",
				facing: "R",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					_, currStartRow, _, _ := getSideRange("B")
					startCol, startRow, _, _ := getSideRange("F")
					return startCol + (rowIndex - currStartRow), startRow
				},
			},
			"B": {
				side:   "A",
				facing: "B",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					return colIndex + 1, rowIndex
				},
			},
			"R": {
				side:   "C",
				facing: "R",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					return colIndex, rowIndex + 1
				},
			},
			"L": {
				side:   "E",
				facing: "R",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					_, _, currFinishCol, _ := getSideRange("B")
					startCol, startRow, _, _ := getSideRange("E")
					return startCol + (currFinishCol - colIndex), startRow
				},
			},
		},
		"C": {
			"T": {
				side:   "F",
				facing: "T",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					_, currStartRow, _, _ := getSideRange("C")
					_, startRow, finishCol, _ := getSideRange("F")
					return finishCol, startRow + (rowIndex - currStartRow)
				},
			},
			"B": {
				side:   "A",
				facing: "L",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					_, currStartRow, _, _ := getSideRange("C")
					startCol, _, _, finishRow := getSideRange("A")
					return startCol + (rowIndex - currStartRow), finishRow
				},
			},
			"R": {
				side:   "D",
				facing: "L",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					_, _, currFinishCol, _ := getSideRange("C")
					startCol, _, _, finishRow := getSideRange("D")
					return startCol + (currFinishCol - colIndex), finishRow
				},
			},
			"L": {
				side:   "B",
				facing: "L",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					return colIndex, rowIndex - 1
				},
			},
		},
		"D": {
			"T": {
				side:   "A",
				facing: "T",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					return colIndex - 1, rowIndex
				},
			},
			"B": {
				side:   "F",
				facing: "L",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					_, currStartRow, _, _ := getSideRange("D")
					startCol, _, _, finishRow := getSideRange("F")
					return startCol + (rowIndex - currStartRow), finishRow
				},
			},
			"R": {
				side:   "C",
				facing: "L",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					_, _, currFinishCol, _ := getSideRange("D")
					startCol, _, _, finishRow := getSideRange("C")
					return startCol + (currFinishCol - colIndex), finishRow
				},
			},
			"L": {
				side:   "E",
				facing: "L",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					return colIndex, rowIndex - 1
				},
			},
		},
		"E": {
			"T": {
				side:   "A",
				facing: "R",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					_, currStartRow, _, _ := getSideRange("E")
					startCol, startRow, _, _ := getSideRange("A")
					return startCol + (rowIndex - currStartRow), startRow
				},
			},
			"B": {
				side:   "F",
				facing: "B",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					return colIndex + 1, rowIndex
				},
			},
			"R": {
				side:   "D",
				facing: "R",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					return colIndex, rowIndex + 1
				},
			},
			"L": {
				side:   "B",
				facing: "R",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					_, _, currFinishCol, _ := getSideRange("E")
					startCol, startRow, _, _ := getSideRange("B")
					return startCol + (currFinishCol - colIndex), startRow
				},
			},
		},
		"F": {
			"T": {
				side:   "E",
				facing: "T",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					return colIndex - 1, rowIndex
				},
			},
			"B": {
				side:   "C",
				facing: "B",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					_, currStartRow, _, _ := getSideRange("F")
					startCol, startRow, _, _ := getSideRange("C")
					return startCol, startRow + (rowIndex - currStartRow)
				},
			},
			"R": {
				side:   "D",
				facing: "T",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					currStartCol, _, _, _ := getSideRange("F")
					_, startRow, finishCol, _ := getSideRange("D")
					return finishCol, startRow + (colIndex - currStartCol)
				},
			},
			"L": {
				side:   "B",
				facing: "B",
				wrapIndex: func(colIndex, rowIndex int) (int, int) {
					currStartCol, _, _, _ := getSideRange("F")
					startCol, startRow, _, _ := getSideRange("B")
					return startCol, startRow + (colIndex - currStartCol)
				},
			},
		},
	}

	currSide := getSide(colIndex, rowIndex)

	nextCol, nextRow := sideMap[currSide][direction].wrapIndex(colIndex, rowIndex)

	return sideMap[currSide][direction].side, sideMap[currSide][direction].facing, nextCol, nextRow
}

func getSide(colIndex, rowIndex int) string {
	sideList := []string{"A", "B", "C", "D", "E", "F"}

	for _, side := range sideList {
		startCol, startRow, finishCol, finishRow := getSideRange(side)

		if colIndex >= startCol && colIndex <= finishCol &&
			rowIndex >= startRow && rowIndex <= finishRow {

			return side
		}
	}

	return ""
}

func getSideRange(side string) (int, int, int, int) {
	sideRanges := map[string]struct {
		startCol  int
		startRow  int
		finishCol int
		finishRow int
	}{
		"A": {
			startCol:  50,
			startRow:  50,
			finishCol: 99,
			finishRow: 99,
		},
		"B": {
			startCol:  0,
			startRow:  50,
			finishCol: 49,
			finishRow: 99,
		},
		"C": {
			startCol:  0,
			startRow:  100,
			finishCol: 49,
			finishRow: 149,
		},
		"D": {
			startCol:  100,
			startRow:  50,
			finishCol: 149,
			finishRow: 99,
		},
		"E": {
			startCol:  100,
			startRow:  0,
			finishCol: 149,
			finishRow: 49,
		},
		"F": {
			startCol:  150,
			startRow:  0,
			finishCol: 199,
			finishRow: 49,
		},
	}

	return sideRanges[side].startCol,
		sideRanges[side].startRow,
		sideRanges[side].finishCol,
		sideRanges[side].finishRow
}

func getDirection(currDirection, direction string) string {
	directionMap := map[string]map[string]string{
		"T": {
			"R": "R",
			"L": "L",
		},
		"B": {
			"R": "L",
			"L": "R",
		},
		"R": {
			"R": "B",
			"L": "T",
		},
		"L": {
			"R": "T",
			"L": "B",
		},
	}

	return directionMap[currDirection][direction]
}

func getFacingPoint(facing string) int {
	switch facing {
	case "R":
		return 0
	case "B":
		return 1
	case "L":
		return 2
	case "T":
		return 3
	}

	return 0
}

func printTilesMap(tilesMap [][]*tile, start *tile, stepNumbers []int, stepDirections []string) {
	for _, row := range tilesMap {
		for _, item := range row {
			if item == start {
				fmt.Print("S")
				continue
			}

			fmt.Print(item.char)
		}
		fmt.Println("|")
	}

	for index, direction := range stepDirections {
		fmt.Print(direction)
		fmt.Print(stepNumbers[index])
	}
	fmt.Println()
}

func printTileSide(tilesMap [][]*tile) {
	for colIndex, row := range tilesMap {
		for rowIndex, item := range row {
			if item.char != " " {
				fmt.Print(getSide(colIndex, rowIndex))
				continue
			}

			fmt.Print(item.char)
		}

		fmt.Println()
	}
}

func isStepNumber(char string) bool {
	return char != "R" && char != "L"
}
