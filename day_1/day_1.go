package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("./input.txt")

	var input_data []int

	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		input_data = append(input_data, i)
	}

	var part_1 = 0
	for i := 1; i < len(input_data); i++ {
		if input_data[i] > input_data[i-1] {
			part_1 += 1
		}
	}

	fmt.Printf("Part 1: %d\n", part_1)

	var part_2 = 0

	var group1 = input_data[0] + input_data[1] + input_data[2]
	var group2 = 0

	for i := 1; i < len(input_data); i++ {
		if i+2 < len(input_data) {
			group2 = input_data[i] + input_data[i+1] + input_data[i+2]
		} else {
			continue
		}

		if group1 < group2 {
			part_2 += 1
		}

		group1 = group2
	}

	fmt.Printf("Part 2: %d\n", part_2)

	file.Close()
}
