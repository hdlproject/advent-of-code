package main

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

type chamber struct {
	highestRock  int
	space        [][]bool
	coloredSpace [][]coloredSpace
	maxLength    int

	base       int
	coverMap   map[int]int
	futureBase int

	turn int

	gasPushes     []string
	lastPushIndex int
}

type coloredSpace struct {
	isOccupied bool
	color      *color.Color
}

func (c *chamber) execute() {
	for i := 0; i < 1000000; i++ {
		rock := newRock(c)

		var isStopped bool
		for !isStopped {
			switch c.gasPushes[(c.lastPushIndex % len(c.gasPushes))] {
			case "<":
				rock.moveLeft(c)
			case ">":
				rock.moveRight(c)
			}
			c.lastPushIndex++

			isStopped = rock.moveDown(c)
		}

		rock.occupySpace(c)

		//if i%2022 == 0 {
		//	c.reset()
		//}

		//c.resetType2()

		c.resetType3()

		c.turn++
	}
}

func (c *chamber) reset() {
	coverMap := make(map[int]struct{})
	var newBase int
	for i := c.highestRock; i >= 0; i-- {
		for j := 0; j < c.maxLength; j++ {
			if c.space[i][j] {
				coverMap[j] = struct{}{}
			}
		}

		if len(coverMap) == c.maxLength {
			newBase = i
			break
		}
	}

	c.highestRock -= newBase
	c.base += newBase
	c.space = c.space[newBase:]
	c.coloredSpace = c.coloredSpace[newBase:]
}

func (c *chamber) resetType2() {
	if len(c.coverMap) < c.maxLength {
		return
	}

	c.reset()
	c.coverMap = make(map[int]int)
}

func (c *chamber) resetType3() {
	if len(c.coverMap) < c.maxLength {
		return
	}

	newBase := math.MaxInt
	for _, cover := range c.coverMap {
		if cover < newBase {
			newBase = cover
		}
	}

	c.highestRock -= newBase
	c.base += newBase
	c.space = c.space[newBase:]
	c.coloredSpace = c.coloredSpace[newBase:]

	c.coverMap = make(map[int]int)
}

