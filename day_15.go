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

type sbPoint struct {
	x         int
	y         int
	pX        int
	pY        int
	isSensor  bool
	mDistance int
}

type blockedRange struct {
	min int
	max int
}

func (s *sbPoint) isInRange(x, y int) bool {
	return calculateManhattanDistance(s.x, s.y, x, y) <= s.mDistance
}

func day15SolutionA() {
	f, err := os.Open("day_15_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sbPos := make(map[int]map[int]*sbPoint)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		values := strings.Split(value, " ")
		xS, _ := strconv.ParseInt(strings.TrimSuffix(strings.TrimPrefix(values[2], "x="), ","), 10, 64)
		yS, _ := strconv.ParseInt(strings.TrimSuffix(strings.TrimPrefix(values[3], "y="), ":"), 10, 64)
		xB, _ := strconv.ParseInt(strings.TrimSuffix(strings.TrimPrefix(values[8], "x="), ","), 10, 64)
		yB, _ := strconv.ParseInt(strings.TrimSuffix(strings.TrimPrefix(values[9], "y="), " "), 10, 64)

		if _, ok := sbPos[int(yS)]; !ok {
			sbPos[int(yS)] = make(map[int]*sbPoint)
		}
		sbPos[int(yS)][int(xS)] = &sbPoint{x: int(xS), y: int(yS), pX: int(xB), pY: int(yB), isSensor: true, mDistance: calculateManhattanDistance(int(xS), int(yS), int(xB), int(yB))}

		if _, ok := sbPos[int(yB)]; !ok {
			sbPos[int(yB)] = make(map[int]*sbPoint)
		}
		sbPos[int(yB)][int(xB)] = &sbPoint{x: int(xB), y: int(yB), pX: int(xS), pY: int(yS)}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	impossiblePos := make(map[int]struct{})
	for _, row := range sbPos {
		for _, point := range row {
			if point.isSensor {
				findRange(2000000, impossiblePos, point.x, point.y, point.mDistance, sbPos)
			}
		}
	}
	var total int
	for range impossiblePos {
		total++
	}

	fmt.Println(total)
}

func findRange(location int, impossiblePos map[int]struct{}, xS, yS, distance int, sbPos map[int]map[int]*sbPoint) {
	var horizontalRange int
	if yS < location {
		horizontalRange = distance - (location - yS)
	} else {
		horizontalRange = distance - (yS - location)
	}

	if horizontalRange < 0 {
		return
	}

	for i := xS - horizontalRange; i <= xS+horizontalRange; i++ {
		if _, ok := sbPos[location][i]; !ok {
			impossiblePos[i] = struct{}{}
		}
	}
}

func findRangeType1(location int, blockedRanges []blockedRange, xS, yS, distance, limit int) []blockedRange {
	var horizontalRange int
	if yS < location {
		horizontalRange = distance - (location - yS)
	} else {
		horizontalRange = distance - (yS - location)
	}

	if horizontalRange < 0 {
		return blockedRanges
	}

	left := int(math.Max(float64(xS-horizontalRange), 0))
	right := int(math.Min(float64(xS+horizontalRange), float64(limit)))

	blockedRanges = append(blockedRanges, blockedRange{
		min: left,
		max: right,
	})

	return blockedRanges
}

func compileBlockedRange(blockedRanges []blockedRange) blockedRange {
	if len(blockedRanges) == 0 {
		return blockedRange{}
	}

	newBlockedRange := blockedRanges[0]
	for range blockedRanges[1:] {
		for _, range2 := range blockedRanges[1:] {
			newBlockedRange, _ = mergeRange(newBlockedRange, range2)
		}
	}

	return newBlockedRange
}

func mergeRange(range1, range2 blockedRange) (blockedRange, bool) {
	if (range2.min - 1) > range1.max {
		return range1, false
	} else if (range2.max + 1) < range1.min {
		return range1, false
	}

	return blockedRange{
		min: int(math.Min(float64(range1.min), float64(range2.min))),
		max: int(math.Max(float64(range1.max), float64(range2.max))),
	}, true
}

func calculateManhattanDistance(xS, yS, xB, yB int) int {
	return int(math.Abs(float64(xS-xB)) + math.Abs(float64(yS-yB)))
}

func day15SolutionB() {
	f, err := os.Open("day_15_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var impossiblePos []*sbPoint
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		values := strings.Split(value, " ")
		xS, _ := strconv.ParseInt(strings.TrimSuffix(strings.TrimPrefix(values[2], "x="), ","), 10, 64)
		yS, _ := strconv.ParseInt(strings.TrimSuffix(strings.TrimPrefix(values[3], "y="), ":"), 10, 64)
		xB, _ := strconv.ParseInt(strings.TrimSuffix(strings.TrimPrefix(values[8], "x="), ","), 10, 64)
		yB, _ := strconv.ParseInt(strings.TrimSuffix(strings.TrimPrefix(values[9], "y="), " "), 10, 64)

		impossiblePos = append(impossiblePos, &sbPoint{
			x: int(xS), y: int(yS),
			pX: int(xB), pY: int(yB),
			isSensor:  true,
			mDistance: calculateManhattanDistance(int(xS), int(yS), int(xB), int(yB)),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	limit := 4000000
	for i := 0; i <= limit; i++ {
		var blockedRanges []blockedRange
		for _, point := range impossiblePos {
			blockedRanges = findRangeType1(i, blockedRanges, point.x, point.y, point.mDistance, limit)
		}
		realRange := compileBlockedRange(blockedRanges)
		if realRange.max-realRange.min >= limit {
			continue
		}

		if realRange.max >= limit {
			fmt.Println(((realRange.min - 1) * 4000000) + i)
			break
		} else {
			fmt.Println(((realRange.max + 1) * 4000000) + i)
			break
		}
	}
}

//func main() {
//	day15SolutionB()
//}
