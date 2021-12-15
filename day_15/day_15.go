package main

import (
	"bufio"
	"fmt"
	"os"
)

type Poisition struct {
	x int
	y int
}

type RiskPoint struct {
	risk_level int
	min_risk_cumulative int
	best_previous_neighbour Poisition
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	risk_map := [][]RiskPoint{}
	for scanner.Scan() {
		var row []RiskPoint
		var v_row []bool
		for _, str := range scanner.Text() {
			row = append(row, RiskPoint{ risk_level: int(str) - 48, min_risk_cumulative: -1, best_previous_neighbour: Poisition{ x: -1, y: -1}} )
			v_row = append(v_row, false)
		}
		risk_map = append(risk_map, row)
	}

	
	fmt.Printf("Part 1: %d\n", find_path(risk_map)) 
	risk_map = make_map_p2(&risk_map)
	fmt.Printf("Part 2: %d\n", find_path(risk_map)) 
}

func find_path(risk_map [][]RiskPoint) int {
	risk_map[0][0].min_risk_cumulative = 0
	i := 0
	for {
		// fmt.Println(i)
		for y := 0; y < len(risk_map); y++ {
			for x := 0; x < len(risk_map[0]); x++ {
				current := &risk_map[y][x]
				p := Poisition{x: x, y: y}
				// left
				if x - 1 >= 0 {
					calculate_risk(current, &risk_map[y][x - 1], p)			
				}
				// right
				if x + 1 < len(risk_map[0]) {
					calculate_risk(current, &risk_map[y][x + 1], p)			
				}
				// above
				if y - 1 >= 0 {
					calculate_risk(current, &risk_map[y - 1][x], p)			
				}
				// bellow
				if y + 1 < len(risk_map) {
					calculate_risk(current, &risk_map[y + 1][x], p)			
				}
			}
		}
		if i == 10 {
			break
		}
		i++
	}

	return risk_map[len(risk_map) - 1][len(risk_map) - 1].min_risk_cumulative
}

func calculate_risk(initial *RiskPoint, other *RiskPoint, pos Poisition) {
	new_min_risk := initial.min_risk_cumulative + other.risk_level 
	if other.min_risk_cumulative != -1 {
		if other.min_risk_cumulative > new_min_risk {
			other.min_risk_cumulative = new_min_risk
			other.best_previous_neighbour = pos
		}
	} else {
		other.min_risk_cumulative = new_min_risk
		other.best_previous_neighbour = pos
	}
}

func make_map_p2(risk_map *[][]RiskPoint) [][]RiskPoint {
	var risk_map_p2 [][]RiskPoint
	// copy risk_map
	for _, row := range *risk_map {
		risk_map_p2 = append(risk_map_p2, row)
	}

	max_x_len := len(risk_map_p2[0])
	max_y_len := len(risk_map_p2)

	for y := 0; y < max_y_len; y++ {
		for i := 0; i < 4; i++ {
			for x := 0; x < max_x_len; x++ {
				point := RiskPoint{
					risk_level: risk_map_p2[y][x + max_x_len * i].risk_level,
					min_risk_cumulative: -1,
					best_previous_neighbour: Poisition{ x: -1, y: -1},
				}
				point.risk_level += 1
				if point.risk_level > 9 {
					point.risk_level = 1
				}
				risk_map_p2[y] = append(risk_map_p2[y], point)
			}
		}
	}

	max_x_len = len(risk_map_p2[0])

	for i := 0; i < 4; i++ {
		for y := 0; y < max_y_len; y++ {
			var new_row []RiskPoint
			for x := 0; x < max_x_len; x++ {
				point := RiskPoint{
					risk_level: risk_map_p2[y + max_y_len * i][x].risk_level,
					min_risk_cumulative: -1,
					best_previous_neighbour: Poisition{ x: -1, y: -1},
				}

				point.risk_level += 1
				if point.risk_level > 9 {
					point.risk_level = 1
				}

				new_row = append(new_row, point)
			}
			risk_map_p2 = append(risk_map_p2, new_row)
		}
	}

	return risk_map_p2
}
