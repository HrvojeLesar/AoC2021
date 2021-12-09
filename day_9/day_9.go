package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pair struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var heightmap [][]int
	for scanner.Scan() {
		var row []int
		for _, digit := range scanner.Text() {
			row = append(row, int(digit)-48)
		}
		heightmap = append(heightmap, row)
	}
	low_points := part_1(&heightmap)
	var basin_sizes []int
	for _, point := range low_points {
		sum := 0
		part_2(point.x, point.y, &heightmap, &sum)
		basin_sizes = append(basin_sizes, sum)
	}

	sort.Ints(basin_sizes)
	ind := len(basin_sizes) - 1
	fmt.Printf("Part 2: %d\n", basin_sizes[ind]*basin_sizes[ind-1]*basin_sizes[ind-2])
}

func part_1(heightmap *[][]int) []Pair {
	var low_points []Pair
	horizontal_length := len(*heightmap)
	vertical_length := len((*heightmap)[0])
	sum := 0
	for x := 0; x < horizontal_length; x++ {
		for y := 0; y < vertical_length; y++ {
			if y != 0 {
				// check left
				if (*heightmap)[x][y] >= (*heightmap)[x][y-1] {
					continue
				}
			}

			if y != vertical_length-1 {
				// check right
				if (*heightmap)[x][y] >= (*heightmap)[x][y+1] {
					continue
				}
			}

			if x != 0 {
				// check above
				if (*heightmap)[x][y] >= (*heightmap)[x-1][y] {
					continue
				}
			}

			if x != horizontal_length-1 {
				// check bellow
				if (*heightmap)[x][y] >= (*heightmap)[x+1][y] {
					continue
				}
			}
			low_points = append(low_points, Pair{x: x, y: y})
			sum += (*heightmap)[x][y] + 1
		}
	}
	fmt.Printf("Part 1: %d\n", sum)
	return low_points
}

func part_2(x int, y int, heightmap *[][]int, basin_size *int) {
	horizontal_length := len(*heightmap)
	vertical_length := len((*heightmap)[0])

	if (*heightmap)[x][y] != 9 && (*heightmap)[x][y] != -1 {
		(*heightmap)[x][y] = -1
		*basin_size += 1
	}
	// go left
	if y-1 >= 0 && (*heightmap)[x][y-1] != -1 && (*heightmap)[x][y-1] != 9 {
		part_2(x, y-1, heightmap, basin_size)
	}
	// go right
	if y+1 < vertical_length && (*heightmap)[x][y+1] != -1 && (*heightmap)[x][y+1] != 9 {
		part_2(x, y+1, heightmap, basin_size)
	}
	// go up
	if x-1 >= 0 && (*heightmap)[x-1][y] != -1 && (*heightmap)[x-1][y] != 9 {
		part_2(x-1, y, heightmap, basin_size)
	}
	// go down
	if x+1 < horizontal_length && (*heightmap)[x+1][y] != -1 && (*heightmap)[x+1][y] != 9 {
		part_2(x+1, y, heightmap, basin_size)
	}
}
