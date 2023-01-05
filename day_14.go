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

type pointType string

const air pointType = "air"
const rock pointType = "rock"
const sand pointType = "sand"

type cavePoint struct {
	pointType pointType
}

func day14SolutionA() {
	f, err := os.Open("day_14_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cave [][]*cavePoint
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		values := strings.Split(value, "->")
		x1, y1 := -1, -1
		for _, pointVal := range values {
			point := strings.Split(pointVal, ",")
			x, _ := strconv.ParseInt(strings.Trim(point[0], " "), 10, 64)
			y, _ := strconv.ParseInt(strings.Trim(point[1], " "), 10, 64)

			cave = constructCave(x1, y1, int(x), int(y), cave)

			x1 = int(x)
			y1 = int(y)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var index int
	for {
		isStop := putSand(500, 0, cave)
		if isStop {
			fmt.Println(index)
			break
		}
		index++
	}

	printCave(cave)
}

func constructCave(x1, y1, x2, y2 int, cave [][]*cavePoint) [][]*cavePoint {
	y := int(math.Max(float64(y1), float64(y2)))
	x := int(math.Max(float64(x1), float64(x2)))

	if y >= len(cave) {
		for i := len(cave) - 1; i < y; i++ {
			var additionalPoints []*cavePoint
			if len(cave) > 0 {
				for j := 0; j < len(cave[0]); j++ {
					additionalPoints = append(additionalPoints, &cavePoint{pointType: air})
				}
			}
			cave = append(cave, additionalPoints)
		}
	}

	if x >= len(cave[0]) {
		for i := len(cave[0]) - 1; i < x; i++ {
			for j := 0; j < len(cave); j++ {
				cave[j] = append(cave[j], &cavePoint{pointType: air})
			}
		}
	}

	if x1 == -1 {
		x1 = x2
	}
	if y1 == -1 {
		y1 = y2
	}
	minX := int(math.Min(float64(x1), float64(x2)))
	maxX := int(math.Max(float64(x1), float64(x2)))
	minY := int(math.Min(float64(y1), float64(y2)))
	maxY := int(math.Max(float64(y1), float64(y2)))

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			cave[i][j].pointType = rock
		}
	}

	return cave
}

func additionalFloor(cave [][]*cavePoint) [][]*cavePoint {
	for i := 0; i < len(cave); i++ {
		for j := 0; j < len(cave)+2; j++ {
			cave[i] = append(cave[i], &cavePoint{pointType: air})
		}
	}

	var floor []*cavePoint
	for i := 0; i < len(cave[0]); i++ {
		floor = append(floor, &cavePoint{pointType: air})
	}
	cave = append(cave, floor)

	floor = []*cavePoint{}
	for i := 0; i < len(cave[0]); i++ {
		floor = append(floor, &cavePoint{pointType: rock})
	}
	cave = append(cave, floor)

	return cave
}

func putSand(x, y int, cave [][]*cavePoint) (isStop bool) {
	if (y + 1) < (len(cave)) {
		if cave[y+1][x].pointType == air {
			return putSand(x, y+1, cave)
		} else {
			if (x-1) > 0 && cave[y+1][x-1].pointType == air {
				return putSand(x-1, y+1, cave)
			} else if (x+1) < len(cave[0]) && cave[y+1][x+1].pointType == air {
				return putSand(x+1, y+1, cave)
			} else {
				cave[y][x].pointType = sand
				return false
			}
		}
	}

	return true
}

func putSandType2(x, y int, cave [][]*cavePoint) (isStop bool) {
	if (y + 1) < (len(cave)) {
		if cave[y+1][x].pointType == air {
			return putSandType2(x, y+1, cave)
		} else {
			if (x-1) > 0 && cave[y+1][x-1].pointType == air {
				return putSandType2(x-1, y+1, cave)
			} else if (x+1) < len(cave[0]) && cave[y+1][x+1].pointType == air {
				return putSandType2(x+1, y+1, cave)
			} else {
				cave[y][x].pointType = sand
				if y == 0 && x == 500 {
					return true
				}

				return false
			}
		}
	}

	return true
}

func printCave(cave [][]*cavePoint) {
	for _, row := range cave {
		for _, point := range row {
			if point.pointType == rock {
				fmt.Print("#")
			} else if point.pointType == sand {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func day14SolutionB() {
	f, err := os.Open("day_14_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cave [][]*cavePoint
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		values := strings.Split(value, "->")
		x1, y1 := -1, -1
		for _, pointVal := range values {
			point := strings.Split(pointVal, ",")
			x, _ := strconv.ParseInt(strings.Trim(point[0], " "), 10, 64)
			y, _ := strconv.ParseInt(strings.Trim(point[1], " "), 10, 64)

			cave = constructCave(x1, y1, int(x), int(y), cave)

			x1 = int(x)
			y1 = int(y)
		}
	}

	cave = additionalFloor(cave)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var index int
	for {
		isStop := putSandType2(500, 0, cave)
		if isStop {
			fmt.Println(index + 1)
			break
		}
		index++
	}

	printCave(cave)
}

//func main() {
//	day14SolutionB()
//}
