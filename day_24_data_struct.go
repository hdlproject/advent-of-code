package main

import (
	"fmt"
)

type valley struct {
	finishY int
	finishX int

	tiles     [][][]*blizzard
	nextTiles [][][]*blizzard

	blizzards []*blizzard

	currY int
	currX int
	round int
}

func newValley() *valley {
	return &valley{
		currY: -1,
		currX: 0,
	}
}

func (v *valley) setFinish() {
	v.finishY = len(v.tiles)
	v.finishX = len(v.tiles[0]) - 1
}

type blizzard struct {
	id        int
	direction string
	y         int
	x         int
}

func (v *valley) move() {
	for _, b := range v.blizzards {
		switch b.direction {
		case "^":
			v.moveUp(b)
		case "v":
			v.moveDown(b)
		case ">":
			v.moveRight(b)
		case "<":
			v.moveLeft(b)
		}
	}
}

func (v *valley) moveUp(b *blizzard) {
	nextY := b.y - 1
	if nextY < 0 {
		nextY = len(v.tiles) - 1
	}

	v.moveVertical(b, nextY)
}

func (v *valley) moveDown(b *blizzard) {
	nextY := b.y + 1
	if nextY >= len(v.tiles) {
		nextY = 0
	}

	v.moveVertical(b, nextY)
}

func (v *valley) moveVertical(b *blizzard, nextY int) {
	var blizzards []*blizzard
	for _, currB := range v.tiles[b.y][b.x] {
		if currB.id == b.id {
			continue
		}

		blizzards = append(blizzards, currB)
	}

	v.tiles[nextY][b.x] = append(v.tiles[nextY][b.x], b)
	v.tiles[b.y][b.x] = blizzards

	b.y = nextY
}

func (v *valley) moveRight(b *blizzard) {
	nextX := b.x + 1
	if nextX >= len(v.tiles[0]) {
		nextX = 0
	}

	v.moveHorizontal(b, nextX)
}

func (v *valley) moveLeft(b *blizzard) {
	nextX := b.x - 1
	if nextX < 0 {
		nextX = len(v.tiles[0]) - 1
	}

	v.moveHorizontal(b, nextX)
}

func (v *valley) moveHorizontal(b *blizzard, nextX int) {
	var blizzards []*blizzard
	for _, currB := range v.tiles[b.y][b.x] {
		if currB.id == b.id {
			continue
		}

		blizzards = append(blizzards, currB)
	}

	v.tiles[b.y][nextX] = append(v.tiles[b.y][nextX], b)
	v.tiles[b.y][b.x] = blizzards

	b.x = nextX
}

func (v *valley) print() {
	for y, row := range v.tiles {
		for x, item := range row {
			if y == v.currY && x == v.currX {
				fmt.Print("E")
				continue
			}

			switch len(item) {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print(item[0].direction)
			default:
				fmt.Print(len(item))
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println(v.currY, v.currX)
}

func (v *valley) movePerson() bool {
	nextY := v.currY
	nextX := v.currX

	if v.currY == -1 {
		if len(v.tiles[0][0]) == 0 {
			nextY = 0
			nextX = 0
		}
	} else {
		if v.currY+1 == len(v.tiles) && v.currX == len(v.tiles[0])-1 {
			return true
		}

		if v.currX+1 < len(v.tiles[0]) && len(v.tiles[v.currY][v.currX+1]) == 0 {
			nextY = v.currY
			nextX = v.currX + 1
		} else if v.currY+1 < len(v.tiles) && len(v.tiles[v.currY+1][v.currX]) == 0 {
			nextY = v.currY + 1
			nextX = v.currX
		}

		//else if v.currY-1 >= 0 && len(v.tiles[v.currY-1][v.currX]) == 0 {
		//	nextY = v.currY - 1
		//	nextX = v.currX
		//} else if v.currX-1 >= 0 && len(v.tiles[v.currY][v.currX-1]) == 0 {
		//	nextY = v.currY
		//	nextX = v.currX - 1
		//}
	}

	if nextY == v.currY && nextX == v.currX {
		return false
	}

	v.currY = nextY
	v.currX = nextX

	return false
}
