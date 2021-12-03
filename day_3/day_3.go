package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var rows []string

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	file.Close()

	count := part_1(rows)
	part_2(rows, count)
}

func part_1(rows []string) []uint32 {
	var count []uint32
	for i := 0; i < len(rows[0]); i++ {
		count = append(count, 0)
	}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j] == '1' {
				count[j] += 1
			}
		}
	}

	num_rows := len(rows)
	var gamma_rate uint32 = 0
	var epsilon_rate uint32 = 0
	for i := 0; i < len(count); i++ {
		if count[i] > uint32(num_rows)/2 {
			gamma_rate |= 1 << uint32(len(count)-1-i)
		} else {
			epsilon_rate |= 1 << uint32(len(count)-1-i)
		}
	}

	fmt.Printf("Part 1: %d\n", gamma_rate*epsilon_rate)
	return count
}

func part_2(rows []string, count []uint32) {
	or := get_oxygen_rating(rows, count)
	csr := get_co2_scrubber_rating(rows, count)

	or_val := string_to_number(&or)
	csr_val := string_to_number(&csr)

	fmt.Printf("Part 2: %d\n", or_val*csr_val)
}

func get_oxygen_rating(rows []string, count []uint32) string {
	initial_len := len(rows)
	for i := 0; i < initial_len; i++ {
		one, zero := count_bits(i, &rows)
		if one >= zero {
			rows = new_rows(i, '1', &rows)
		} else {
			rows = new_rows(i, '0', &rows)
		}

		if len(rows) == 1 {
			break
		}
	}
	return rows[0]
}

func get_co2_scrubber_rating(rows []string, count []uint32) string {
	initial_len := len(rows)
	for i := 0; i < initial_len; i++ {
		one, zero := count_bits(i, &rows)
		if one < zero {
			rows = new_rows(i, '1', &rows)
		} else {
			rows = new_rows(i, '0', &rows)
		}

		if len(rows) == 1 {
			break
		}
	}
	return rows[0]
}

func count_bits(pos int, rows *[]string) (int, int) {
	zero := 0
	one := 0
	for i := 0; i < len(*rows); i++ {
		if (*rows)[i][pos] == '1' {
			one += 1
		} else {
			zero += 1
		}
	}

	return one, zero
}

func new_rows(pos int, value byte, rows *[]string) []string {
	var new_rows []string
	for i := 0; i < len(*rows); i++ {
		if (*rows)[i][pos] == value {
			new_rows = append(new_rows, (*rows)[i])
		}
	}

	return new_rows
}

func string_to_number(str *string) uint32 {
	var val uint32 = 0
	for i := 0; i < len(*str); i++ {
		if (*str)[i] == '1' {
			val |= 1 << uint32(len(*str)-1-i)
		}
	}

	return val
}
