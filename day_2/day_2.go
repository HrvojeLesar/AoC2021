package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PositionP1 struct {
	horizontal int32
	depth      int32
}

func (p *PositionP1) passCommand(command *[]string) {
	units, _ := strconv.Atoi((*command)[1])
	switch (*command)[0][0] {
	case 'f':
		p.horizontal += int32(units)
	case 'd':
		p.depth += int32(units)
	case 'u':
		p.depth -= int32(units)
	}
}

type PositionP2 struct {
	horizontal int32
	depth      int32
	aim        int32
}

func (p *PositionP2) passCommand(command *[]string) {
	units, _ := strconv.Atoi((*command)[1])
	switch (*command)[0][0] {
	case 'f':
		p.horizontal += int32(units)
		p.depth += p.aim * int32(units)
	case 'd':
		p.aim += int32(units)
	case 'u':
		p.aim -= int32(units)
	}
}

func main() {
	file, _ := os.Open("input.txt")

	position_p1 := PositionP1{
		horizontal: 0,
		depth:      0,
	}

	position_p2 := PositionP2{
		horizontal: 0,
		depth:      0,
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		position_p1.passCommand(&command)
		position_p2.passCommand(&command)
	}

	fmt.Printf("Part 1: %d\n", position_p1.horizontal*position_p1.depth)
	fmt.Printf("Part 2: %d\n", position_p2.horizontal*position_p2.depth)

	file.Close()
}
