package main

import (
	"bufio"
	"fmt"
	"os"
)

var hex_to_binary_const = map[rune][]rune{
	'0': {'0', '0', '0', '0'},
	'1': {'0', '0', '0', '1'},
	'2': {'0', '0', '1', '0'},
	'3': {'0', '0', '1', '1'},
	'4': {'0', '1', '0', '0'},
	'5': {'0', '1', '0', '1'},
	'6': {'0', '1', '1', '0'},
	'7': {'0', '1', '1', '1'},
	'8': {'1', '0', '0', '0'},
	'9': {'1', '0', '0', '1'},
	'A': {'1', '0', '1', '0'},
	'B': {'1', '0', '1', '1'},
	'C': {'1', '1', '0', '0'},
	'D': {'1', '1', '0', '1'},
	'E': {'1', '1', '1', '0'},
	'F': {'1', '1', '1', '1'},
}

// packet header
// 0-2 packet version
// 3-5 type ID

// literal type
// 5 bit groups
// last group prefix 0

// operator type
// 6 lenght type ID (0 | 1)

var version_count uint = 0

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	packet := hex_packet_to_bin(scanner.Text())
	_, result := parse_packet(packet, true)
	fmt.Printf("Part 1: %d\n", version_count)
	fmt.Printf("Part 2: %d\n", result)
}

func hex_packet_to_bin(packet string) []rune {
	var binary_runes []rune
	for _, r := range []rune(packet) {
		binary_runes = append(binary_runes, hex_to_binary_const[r]...)
	}
	return binary_runes
}

func parse_packet(packet []rune, is_main bool) (int, uint) {
	var literal uint = 0
	var values []uint = []uint{}
	bits_read := 6

	version := append([]rune{'0'}, packet[0:3]...)
	type_id := append([]rune{'0'}, packet[3:6]...)
	version_count += rune_to_int(version)
	type_id_uint := rune_to_int(type_id)
	if bits_read > len(packet) {
		return bits_read, 0
	}

	var bits_used int = 0
	switch type_id_uint {
	case 4:
		literal, bits_used = parse_literal(packet[bits_read:])
		bits_read += bits_used
	default:
		bits_used, values = parse_operator(packet[bits_read:])
		bits_read += bits_used
	}

	switch type_id_uint {
	case 0:
		var sum uint = 0	
		for _, v := range values {
			sum += v
		}
		literal = sum
	case 1:
		var product uint = 1	
		for _, v := range values {
			product *= v
		}
		literal = product
	case 2:
		min := values[0]
		for _, v := range values {
			if min > v {
				min = v
			}
		}
		literal = min
	case 3:
		max := values[0]
		for _, v := range values {
			if max < v {
				max = v
			}
		}
		literal = max
	case 5:
		greater_than := 0
		if values[0] > values[1] {
			greater_than = 1
		}
		literal = uint(greater_than)
	case 6:
		less_than := 0
		if values[0] < values[1] {
			less_than = 1
		}
		literal = uint(less_than)
	case 7:
		equal_to := 0
		if values[0] == values[1] {
			equal_to = 1
		}
		literal = uint(equal_to)
	}

	if is_main {
		for bits_read % 4 != 0 {
			bits_read += 1
		}
	}

	if is_main && bits_read < len(packet) {
		parse_packet(packet[bits_read:], true)
	}

	return bits_read, literal
}

func parse_literal(packet []rune) (uint, int) {
	i := 0
	var literal []rune
	for packet[i:i + 5][0] != '0' {
		literal = append(literal, packet[i+1:i+5]...)
		i += 5
	}
	return rune_to_int(append(literal, packet[i+1:i+5]...)), i + 5
}

func parse_operator(packet []rune) (int, []uint) {
	length_type_id := packet[0]
	bits_used := 1
	values := []uint{}
	if length_type_id == '0' {
		length := rune_to_int(packet[1:15 + 1])
		bits_used += 15
		for length > 0 {
			bits, value := parse_packet(packet[bits_used:bits_used + int(length) + 1], false)
			length -= uint(bits)	
			bits_used += bits
			values = append(values, value)
		}
	} else {
		num_of_subpackets := rune_to_int(packet[1:11 + 1])
		bits_used += 11
		for i := 1; i <= int(num_of_subpackets); i++ {
			bits, value := parse_packet(packet[bits_used:], false) 
			bits_used += bits
			values = append(values, value)
		}
	}

	return bits_used, values
}

func rune_to_int(packet []rune) uint {
	var num uint = 0
	for _, r := range packet {
		if r == '0' {
			num = num << 1
		} else {
			num = (num << 1) | 1
		}
	}
	return num
}