func (c *chamber) printSpace() {
	for i := len(c.space) - 1; i >= 0; i-- {
		for j := 0; j < c.maxLength; j++ {
			if c.space[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (c *chamber) printColoredSpace() {
	for i := len(c.coloredSpace) - 1; i >= 0; i-- {
		for j := 0; j < c.maxLength; j++ {
			if c.coloredSpace[i][j].isOccupied {
				c.coloredSpace[i][j].color.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type minus struct {
	left          int
	right         int
	currentHeight int
}

func (r *minus) moveRight(chamber *chamber) {
	if r.right+1 >= chamber.maxLength {
		return
	}

	if chamber.space[r.currentHeight][r.right+1] {
		return
	}

	r.left++
	r.right++
}

func (r *minus) moveLeft(chamber *chamber) {
	if r.left-1 < 0 {
		return
	}

	if chamber.space[r.currentHeight][r.left-1] {
		return
	}

	r.left--
	r.right--
}

func (r *minus) moveDown(chamber *chamber) bool {
	if r.currentHeight-1 < 0 {
		return true
	}

	for i := r.left; i <= r.right; i++ {
		if chamber.space[r.currentHeight-1][i] {
			return true
		}
	}

	r.currentHeight--
	return false
}

func (r *minus) occupySpace(chamber *chamber) {
	c := color.New(color.FgBlue).Add(color.Underline)

	for i := r.left; i <= r.right; i++ {
		chamber.space[r.currentHeight][i] = true
		chamber.coloredSpace[r.currentHeight][i] = coloredSpace{
			isOccupied: true,
			color:      c,
		}

	}

	chamber.highestRock = int(math.Max(float64(chamber.highestRock), float64(r.currentHeight)))

	for i := r.left; i <= r.right; i++ {
		chamber.coverMap[i] = int(math.Max(float64(chamber.coverMap[i]), float64(r.currentHeight)))
	}
}

type plus struct {
	left          int
	right         int
	currentHeight int
}

func (r *plus) moveRight(chamber *chamber) {
	if r.right+1 >= chamber.maxLength {
		return
	}

	if chamber.space[r.currentHeight+2][r.right] {
		return
	}

	if chamber.space[r.currentHeight+1][r.right+1] {
		return
	}

	if chamber.space[r.currentHeight][r.right] {
		return
	}

	r.left++
	r.right++
}

func (r *plus) moveLeft(chamber *chamber) {
	if r.left-1 < 0 {
		return
	}

	if chamber.space[r.currentHeight+2][r.left] {
		return
	}

	if chamber.space[r.currentHeight+1][r.left-1] {
		return
	}

	if chamber.space[r.currentHeight][r.left] {
		return
	}

	r.left--
	r.right--
}

func (r *plus) moveDown(chamber *chamber) bool {
	if r.currentHeight-1 < 0 {
		return true
	}

	if chamber.space[r.currentHeight][r.left] {
		return true
	}

	if chamber.space[r.currentHeight][r.right] {
		return true
	}

	if chamber.space[r.currentHeight-1][r.left+1] {
		return true
	}

	r.currentHeight--
	return false
}

func (r *plus) occupySpace(chamber *chamber) {
	c := color.New(color.FgRed).Add(color.Underline)

	chamber.space[r.currentHeight+2][r.left+1] = true
	chamber.coloredSpace[r.currentHeight+2][r.left+1] = coloredSpace{
		isOccupied: true,
		color:      c,
	}

	for i := r.left; i <= r.right; i++ {
		chamber.space[r.currentHeight+1][i] = true
		chamber.coloredSpace[r.currentHeight+1][i] = coloredSpace{
			isOccupied: true,
			color:      c,
		}
	}

	chamber.space[r.currentHeight][r.left+1] = true
	chamber.coloredSpace[r.currentHeight][r.left+1] = coloredSpace{
		isOccupied: true,
		color:      c,
	}

	chamber.highestRock = int(math.Max(float64(chamber.highestRock), float64(r.currentHeight+2)))

	chamber.coverMap[r.left] = int(math.Max(float64(chamber.coverMap[r.left]), float64(r.currentHeight+1)))
	chamber.coverMap[r.left+1] = int(math.Max(float64(chamber.coverMap[r.left+1]), float64(r.currentHeight+2)))
	chamber.coverMap[r.right] = int(math.Max(float64(chamber.coverMap[r.left+1]), float64(r.currentHeight+1)))
}

type elbow struct {
	left          int
	right         int
	currentHeight int
}

func (r *elbow) moveRight(chamber *chamber) {
	if r.right+1 >= chamber.maxLength {
		return
	}

	for i := 0; i <= 2; i++ {
		if chamber.space[r.currentHeight+i][r.right+1] {
			return
		}
	}

	r.left++
	r.right++
}

func (r *elbow) moveLeft(chamber *chamber) {
	if r.left-1 < 0 {
		return
	}

	if chamber.space[r.currentHeight][r.left-1] {
		return
	}

	for i := 1; i <= 2; i++ {
		if chamber.space[r.currentHeight+i][r.left+1] {
			return
		}
	}

	r.left--
	r.right--
}

func (r *elbow) moveDown(chamber *chamber) bool {
	if r.currentHeight-1 < 0 {
		return true
	}

	for i := r.left; i <= r.right; i++ {
		if chamber.space[r.currentHeight-1][i] {
			return true
		}
	}

	r.currentHeight--
	return false
}

func (r *elbow) occupySpace(chamber *chamber) {
	c := color.New(color.FgGreen).Add(color.Underline)

	for i := r.left; i <= r.right; i++ {
		chamber.space[r.currentHeight][i] = true
		chamber.coloredSpace[r.currentHeight][i] = coloredSpace{
			isOccupied: true,
			color:      c,
		}
	}

	chamber.space[r.currentHeight+1][r.right] = true
	chamber.coloredSpace[r.currentHeight+1][r.right] = coloredSpace{
		isOccupied: true,
		color:      c,
	}

	chamber.space[r.currentHeight+2][r.right] = true
	chamber.coloredSpace[r.currentHeight+2][r.right] = coloredSpace{
		isOccupied: true,
		color:      c,
	}

	chamber.highestRock = int(math.Max(float64(chamber.highestRock), float64(r.currentHeight+2)))

	chamber.coverMap[r.left] = int(math.Max(float64(chamber.coverMap[r.left]), float64(r.currentHeight)))
	chamber.coverMap[r.left+1] = int(math.Max(float64(chamber.coverMap[r.left]), float64(r.currentHeight)))
	chamber.coverMap[r.right] = int(math.Max(float64(chamber.coverMap[r.left]), float64(r.currentHeight+2)))
}

type pile struct {
	pos           int
	currentHeight int
}

func (r *pile) moveRight(chamber *chamber) {
	if r.pos+1 >= chamber.maxLength {
		return
	}

	for i := 0; i <= 3; i++ {
		if chamber.space[r.currentHeight+i][r.pos+1] {
			return
		}
	}

	r.pos++
}

func (r *pile) moveLeft(chamber *chamber) {
	if r.pos-1 < 0 {
		return
	}

	for i := 0; i <= 3; i++ {
		if chamber.space[r.currentHeight+i][r.pos-1] {
			return
		}
	}

	r.pos--
}

func (r *pile) moveDown(chamber *chamber) bool {
	if r.currentHeight-1 < 0 {
		return true
	}

	if chamber.space[r.currentHeight-1][r.pos] {
		return true
	}

	r.currentHeight--
	return false
}

func (r *pile) occupySpace(chamber *chamber) {
	c := color.New(color.FgMagenta).Add(color.Underline)

	for i := 0; i <= 3; i++ {
		chamber.space[r.currentHeight+i][r.pos] = true
		chamber.coloredSpace[r.currentHeight+i][r.pos] = coloredSpace{
			isOccupied: true,
			color:      c,
		}
	}

	chamber.highestRock = int(math.Max(float64(chamber.highestRock), float64(r.currentHeight+3)))

	chamber.coverMap[r.pos] = int(math.Max(float64(chamber.coverMap[r.pos]), float64(r.currentHeight+3)))
}

type square struct {
	left          int
	right         int
	currentHeight int
}

func (r *square) moveRight(chamber *chamber) {
	if r.right+1 >= chamber.maxLength {
		return
	}

	for i := 0; i <= 1; i++ {
		if chamber.space[r.currentHeight+i][r.right+1] {
			return
		}
	}

	r.right++
	r.left++
}

func (r *square) moveLeft(chamber *chamber) {
	if r.left-1 < 0 {
		return
	}

	for i := 0; i <= 1; i++ {
		if chamber.space[r.currentHeight+i][r.left-1] {
			return
		}
	}

	r.right--
	r.left--
}

func (r *square) moveDown(chamber *chamber) bool {
	if r.currentHeight-1 < 0 {
		return true
	}

	if chamber.space[r.currentHeight-1][r.left] {
		return true
	}

	if chamber.space[r.currentHeight-1][r.right] {
		return true
	}

	r.currentHeight--
	return false
}

func (r *square) occupySpace(chamber *chamber) {
	c := color.New(color.FgHiYellow).Add(color.Underline)

	for i := r.left; i <= r.right; i++ {
		chamber.space[r.currentHeight][i] = true
		chamber.coloredSpace[r.currentHeight][i] = coloredSpace{
			isOccupied: true,
			color:      c,
		}
	}

	for i := r.left; i <= r.right; i++ {
		chamber.space[r.currentHeight+1][i] = true
		chamber.coloredSpace[r.currentHeight+1][i] = coloredSpace{
			isOccupied: true,
			color:      c,
		}
	}

	chamber.highestRock = int(math.Max(float64(chamber.highestRock), float64(r.currentHeight+1)))

	chamber.coverMap[r.left] = int(math.Max(float64(chamber.coverMap[r.left]), float64(r.currentHeight+1)))
	chamber.coverMap[r.right] = int(math.Max(float64(chamber.coverMap[r.right]), float64(r.currentHeight+1)))
}

type rockInf interface {
	moveRight(*chamber)
	moveLeft(*chamber)
	moveDown(*chamber) bool
	occupySpace(*chamber)
}

func newRock(chamber *chamber) rockInf {
	switch chamber.turn % 5 {
	case 0:
		highest := chamber.highestRock
		newHeight := int(math.Max(float64(len(chamber.space)-1), float64(highest+4)))
		for i := len(chamber.space); i < newHeight+1; i++ {
			chamber.space = append(chamber.space, make([]bool, 7))
			chamber.coloredSpace = append(chamber.coloredSpace, make([]coloredSpace, 7))
		}

		return &minus{
			left:          2,
			right:         5,
			currentHeight: chamber.highestRock + 4,
		}
	case 1:
		highest := chamber.highestRock
		newHeight := int(math.Max(float64(len(chamber.space)-1), float64(highest+6)))
		for i := len(chamber.space); i < newHeight+1; i++ {
			chamber.space = append(chamber.space, make([]bool, 7))
			chamber.coloredSpace = append(chamber.coloredSpace, make([]coloredSpace, 7))
		}

		return &plus{
			left:          2,
			right:         4,
			currentHeight: chamber.highestRock + 4,
		}
	case 2:
		highest := chamber.highestRock
		newHeight := int(math.Max(float64(len(chamber.space)-1), float64(highest+6)))
		for i := len(chamber.space); i < newHeight+1; i++ {
			chamber.space = append(chamber.space, make([]bool, 7))
			chamber.coloredSpace = append(chamber.coloredSpace, make([]coloredSpace, 7))
		}

		return &elbow{
			left:          2,
			right:         4,
			currentHeight: chamber.highestRock + 4,
		}
	case 3:
		highest := chamber.highestRock
		newHeight := int(math.Max(float64(len(chamber.space)-1), float64(highest+7)))
		for i := len(chamber.space); i < newHeight+1; i++ {
			chamber.space = append(chamber.space, make([]bool, 7))
			chamber.coloredSpace = append(chamber.coloredSpace, make([]coloredSpace, 7))
		}

		return &pile{
			pos:           2,
			currentHeight: chamber.highestRock + 4,
		}
	case 4:
		highest := chamber.highestRock
		newHeight := int(math.Max(float64(len(chamber.space)-1), float64(highest+5)))
		for i := len(chamber.space); i < newHeight+1; i++ {
			chamber.space = append(chamber.space, make([]bool, 7))
			chamber.coloredSpace = append(chamber.coloredSpace, make([]coloredSpace, 7))
		}

		return &square{
			left:          2,
			right:         3,
			currentHeight: chamber.highestRock + 4,
		}
	}

	return nil
}
