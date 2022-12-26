package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	vertOffset = 3
	leftOffset
	width = 7
)

func main() {
	fmt.Println(part1("example_input.txt"))
	//fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	//fmt.Println(part2("input.txt"))
}

type Shape struct {
	//uLim, lLim, rLim, dLim int
	Coords []utils.Coordinate
}

type grid struct {
	arr [][]int
}

func (s Shape) move(x utils.Coordinate) Shape {
	var res Shape
	for _, c := range s.Coords {
		res.Coords = append(res.Coords, c.Add(x))
	}
	return res
}

func (s Shape) leftOne() Shape {
	return s.move(utils.Coordinate{-1, 0})
}

func (s Shape) rightOne() Shape {
	return s.move(utils.Coordinate{1, 0})
}

func (s Shape) downOne() Shape {
	return s.move(utils.Coordinate{0, -1})
}

func (s Shape) yMax() int {
	max := math.MinInt64
	for _, c := range s.Coords {
		if c.Y > max {
			max = c.Y
		}
	}
	return max
}
func (s Shape) contains(c utils.Coordinate) bool {
	for _, coord := range s.Coords {
		if coord == c {
			return true
		}
	}
	return false
}

func (s Shape) hitWall() bool {
	for i := 0; i < len(s.Coords); i++ {
		if s.Coords[i].X < 0 || s.Coords[i].X >= 7 {
			return true
		}
	}
	return false
}

func attemptLRMove(move rune, s Shape, grid [][]bool) Shape {
	var res Shape
	switch move {
	case '<':
		fmt.Println("moving left")
		res = s.leftOne()
	case '>':
		fmt.Println("moving right")
		res = s.rightOne()
	}
	isHitWall := res.hitWall()
	if isHitWall {
		return s
	}
	isGridCollision := gridCollision(res, grid)
	if isGridCollision {
		fmt.Println("grid collision")
		return s
	}
	return res
}

func attemptDownMove(s Shape, grid [][]bool) (isHitBottom bool, s2 Shape) {
	res := s.downOne()
	if isGridCollision := gridCollision(res, grid); isGridCollision {
		return true, s
	}
	return false, res
}

func addToGrid(s Shape, grid [][]bool) [][]bool {
	for _, c := range s.Coords {
		grid[c.Y][c.X] = true
	}
	return grid
}

func gridCollision(s Shape, grid [][]bool) bool {
	for _, c := range s.Coords {
		if c.Y < 0 || grid[c.Y][c.X] {
			return true
		}
	}
	return false
}

func printGrid(s Shape, grid [][]bool) {

	for y := len(grid) - 1; y >= 0; y-- {
		line := "|"
		for x := 0; x < len(grid[0]); x++ {
			if s.contains(utils.Coordinate{x, y}) {
				line += "@"
			} else if grid[y][x] == true {
				line += "#"
			} else {
				line += "."
			}
		}
		line += "|"
		fmt.Println(line)
	}
	fmt.Println("+-------+")
	// print floor
}

func loadInput(filename string) (cmds []rune) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range line {
			cmds = append(cmds, r)
		}
		return cmds
	}
	utils.CheckErr(scanner.Err())
	return []rune{}
}

func part1(filename string) int {
	cmds := loadInput(filename)
	fmt.Printf("%c\n", cmds)

	horizontal := Shape{
		Coords: []utils.Coordinate{{2, 0}, {3, 0}, {4, 0}, {5, 0}},
	}
	plus := Shape{
		Coords: []utils.Coordinate{{3, 2}, {2, 1}, {3, 1}, {4, 1}, {3, 0}},
	}
	backwardL := Shape{
		Coords: []utils.Coordinate{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}},
	}
	allShapes := []Shape{horizontal, plus, backwardL}
	nShapes := len(allShapes)
	nRocks := 3
	ymax := -1
	idxCmd := 0
	m, n := 10, 7
	grid := make([][]bool, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]bool, n)
	}
	// while still got pieces to go forth:
	for i := 0; i < nRocks; i++ {
		// create the right new shape
		shapeToPlace := allShapes[i%nShapes]

		// put it in the correct v offset compared with the grid (wrt to the yMax of the grid)
		shapeToPlace = shapeToPlace.move(utils.Coordinate{X: 0, Y: vertOffset + ymax + 1})
		fmt.Println("\n\nPLACING NEW SHAPE")
		for {
			fmt.Println("start of loop")
			fmt.Println(shapeToPlace)
			printGrid(shapeToPlace, grid)
			// move it across if possible
			shapeToPlace = attemptLRMove(cmds[idxCmd], shapeToPlace, grid)
			idxCmd++ // never reuse
			// move down if possible, and repeat else break
			var isHitBottom bool
			isHitBottom, shapeToPlace = attemptDownMove(shapeToPlace, grid)

			// place shape into grid, and update yMax
			if isHitBottom {
				fmt.Println("placing: ", shapeToPlace)
				grid = addToGrid(shapeToPlace, grid)
				fmt.Println("\n\nplaced shape")
				printGrid(shapeToPlace, grid)
				if shapeToPlace.yMax() > ymax {
					ymax = shapeToPlace.yMax()
					fmt.Println("updated ymax to ", ymax)
				}
				break
			}
		}
	}

	return 0
}

func part2(filename string) int {
	return 0
}
