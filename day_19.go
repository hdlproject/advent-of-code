package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type inventory struct {
	ore      int
	clay     int
	obsidian int
	geode    int

	oreRobot      int
	clayRobot     int
	obsidianRobot int
	geodeRobot    int
}

type blueprint struct {
	id int

	oreOreReq        int
	clayOreReq       int
	obsidianOreReq   int
	obsidianClayReq  int
	geodeOreReq      int
	geodeObsidianReq int

	inventory
}

func (b *blueprint) calculate() int {
	for i := 0; i < 10; i++ {
		if b.checkGeodeRobot() {
			fmt.Println("geode")
			b.gather()
			b.createGeodeRobot()
		} else if b.checkObsidianRobot() {
			fmt.Println("obsidian")
			b.gather()
			b.createObsidianRobot()
		} else if b.checkClayRobot() {
			fmt.Println("clay")
			b.gather()
			b.createClayRobot()
		} else if b.checkOreRobot() {
			fmt.Println("ore")
			b.gather()
			b.createOreRobot()
		} else {
			b.gather()
		}

		spew.Dump(b.inventory)
	}

	return b.geode
}

func (b *blueprint) gather() {
	b.ore += b.oreRobot
	b.clay += b.clayRobot
	b.obsidian += b.obsidianRobot
	b.geode += b.geodeRobot
}

func (b *blueprint) checkOreRobot() bool {
	return b.ore >= b.oreOreReq
}

func (b *blueprint) createOreRobot() {
	if b.checkOreRobot() {
		b.oreRobot++
		b.ore -= b.oreOreReq
	}
}

func (b *blueprint) checkClayRobot() bool {
	return b.ore >= b.clayOreReq
}

func (b *blueprint) createClayRobot() {
	if b.checkClayRobot() {
		b.clayRobot++
		b.ore -= b.clayOreReq
	}
}

func (b *blueprint) checkObsidianRobot() bool {
	return b.ore >= b.obsidianOreReq && b.clay >= b.obsidianClayReq
}

func (b *blueprint) createObsidianRobot() {
	if b.checkObsidianRobot() {
		b.obsidianRobot++
		b.ore -= b.obsidianOreReq
		b.clay -= b.obsidianClayReq
	}
}

func (b *blueprint) checkGeodeRobot() bool {
	return b.ore >= b.geodeOreReq && b.obsidian >= b.geodeObsidianReq
}

func (b *blueprint) createGeodeRobot() {
	if b.checkGeodeRobot() {
		b.geodeRobot++
		b.ore -= b.geodeOreReq
		b.obsidian -= b.geodeObsidianReq
	}
}

func day19SolutionA() {
	f, err := os.Open("day_19_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var blueprints []blueprint

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		values := strings.Split(value, " ")

		id, _ := strconv.ParseInt(strings.TrimSuffix(values[1], ":"), 10, 64)
		oreReq, _ := strconv.ParseInt(values[6], 10, 64)
		clayReq, _ := strconv.ParseInt(values[12], 10, 64)
		obsidianOreReq, _ := strconv.ParseInt(values[18], 10, 64)
		obsidianClayReq, _ := strconv.ParseInt(values[21], 10, 64)
		geodeOreReq, _ := strconv.ParseInt(values[27], 10, 64)
		geodeObsidianReq, _ := strconv.ParseInt(values[30], 10, 64)

		blueprints = append(blueprints, blueprint{
			id:               int(id),
			oreOreReq:        int(oreReq),
			clayOreReq:       int(clayReq),
			obsidianOreReq:   int(obsidianOreReq),
			obsidianClayReq:  int(obsidianClayReq),
			geodeOreReq:      int(geodeOreReq),
			geodeObsidianReq: int(geodeObsidianReq),
			inventory: inventory{
				oreRobot: 1,
			},
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(blueprints[0].calculate())

}

func day19SolutionB() {

}

//func main() {
//	day19SolutionA()
//}
