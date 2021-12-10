package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var illegal_scores_const = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var incomplete_scores_const = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var bracket_pairs_const = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	illegal_count := make(map[rune]int)
	var incomplete_scores []int
	for scanner.Scan() {
		find_illegal(scanner.Text(), &illegal_count, &incomplete_scores)
	}
	part_1(&illegal_count)
	part_2(&incomplete_scores)
}

func find_illegal(line string, illegal_count *map[rune]int, incomplete_scores *[]int) {
	var expected_closing_bracket []rune
	for _, char := range line {
		closing_bracket, is_opening := bracket_pairs_const[char]
		if is_opening {
			expected_closing_bracket = append(expected_closing_bracket, closing_bracket)
		} else {
			expected := pop(&expected_closing_bracket)
			if expected != char {
				(*illegal_count)[char] += 1
				return
			}
		}
	}

	score := 0
	for i := len(expected_closing_bracket) - 1; i >= 0; i-- {
		score *= 5
		score += incomplete_scores_const[expected_closing_bracket[i]]
	}
	*incomplete_scores = append(*incomplete_scores, score)
}

func pop(runes *[]rune) rune {
	ret := (*runes)[len(*runes)-1]
	*runes = (*runes)[:len(*runes)-1]
	return ret
}

func part_1(illegal_count *map[rune]int) {
	score := 0
	for key, _ := range *illegal_count {
		score += illegal_scores_const[key] * (*illegal_count)[key]
	}
	fmt.Printf("Part 1: %d\n", score)
}

func part_2(incomplete_scores *[]int) {
	sort.Ints(*incomplete_scores)
	fmt.Printf("Part 2: %d\n", (*incomplete_scores)[len(*incomplete_scores)/2])
}
