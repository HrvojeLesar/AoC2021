// 0: a,b,c,e,f,g 	| 6
// 1: c,f			| 2
// 2: a,c,d,e,g		| 5
// 3: a,c,d,f,g		| 5
// 4: b,c,d,f		| 4
// 5: a,b,d,f,g		| 5
// 6: a,b,d,e,f,g	| 6
// 7: a,c,f			| 3
// 8: a,b,c,d,e,f,g | 7
// 9: a,b,c,d,f,g	| 6

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	part_1_count := 0
	var part_2_count uint64 = 0
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " | ")
		part_1_count += part_1(input[1])
		part_2_count += part_2(input)
	}

	fmt.Printf("Part 1: %d\n", part_1_count)
	fmt.Printf("Part 2: %d\n", part_2_count)
}

func part_1(output_values string) int {
	count := 0
	values := strings.Split(output_values, " ")
	for _, val := range values {
		val_len := len(val)
		if val_len == 2 || val_len == 4 || val_len == 3 || val_len == 7 {
			count += 1
		}
	}
	return count
}

func part_2(input []string) uint64 {
	// 0: top, 1: left, 2: right
	// 3: sredina
	// 4: left, 5: right, 6: down
	var positions [7]rune
	uniques := strings.Split(input[0], " ")
	values := sort_array(strings.Split(input[1], " "))
	digits, others := get_uniques_others(uniques)
	positions[0] = not_in_second(digits[7], digits[1])
	find_zero_six_nine(others, digits[1], &digits, &positions)
	find_two_three_five(others, digits[1], &digits, &positions)
	sort_digits(&digits)
	return parse(&digits, values)
}

func get_uniques_others(in []string) ([10]([]rune), []([]rune)) {
	var digits [10]([]rune)
	var others []([]rune)
	for _, val := range in {
		val_len := len(val)
		switch val_len {
		case 2:
			// 1
			digits[1] = []rune(val)
		case 3:
			// 7
			digits[7] = []rune(val)
		case 4:
			// 4
			digits[4] = []rune(val)
		case 7:
			// 8
			digits[8] = []rune(val)
		default:
			others = append(others, []rune(val))
		}
	}
	return digits, others
}

// actual garbage function, who knows how it works
func not_in_second(first []rune, second []rune) rune {
	for _, r := range first {
		if !strings.ContainsRune(string(second), r) {
			return r
		}
	}
	return -1
}

func find_zero_six_nine(others []([]rune), one []rune, digits *[10]([]rune), positions *[7]rune) {
	var six_nine []([]rune)
	for _, other := range others {
		if len(other) == 6 {
			// dobim 0 i 9, neznam koji je koji
			if strings.ContainsRune(string(other), one[0]) && strings.ContainsRune(string(other), one[1]) {
				six_nine = append(six_nine, other)
			} else {
				digits[6] = other
				(*positions)[2] = not_in_second(one, (*digits)[6])
			}
		}
	}

	if not_in_second(digits[4], six_nine[0]) == -1 {
		digits[0] = six_nine[1]
		digits[9] = six_nine[0]
	} else {
		digits[0] = six_nine[0]
		digits[9] = six_nine[1]
	}
}

func find_two_three_five(others []([]rune), one []rune, digits *[10]([]rune), positions *[7]rune) {
	for _, other := range others {
		if len(other) == 5 {
			if strings.ContainsRune(string(other), one[0]) && strings.ContainsRune(string(other), one[1]) {
				digits[3] = other
			} else {
				if strings.ContainsRune(string(other), (*positions)[2]) {
					// 2
					digits[2] = other
				} else {
					digits[5] = other
				}
			}
		}
	}
}

func sort_digits(digits *[10]([]rune)) {
	for _, d := range *digits {
		sort.Slice(d, func(i, j int) bool {
			return d[i] < d[j]
		})
	}
}

func sort_string(in string) string {
	var runes []rune
	for _, r := range in {
		runes = append(runes, r)
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func sort_array(in []string) []string {
	var out []string
	for _, val := range in {
		out = append(out, sort_string(val))
	}
	return out
}

func parse(digits *[10]([]rune), values []string) uint64 {
	var out uint64 = 0
	for _, val := range values {
		for i, digit := range *digits {
			if strings.Compare(val, string(digit)) == 0 {
				out = out*uint64(10) + uint64(i)
			}
		}
	}
	return out
}
