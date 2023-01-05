package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day23SolutionA() {
	f, err := os.Open("day_23_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	field := newField()

	var y int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		values := strings.Split(value, "")

		var currRow []fieldTile
		var x int
		for _, val := range values {
			switch val {
			case ".":
				currRow = append(currRow, fieldTile{})
			default:
				elf := &elf{
					y:  y,
					x:  x,
					id: val,
				}

				currRow = append(currRow, fieldTile{
					elf: elf,
				})

				field.elves = append(field.elves, elf)
			}

			x++
		}
		field.tiles = append(field.tiles, currRow)
		y++
	}

	field.colLen = len(field.tiles)
	field.rowLen = len(field.tiles[0])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		field.planMove()
		field.move()
	}
	//field.print()

	fmt.Println(field.count())
}

func day23SolutionB() {
	f, err := os.Open("day_23_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	field := newField()

	var y int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		values := strings.Split(value, "")

		var currRow []fieldTile
		var x int
		for _, val := range values {
			switch val {
			case ".":
				currRow = append(currRow, fieldTile{})
			default:
				elf := &elf{
					y:  y,
					x:  x,
					id: val,
				}

				currRow = append(currRow, fieldTile{
					elf: elf,
				})

				field.elves = append(field.elves, elf)
			}

			x++
		}
		field.tiles = append(field.tiles, currRow)
		y++
	}

	field.colLen = len(field.tiles)
	field.rowLen = len(field.tiles[0])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	isMoved := true
	for isMoved {
		field.planMove()
		isMoved = field.move()
	}
	//field.print()

	fmt.Println(field.round - 1)
}

//func main() {
//	day23SolutionB()
//}
