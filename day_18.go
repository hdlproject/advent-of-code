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

type cube struct {
	top       bool
	bottom    bool
	right     bool
	left      bool
	front     bool
	back      bool
	isStone   bool
	isVisited bool
}

func (c cube) getSurface() int {
	var count int
	if !c.top {
		count++
	}
	if !c.bottom {
		count++
	}
	if !c.right {
		count++
	}
	if !c.left {
		count++
	}
	if !c.front {
		count++
	}
	if !c.back {
		count++
	}
	return count
}

func day18SolutionA() {
	f, err := os.Open("day_18_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cubes := make(map[int]map[int]map[int]cube)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		values := strings.Split(value, ",")
		x, _ := strconv.ParseInt(values[0], 10, 64)
		y, _ := strconv.ParseInt(values[1], 10, 64)
		z, _ := strconv.ParseInt(values[2], 10, 64)

		buildSurface(cubes, int(y), int(x), int(z))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var total int
	for _, row := range cubes {
		for _, depth := range row {
			for _, item := range depth {
				total += item.getSurface()
			}
		}
	}

	fmt.Println(total)
}

func buildSurface(cubes map[int]map[int]map[int]cube, y, x, z int) {
	if _, ok := cubes[y]; !ok {
		cubes[y] = make(map[int]map[int]cube)
	}
	if _, ok := cubes[y][x]; !ok {
		cubes[y][x] = make(map[int]cube)
	}
	if _, ok := cubes[y][x][z]; !ok {
		currCube := cube{
			top:    isExist(cubes, y-1, x, z),
			bottom: isExist(cubes, y+1, x, z),
			right:  isExist(cubes, y, x+1, z),
			left:   isExist(cubes, y, x-1, z),
			front:  isExist(cubes, y, x, z-1),
			back:   isExist(cubes, y, x, z+1),
		}

		cubes[y][x][z] = currCube

		if currCube.top {
			adjCube := cubes[y-1][x][z]
			adjCube.bottom = true
			cubes[y-1][x][z] = adjCube
		}
		if currCube.bottom {
			adjCube := cubes[y+1][x][z]
			adjCube.top = true
			cubes[y+1][x][z] = adjCube
		}
		if currCube.right {
			adjCube := cubes[y][x+1][z]
			adjCube.left = true
			cubes[y][x+1][z] = adjCube
		}
		if currCube.left {
			adjCube := cubes[y][x-1][z]
			adjCube.right = true
			cubes[y][x-1][z] = adjCube
		}
		if currCube.front {
			adjCube := cubes[y][x][z-1]
			adjCube.back = true
			cubes[y][x][z-1] = adjCube
		}
		if currCube.back {
			adjCube := cubes[y][x][z+1]
			adjCube.front = true
			cubes[y][x][z+1] = adjCube
		}
	}
}

func isExist(cubes map[int]map[int]map[int]cube, y, x, z int) bool {
	if _, ok := cubes[y]; !ok {
		return false
	}

	if _, ok := cubes[y][x]; !ok {
		return false
	}

	if _, ok := cubes[y][x][z]; !ok {
		return false
	}

	return true
}

func day18SolutionB() {
	f, err := os.Open("day_18_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cubes := make(map[int]map[int]map[int]cube)
	minY, minX, minZ := math.MaxInt, math.MaxInt, math.MaxInt
	maxY, maxX, maxZ := math.MinInt, math.MinInt, math.MinInt

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()

		values := strings.Split(value, ",")
		x, _ := strconv.ParseInt(values[0], 10, 64)
		y, _ := strconv.ParseInt(values[1], 10, 64)
		z, _ := strconv.ParseInt(values[2], 10, 64)

		minY = int(math.Min(float64(minY), float64(y)))
		minX = int(math.Min(float64(minX), float64(x)))
		minZ = int(math.Min(float64(minZ), float64(z)))

		maxY = int(math.Max(float64(maxY), float64(y)))
		maxX = int(math.Max(float64(maxX), float64(x)))
		maxZ = int(math.Max(float64(maxZ), float64(z)))

		buildSurface(cubes, int(y), int(x), int(z))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cubeArr := rebuildSurface(cubes, minY, maxY, minX, maxX, minZ, maxZ)

	cubeArr = traverseCube(cubeArr, 0, 0, 0)

	for y, row := range cubeArr {
		for x, depth := range row {
			for z, item := range depth {
				if !item.isVisited && !item.isStone {
					cubeArr[y-1][x][z].bottom = true
					cubeArr[y+1][x][z].top = true
					cubeArr[y][x+1][z].left = true
					cubeArr[y][x-1][z].right = true
					cubeArr[y][x][z-1].back = true
					cubeArr[y][x][z+1].front = true
				}
			}
		}
	}

	var total int
	for _, row := range cubeArr {
		for _, depth := range row {
			for _, item := range depth {
				if item.isStone {
					total += item.getSurface()
				}
			}
		}
	}

	fmt.Println(total)
}

func rebuildSurface(cubes map[int]map[int]map[int]cube, minY, maxY, minX, maxX, minZ, maxZ int) [][][]cube {
	var cubeArr [][][]cube
	for i := 0; i <= maxY+1; i++ {

		var cubeRow [][]cube
		for j := 0; j <= maxX+1; j++ {

			var cubeDepth []cube
			for k := 0; k <= maxZ+1; k++ {
				if _, ok := cubes[i][j][k]; ok {
					currCube := cubes[i][j][k]
					currCube.isStone = true
					cubeDepth = append(cubeDepth, currCube)
				} else {
					cubeDepth = append(cubeDepth, cube{})
				}
			}

			cubeRow = append(cubeRow, cubeDepth)
		}

		cubeArr = append(cubeArr, cubeRow)
	}

	return cubeArr
}

func traverseCube(cubes [][][]cube, y, x, z int) [][][]cube {
	cubes[y][x][z].isVisited = true

	if (y - 1) >= 0 {
		if !cubes[y-1][x][z].isStone && !cubes[y-1][x][z].isVisited {
			cubes = traverseCube(cubes, y-1, x, z)
		}
	}

	if (y+1) < len(cubes) && !cubes[y+1][x][z].isStone && !cubes[y+1][x][z].isVisited {
		cubes = traverseCube(cubes, y+1, x, z)
	}

	if (x+1) < len(cubes[0]) && !cubes[y][x+1][z].isStone && !cubes[y][x+1][z].isVisited {
		cubes = traverseCube(cubes, y, x+1, z)
	}

	if (x-1) >= 0 && !cubes[y][x-1][z].isStone && !cubes[y][x-1][z].isVisited {
		cubes = traverseCube(cubes, y, x-1, z)
	}

	if (z-1) >= 0 && !cubes[y][x][z-1].isStone && !cubes[y][x][z-1].isVisited {
		cubes = traverseCube(cubes, y, x, z-1)
	}

	if (z+1) < len(cubes[0][0]) && !cubes[y][x][z+1].isStone && !cubes[y][x][z+1].isVisited {
		cubes = traverseCube(cubes, y, x, z+1)
	}

	return cubes
}

//func main() {
//	day18SolutionB()
//}
