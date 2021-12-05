package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type VentLine struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

func new_vent_lines(coords []int) *VentLine {
	return &VentLine{
		x1: coords[0],
		x2: coords[2],
		y1: coords[1],
		y2: coords[3],
	}
}

func (v *VentLine) from_to_coords() (int, int, int, int) {
	if v.x1 < v.x2 {
		return v.x1, v.y1, v.x2, v.y2
	} else {
		return v.x2, v.y2, v.x1, v.y1
	}
}

func main() {

	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var vent_lines []VentLine

	for scanner.Scan() {
		row := scanner.Text()
		vent_lines = append(vent_lines, parse_row(&row))
	}
	part_1(&vent_lines)
	part_2(&vent_lines)
}

func parse_row(row *string) VentLine {
	var coords []int
	split_row := strings.Split(strings.Join(strings.Split((*row), " "), ""), "->")
	for _, r := range split_row {
		two_coords := strings.Split(r, ",")
		result, _ := strconv.Atoi(two_coords[0])
		coords = append(coords, result)
		result, _ = strconv.Atoi(two_coords[1])
		coords = append(coords, result)
	}
	return *new_vent_lines(coords)
}

func from_to(num1 int, num2 int) (int, int) {
	if num1 < num2 {
		return num1, num2
	}
	return num2, num1
}

func part_1(vent_lines *[]VentLine) {
	var vents_map = map[int]map[int]int{}

	for _, vent_line := range *vent_lines {
		var from int
		var to int
		var is_x bool = false
		if vent_line.x1 == vent_line.x2 {
			is_x = true
			from, to = from_to(vent_line.y1, vent_line.y2)
		} else if vent_line.y1 == vent_line.y2 {
			from, to = from_to(vent_line.x1, vent_line.x2)
		} else {
			continue
		}
		for ; from <= to; from++ {
			if is_x {
				if vents_map[vent_line.x1] == nil {
					vents_map[vent_line.x1] = make(map[int]int)
				}
				vents_map[vent_line.x1][from] += 1
			} else {
				if vents_map[from] == nil {
					vents_map[from] = make(map[int]int)
				}
				vents_map[from][vent_line.y1] += 1
			}
		}
	}

	points := 0
	for _, val := range vents_map {
		for _, v := range val {
			if v >= 2 {
				points += 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", points)
}

func part_2(vent_lines *[]VentLine) {
	var vents_map = map[int]map[int]int{}

	for _, vent_line := range *vent_lines {
		var from int
		var to int
		var is_x bool = false
		var is_diagonal bool = false
		if vent_line.x1 == vent_line.x2 {
			is_x = true
			from, to = from_to(vent_line.y1, vent_line.y2)
		} else if vent_line.y1 == vent_line.y2 {
			from, to = from_to(vent_line.x1, vent_line.x2)
		} else {
			is_diagonal = true
		}
		if !is_diagonal {
			for ; from <= to; from++ {
				if is_x {
					if vents_map[vent_line.x1] == nil {
						vents_map[vent_line.x1] = make(map[int]int)
					}
					vents_map[vent_line.x1][from] += 1
				} else {
					if vents_map[from] == nil {
						vents_map[from] = make(map[int]int)
					}
					vents_map[from][vent_line.y1] += 1
				}
			}
		} else {
			x1, y1, x2, y2 := vent_line.from_to_coords()
			steps := x2 - x1
			for i := 0; i <= steps; i++ {
				if vents_map[x1] == nil {
					vents_map[x1] = make(map[int]int)
				}
				vents_map[x1][y1] += 1
				if y1 < y2 {
					y1++
				} else {
					y1--
				}
				x1++
			}
		}
	}

	points := 0
	for _, val := range vents_map {
		for _, v := range val {
			if v >= 2 {
				points += 1
			}
		}
	}

	fmt.Printf("Part 2: %d\n", points)
}
