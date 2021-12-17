package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TargetArea struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (t *TargetArea) is_coord_in(x int, y int) bool {
	if x >= t.x1 && x <= t.x2 && y >= t.y1 && y <= t.y2 {
		return true
	}
	return false
}

func (t *TargetArea) is_missed(x int, y int) bool {
	if x > t.x2 || y < t.y1 {
		return true
	}
	return false
}

var highest_y int = 0
var in_area_count int = 0

func main() {
	file, _ := os.Open("input.txt")
	
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	target_area_parsing := strings.Split(scanner.Text(), " ")
	x1, x2 := parse(target_area_parsing[2])
	y1, y2 := parse(target_area_parsing[3])
	target_area := TargetArea{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}
	probe_launch(target_area)
	fmt.Printf("Part 1: %d\n", highest_y)
	fmt.Printf("Part 2: %d\n", in_area_count)
}

func parse(coord string) (int, int) {
	coords := []int{}
	i := 0
	is_negative := false
	for _, char := range coord {
		if char == '-'  {
			is_negative = true
		}
		if char >= '0' && char <= '9' {
			if len(coords) == i {
				coords = append(coords, int(char - '0'))
			} else {
				coords[i] = coords[i] * 10 + int(char - '0')
			}
		} else if len(coords) != i {
			if is_negative {
				coords[i] = -coords[i]
			}
			is_negative = false
			i++
		}
	}
	if is_negative {
		coords[i] = -coords[i]
	}
	return coords[0], coords[1] 
}

func probe_launch(targetArea TargetArea) {
	for y := -3000; y < 3000; y++ {
		for x := 0; x < 3000; x++ {
			simulate_proble_launch(x, y, targetArea)
		}
	}
}

func simulate_proble_launch(x int, y int, targetArea TargetArea) {
	max_y := 0
	pos_x := 0
	pos_y := 0
	for {
		if pos_y > max_y {
			max_y = pos_y
		}
		pos_x += x
		pos_y += y
		if x > 0 {
			x -= 1
		} else if x < 0 {
			x += 1
		}
		y -= 1
		if targetArea.is_missed(pos_x, pos_y) {
			break
		}
		if targetArea.is_coord_in(pos_x, pos_y) {
			in_area_count += 1
			if highest_y < max_y {
				highest_y = max_y
			}
			break
		}
	}
}
