package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var directions [8][2]int = [8][2]int{
	{-1, -1}, // top left
	{0, -1},  // top
	{1, -1},  // top right
	{-1, 0},  // left
	{1, 0},   // right
	{-1, 1},  // bottom left
	{0, 1},   // bottom
	{1, 1},   // bottom right
}

type Octopus struct {
	energy_level int
	has_flashed  bool
}

func (o *Octopus) increment() {
	o.energy_level += 1
}

func (o *Octopus) flash() {
	if !o.has_flashed {
		o.has_flashed = true
	}
}

func (o *Octopus) remove_fashed() int {
	if o.has_flashed {
		o.energy_level = 0
		o.has_flashed = false
		return 1
	}
	return 0
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var energy_levels [][]Octopus
	for scanner.Scan() {
		var row []Octopus
		for _, r := range scanner.Text() {
			digit, _ := strconv.Atoi(string(r))
			row = append(row, Octopus{energy_level: digit, has_flashed: false})
		}
		energy_levels = append(energy_levels, row)
	}

	simulate_energy_levels(100, &energy_levels)
	simulate_until_all_flash(&energy_levels)
}

func simulate_energy_levels(steps int, energy_levels *[][]Octopus) {
	total_flashes := 0
	for step := 0; step < steps; step++ {
		increase_energy_level(energy_levels)
		flash(energy_levels)
		total_flashes += count_flashes(energy_levels)
	}
	fmt.Printf("Part 1: %d\n", total_flashes)
}

func simulate_until_all_flash(energy_levels *[][]Octopus) {
	i := 100
	for {
		i++
		increase_energy_level(energy_levels)
		flash(energy_levels)
		if count_flashes(energy_levels) == 100 {
			break
		}
	}
	fmt.Printf("Part 2: %d\n", i)
}

func increase_energy_level(energy_levels *[][]Octopus) {
	for y := 0; y < len(*energy_levels); y++ {
		for x := 0; x < len((*energy_levels)[0]); x++ {
			(*energy_levels)[y][x].increment()
		}
	}
}

func flash(energy_levels *[][]Octopus) {
	for y := 0; y < len(*energy_levels); y++ {
		for x := 0; x < len((*energy_levels)[0]); x++ {
			if (*energy_levels)[y][x].energy_level > 9 && !(*energy_levels)[y][x].has_flashed {
				(*energy_levels)[y][x].flash()
				simulate_flashes(x, y, energy_levels)
			}
		}
	}
}

func simulate_flashes(x int, y int, energy_levels *[][]Octopus) {
	max_vertical := len(*energy_levels)
	max_horizontal := len((*energy_levels)[0])
	for _, direction := range directions {
		dir_x := x + direction[0]
		dir_y := y + direction[1]
		if dir_x >= 0 && dir_x < max_horizontal && dir_y >= 0 && dir_y < max_vertical {
			(*energy_levels)[dir_y][dir_x].increment()
			if (*energy_levels)[dir_y][dir_x].energy_level > 9 && !(*energy_levels)[dir_y][dir_x].has_flashed {
				(*energy_levels)[dir_y][dir_x].flash()
				simulate_flashes(dir_x, dir_y, energy_levels)
			}
		}
	}
}

func count_flashes(energy_levels *[][]Octopus) int {
	flashes := 0
	for y := 0; y < len(*energy_levels); y++ {
		for x := 0; x < len((*energy_levels)[0]); x++ {
			flashes += (*energy_levels)[y][x].remove_fashed()
		}
	}
	return flashes
}
