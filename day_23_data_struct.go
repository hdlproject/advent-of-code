package main

import (
	"fmt"
	"math"
)

type elf struct {
	y  int
	x  int
	id string
}

type fieldTile struct {
	elf *elf
}

type field struct {
	colLen           int
	rowLen           int
	priority         []string
	elves            []*elf
	tiles            [][]fieldTile
	proposedMovement map[int]map[int][]*elf
	round            int
}

func newField() *field {
	return &field{
		priority:         []string{"N", "S", "W", "E"},
		proposedMovement: make(map[int]map[int][]*elf),
		round:            1,
	}
}

func (f *field) needMove(e *elf) bool {
	if e.y-1 < 0 {
		f.appendFieldTop()
	}
	if e.y+1 >= f.colLen {
		f.appendFieldBottom()
	}
	if e.x-1 < 0 {
		f.appendFieldLeft()
	}
	if e.x+1 >= f.rowLen {
		f.appendFieldRight()
	}

	for i := e.y - 1; i <= e.y+1; i++ {
		for j := e.x - 1; j <= e.x+1; j++ {
			if i == e.y && j == e.x {
				continue
			}

			if f.tiles[i][j].elf != nil {
				return true
			}
		}
	}

	return false
}

func (f *field) appendFieldTop() {
	var newField []fieldTile
	for i := 0; i < f.rowLen; i++ {
		newField = append(newField, fieldTile{})
	}

	f.tiles = append([][]fieldTile{newField}, f.tiles...)
	f.colLen++

	for _, currElf := range f.elves {
		currElf.y++
	}
}

func (f *field) appendFieldBottom() {
	var newField []fieldTile
	for i := 0; i < f.rowLen; i++ {
		newField = append(newField, fieldTile{})
	}

	f.tiles = append(f.tiles, newField)
	f.colLen++
}

func (f *field) appendFieldLeft() {
	for i := 0; i < f.colLen; i++ {
		f.tiles[i] = append([]fieldTile{{}}, f.tiles[i]...)
	}

	f.rowLen++

	for _, currElf := range f.elves {
		currElf.x++
	}
}

func (f *field) appendFieldRight() {
	for i := 0; i < f.colLen; i++ {
		f.tiles[i] = append(f.tiles[i], fieldTile{})
	}

	f.rowLen++
}

func (f *field) whereToMove(e *elf) string {
	for _, item := range f.priority {
		switch item {
		case "N":
			if f.canMoveNorth(e) {
				return item
			}
		case "S":
			if f.canMoveSouth(e) {
				return item
			}
		case "W":
			if f.canMoveWest(e) {
				return item
			}
		case "E":
			if f.canMoveEast(e) {
				return item
			}
		}
	}

	return ""
}

func (f *field) canMoveNorth(e *elf) bool {
	//if e.y-1 < 0 {
	//	f.appendFieldTop()
	//}
	//
	//if e.x-1 < 0 {
	//	f.appendFieldLeft()
	//}
	//if e.x+1 >= f.rowLen {
	//	f.appendFieldRight()
	//}

	for i := e.x - 1; i <= e.x+1; i++ {
		if f.tiles[e.y-1][i].elf != nil {
			return false
		}
	}

	return true
}

func (f *field) canMoveSouth(e *elf) bool {
	//if e.y+1 >= f.colLen {
	//	f.appendFieldBottom()
	//}
	//
	//if e.x-1 < 0 {
	//	f.appendFieldLeft()
	//}
	//if e.x+1 >= f.rowLen {
	//	f.appendFieldRight()
	//}

	for i := e.x - 1; i <= e.x+1; i++ {
		if f.tiles[e.y+1][i].elf != nil {
			return false
		}
	}

	return true
}

func (f *field) canMoveWest(e *elf) bool {
	//if e.x-1 < 0 {
	//	f.appendFieldLeft()
	//}
	//
	//if e.y-1 < 0 {
	//	f.appendFieldTop()
	//}
	//if e.y+1 >= f.colLen {
	//	f.appendFieldBottom()
	//}

	for i := e.y - 1; i <= e.y+1; i++ {
		if f.tiles[i][e.x-1].elf != nil {
			return false
		}
	}

	return true
}

func (f *field) canMoveEast(e *elf) bool {
	//if e.x+1 >= f.rowLen {
	//	f.appendFieldRight()
	//}
	//
	//if e.y-1 < 0 {
	//	f.appendFieldTop()
	//}
	//if e.y+1 >= f.colLen {
	//	f.appendFieldBottom()
	//}

	for i := e.y - 1; i <= e.y+1; i++ {
		if f.tiles[i][e.x+1].elf != nil {
			return false
		}
	}

	return true
}

func (f *field) planMove() {
	var needMoveElves []*elf
	for _, currElf := range f.elves {
		if f.needMove(currElf) {
			needMoveElves = append(needMoveElves, currElf)
		}
	}

	for _, currElf := range needMoveElves {
		proposedY := currElf.y
		proposedX := currElf.x

		whereToMove := f.whereToMove(currElf)
		switch whereToMove {
		case "N":
			proposedY -= 1
		case "S":
			proposedY += 1
		case "W":
			proposedX -= 1
		case "E":
			proposedX += 1
		}

		if _, ok := f.proposedMovement[proposedY]; !ok {
			f.proposedMovement[proposedY] = make(map[int][]*elf)
		}
		f.proposedMovement[proposedY][proposedX] = append(f.proposedMovement[proposedY][proposedX], currElf)
	}
}

func (f *field) move() bool {
	var isMoved bool

	for y, movementY := range f.proposedMovement {
		for x, movementX := range movementY {
			if len(movementX) != 1 {
				continue
			}

			currElf := movementX[0]

			oldY := currElf.y
			oldX := currElf.x

			currElf.y = y
			currElf.x = x

			if y == oldY && x == oldX {
				continue
			}

			f.tiles[y][x].elf = currElf
			f.tiles[oldY][oldX].elf = nil

			isMoved = true
		}
	}

	f.proposedMovement = make(map[int]map[int][]*elf)

	leastPriority := f.priority[0]
	f.priority = f.priority[1:]
	f.priority = append(f.priority, leastPriority)

	f.round++

	return isMoved
}

func (f *field) count() int {
	var count int

	minY := math.MaxInt
	maxY := math.MinInt
	minX := math.MaxInt
	maxX := math.MinInt
	for _, e := range f.elves {
		minY = int(math.Min(float64(minY), float64(e.y)))
		maxY = int(math.Max(float64(maxY), float64(e.y)))
		minX = int(math.Min(float64(minX), float64(e.x)))
		maxX = int(math.Max(float64(maxX), float64(e.x)))
	}

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if f.tiles[i][j].elf == nil {
				count++
			}
		}
	}

	return count
}

func (f *field) print() {
	for _, row := range f.tiles {
		for _, item := range row {
			if item.elf != nil {
				fmt.Print(item.elf.id)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
