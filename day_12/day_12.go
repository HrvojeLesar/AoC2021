package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cave struct {
	name        string
	connects_to []string
	is_small    bool
}

func NewCave(name string) *Cave {
	first_char := []rune(name)[0]
	return &Cave{
		name:        name,
		connects_to: make([]string, 0),
		is_small:    first_char >= 97 && first_char <= 122,
	}
}

func (c *Cave) ConnectTo(ct string) {
	c.connects_to = append(c.connects_to, ct)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	caves := make(map[string]*Cave)

	for scanner.Scan() {
		connected_caves := strings.Split(scanner.Text(), "-")
		_, first_cave_exists := caves[connected_caves[0]]
		_, second_cave_exists := caves[connected_caves[1]]
		if !first_cave_exists {
			caves[connected_caves[0]] = NewCave(connected_caves[0])
		}
		caves[connected_caves[0]].ConnectTo(connected_caves[1])

		if !second_cave_exists {
			caves[connected_caves[1]] = NewCave(connected_caves[1])
		}
		caves[connected_caves[1]].ConnectTo(connected_caves[0])
	}
	path := []string{"start"}
	total_paths := 0
	total_paths_2 := 0
	for _, cave_name := range caves["start"].connects_to {
		find_path(cave_name, caves, path, &total_paths)
		find_path2(cave_name, caves, path, &total_paths_2, false)
	}
	fmt.Printf("Part 1: %d\n", total_paths)
	fmt.Printf("Part 2: %d\n", total_paths_2)
}

func find_path(cave_name string, caves map[string]*Cave, path []string, total_paths *int) {
	path_copy := path
	path_copy = append(path_copy, cave_name)
	for _, cave := range caves[cave_name].connects_to {
		if caves[cave].name == "start" {
			continue
		}
		if caves[cave].name == "end" {
			*total_paths += 1
			continue
		}
		if caves[cave].is_small {
			if !is_in_path(cave, &path_copy) {
				find_path(cave, caves, path_copy, total_paths)
			}
		} else {
			find_path(cave, caves, path_copy, total_paths)
		}
	}
}

func find_path2(cave_name string, caves map[string]*Cave, path []string, total_paths *int, small_used_twice bool) {
	path_copy := path
	path_copy = append(path_copy, cave_name)
	for _, cave := range caves[cave_name].connects_to {
		if caves[cave].name == "start" {
			continue
		}
		if caves[cave].name == "end" {
			*total_paths += 1
			continue
		}
		if caves[cave].is_small {
			is_in_path := is_in_path(cave, &path_copy)
			if !small_used_twice && is_in_path {
				find_path2(cave, caves, path_copy, total_paths, true)
			} else if !is_in_path {
				find_path2(cave, caves, path_copy, total_paths, small_used_twice)
			}
		} else {
			find_path2(cave, caves, path_copy, total_paths, small_used_twice)
		}
	}
}

func is_in_path(val string, path *[]string) bool {
	contains := false
	for _, p := range *path {
		if p == val {
			contains = true
			break
		}
	}
	return contains
}
