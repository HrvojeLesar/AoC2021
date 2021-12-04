package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoCard struct {
	card_numbers []int
	mask         []bool
}

func new_bingo_card() *BingoCard {
	return &BingoCard{
		mask: make([]bool, 25),
	}
}

func (p *BingoCard) mark_num(number int) {
	for i := 0; i < 25; i++ {
		if p.card_numbers[i] == number {
			p.mask[i] = true
		}
	}
}

func (p *BingoCard) is_winner() bool {
	// check rows
	is_winner := false
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !p.mask[5*i+j] {
				break
			}

			if j == 4 {
				is_winner = true
			}
		}
	}

	if is_winner {
		return is_winner
	}

	// check columns
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !p.mask[5*j+i] {
				break
			}

			if j == 4 {
				is_winner = true
			}
		}
	}

	return is_winner
}

func (p *BingoCard) sum_unmarked() int {
	sum := 0
	for i := 0; i < 25; i++ {
		if !p.mask[i] {
			sum += p.card_numbers[i]
		}
	}

	return sum
}

func main() {

	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var bingo_numbers []int
	var bingo_cards []BingoCard

	scanner.Scan()
	bingo_numbers = parse_numbers(scanner.Text(), ",")
	scanner.Scan()

	var bingo_card_numbers []int
	for {
		scan_res := scanner.Scan()
		if scanner.Text() != "" {
			bingo_card_numbers = append(bingo_card_numbers, parse_numbers(scanner.Text(), " ")...)
		} else {
			card := new_bingo_card()
			card.card_numbers = bingo_card_numbers
			bingo_cards = append(bingo_cards, *card)
			bingo_card_numbers = nil
		}

		if !scan_res {
			break
		}
	}

	sum, number := part_1(&bingo_numbers, bingo_cards)
	sum2, number2 := part_2(&bingo_numbers, bingo_cards)
	fmt.Println(sum * number)
	fmt.Println(sum2 * number2)
}

func part_1(bingo_numbers *[]int, bingo_cards []BingoCard) (int, int) {

	for i, num := range *bingo_numbers {
		for _, card := range bingo_cards {
			card.mark_num(num)
			if i >= 4 {
				if card.is_winner() {
					return card.sum_unmarked(), num
				}
			}
		}
	}

	return -1, -1
}

func part_2(bingo_numbers *[]int, bingo_cards []BingoCard) (int, int) {
	win_nums := 0
	win_mask := make([]bool, len(bingo_cards))
	for i, num := range *bingo_numbers {
		for card_index, card := range bingo_cards {
			if !win_mask[card_index] {
				card.mark_num(num)
				if i >= 4 {
					if card.is_winner() {
						win_mask[card_index] = true
						win_nums++
						if win_nums == len(bingo_cards) {
							return card.sum_unmarked(), num
						}
					}
				}
			}
		}
	}

	return -1, -1
}

func parse_numbers(numbers_string string, delimiter string) []int {
	var bingo_numbers []int
	split_numbers := strings.Split(numbers_string, delimiter)
	for _, str := range split_numbers {
		if str == "" {
			continue
		}
		num, _ := strconv.Atoi(str)
		bingo_numbers = append(bingo_numbers, num)
	}
	return bingo_numbers
}
