package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	x int
	y int
}

func (p *Pair) to_string() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dots := make(map[string]*Pair)
	var folds []int
	is_parsing_folds := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			is_parsing_folds = true
			continue
		}
		if !is_parsing_folds {
			string_dots := strings.Split(scanner.Text(), ",")
			x, _ := strconv.Atoi(string_dots[0])
			y, _ := strconv.Atoi(string_dots[1])
			dots[scanner.Text()] = &Pair{x: x, y: y}
		} else {
			fold := strings.Split(scanner.Text(), "=")
			amount, _ := strconv.Atoi(fold[1])
			// y < 0
			if fold[0][len(fold[0]) - 1] == 'x' {
				folds = append(folds, amount)
			} else {
				folds = append(folds, -amount)
			}
		}
	}
	for i := range folds {
		fold(folds[i], &dots)
		if i == 0 {
			fmt.Printf("Part 1: %d\n", len(dots))
		}
	}

	part_2(&dots)
}

func fold(fold_value int, dots *map[string]*Pair) {
	// x
	if fold_value >= 0 {
		for i := range *dots {
			if (*dots)[i].x > fold_value {
				(*dots)[i].x -= ((*dots)[i].x - fold_value) * 2
				new_position := (*dots)[i].to_string()
				(*dots)[new_position] = (*dots)[i]
				delete(*dots, i)
			}
		}
	} else { // y
		fold_value = -fold_value
		for i := range *dots {
			if (*dots)[i].y > fold_value {
				(*dots)[i].y -= ((*dots)[i].y - fold_value) * 2
				new_position := (*dots)[i].to_string()
				(*dots)[new_position] = (*dots)[i]
				delete(*dots, i)
			}
		}
	}
}

func part_2(dots *map[string]*Pair) {
	fmt.Println("Part 2:")
	max_x, max_y := max_x_y(dots)
	for y := 0; y <= max_y; y++ {
		for x := 0; x <= max_x; x++ {
			key := fmt.Sprintf("%d,%d", x, y)
			if (*dots)[key] != nil {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func max_x_y(dots *map[string]*Pair) (int, int) {
	first_iter := true
	var max_x int
	var max_y int
	for i := range *dots {
		if first_iter {
			max_x = (*dots)[i].x
			max_y = (*dots)[i].y
			first_iter = false
			continue
		}

		if (*dots)[i].x > max_x {
			max_x = (*dots)[i].x
		}
		if (*dots)[i].y > max_y {
			max_y = (*dots)[i].y
		}
	}
	return max_x, max_y
}
