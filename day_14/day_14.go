package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	polymer_template := []rune(scanner.Text())
	scanner.Scan()

	polymer_template2 := make(map[string]uint64)

	for i := 0; i < len(polymer_template) - 1; i++ {
		polymer_template2[string(polymer_template[i:i+2])] += 1
	}

	polymers := make(map[rune]uint64)
	for _, polymer := range polymer_template {
		polymers[polymer] += 1
	}

	polymer_pairs := make(map[string]rune)
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), " -> ")
		polymer_pairs[pairs[0]] = []rune(pairs[1])[0]
	}

	part_1(&polymer_template2, &polymers, &polymer_pairs)
	part_2(&polymer_template2, &polymers, &polymer_pairs)
}

func part_1(polymer_template *map[string]uint64, polymers *map[rune]uint64, polymer_pairs *map[string]rune) {
	for i := 0; i < 10; i++ {
		grow_polymer(polymer_template, polymers, polymer_pairs)
	}
	min, max := min_max(polymers)
	fmt.Printf("Part 1: %d\n", max - min)
}

func min_max(polymers *map[rune]uint64) (uint64, uint64) {
	is_first := true
	var min uint64
	var max uint64
	for key := range *polymers {
		if is_first {
			min = (*polymers)[key]
			max = (*polymers)[key]
			is_first = false
			continue
		}

		if min > (*polymers)[key] {
			min = (*polymers)[key]
		}

		if max < (*polymers)[key] {
			max = (*polymers)[key]
		}
	}

	return min, max
}

func part_2(polymer_template *map[string]uint64, polymers *map[rune]uint64, polymer_pairs *map[string]rune) {
	for i := 10; i < 40; i++ {
		grow_polymer(polymer_template, polymers, polymer_pairs)
	}
	min, max := min_max(polymers)
	fmt.Printf("Part 2: %d\n", max - min)
}

func grow_polymer(polymer_template *map[string]uint64, polymers *map[rune]uint64, polymer_pairs *map[string]rune) {
	pt_before := make(map[string]uint64)
	// volim koperati
	for k,v := range *polymer_template {
		pt_before[k] = v
	}

	for key := range pt_before {
		repeat_times := pt_before[key]
		new_polymer := (*polymer_pairs)[key]
		first_key := fmt.Sprintf("%s%s", string(key[0:1]), string(new_polymer))
		second_key := fmt.Sprintf("%s%s", string(new_polymer), string(key[1:2]))

		(*polymers)[new_polymer] += repeat_times

		(*polymer_template)[key] -= repeat_times
		(*polymer_template)[first_key] += repeat_times
		(*polymer_template)[second_key] += repeat_times
		
		if (*polymer_template)[key] == 0 {
			delete(*polymer_template, key)
		}
	}
}

