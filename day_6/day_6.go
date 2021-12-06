package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var smart_fish [10]uint64

	scanner.Scan()
	for _, str := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(str)
		smart_fish[num] += 1
	}

	file.Close()

	fmt.Printf("Part 1: %d\n", simulate_fish(smart_fish, 80))
	fmt.Printf("Part 2: %d\n", simulate_fish(smart_fish, 256))
}

func simulate_fish(fish [10]uint64, days int) uint64 {
	for i := 0; i < days; i++ {
		fish[7] += fish[0]
		fish[9] += fish[0]
		for j := 0; j < len(fish); j++ {
			if j == len(fish)-1 {
				fish[j] = 0
			} else {
				fish[j] = fish[j+1]
			}
		}
	}

	var count uint64 = 0
	for i, _ := range fish {
		count += fish[i]
	}

	return count
}
