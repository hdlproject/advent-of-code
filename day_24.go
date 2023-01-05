package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day24SolutionA() {
	f, err := os.Open("day_24_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	valley := newValley()

	id := 1
	var y int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		if strings.Contains(value, "###") {
			continue
		}

		values := strings.Split(value, "")

		var tiles [][]*blizzard
		var x int
		for _, val := range values {
			if val == "#" {
				continue
			}

			var bs []*blizzard
			switch val {
			case "^", "v", ">", "<":
				b := &blizzard{
					direction: val,
					id:        id,
					y:         y,
					x:         x,
				}

				bs = []*blizzard{b}

				valley.blizzards = append(valley.blizzards, b)

				id++
			}

			x++
			tiles = append(tiles, bs)
		}

		y++
		valley.tiles = append(valley.tiles, tiles)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	valley.setFinish()

	//for i := 0; i < 17; i++ {
	//	valley.move()
	//	valley.movePerson()
	//}
	//valley.print()

	var isFinished bool
	round := 1
	for !isFinished {
		valley.move()
		isFinished = valley.movePerson()
		if isFinished {
			break
		}

		round++
	}
	valley.print()

	fmt.Println(round)
}

func day24SolutionB() {

}

//func main() {
//	day24SolutionA()
//
//	//593 too high
//	//338 too high
//}
