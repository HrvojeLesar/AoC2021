package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func median(numbers []float64) float64 {
	length := len(numbers)
	sort.Float64s(numbers)
	if length%2 == 0 {
		return (numbers[length/2-1] + numbers[length/2]) / 2
	} else {
		return numbers[length/2-1]
	}
}

func get_fuel_values(move_to float64, horizontal_positions []float64) float64 {
	sum := 0.0
	for _, pos := range horizontal_positions {
		sum += math.Abs(pos - move_to)
	}
	return sum
}

func sum_to(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func brute_force_p2(positions []float64) int {
	sort.Float64s(positions)
	max := positions[len(positions)-1]
	var sums []int
	for i := 0.0; i <= max; i++ {
		sum := 0
		for _, pos := range positions {
			sum += sum_to(int(math.Abs(pos - i)))
		}
		sums = append(sums, sum)
	}

	sort.Ints(sums)

	return sums[0]
}

func main() {

	file, _ := os.Open("input.txt")

	var horiontal_positions []float64
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, str := range strings.Split(scanner.Text(), ",") {
		horizontal_position, _ := strconv.Atoi(str)
		horiontal_positions = append(horiontal_positions, float64(horizontal_position))
	}

	file.Close()

	median := median(horiontal_positions)

	fmt.Printf("Part 1: %f\n", get_fuel_values(median, horiontal_positions))
	fmt.Printf("Part 2: %d\n", brute_force_p2(horiontal_positions))
}
