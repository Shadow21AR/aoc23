package day5

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day5() {
	// startTime := time.Now()
	file, err := os.Open("day5/day5.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	data := (strings.Split(string(content), "\r\n\r\n"))
	part(data) //part1
	//part2...nah
}
func part(data []string) {
	almanac := make(map[int][8]int)
	for _, line := range data {
		instruction := strings.Split(line, ":")
		item := instruction[0]
		switch {
		case strings.Contains(item, "seeds"):
			addSeeds(&almanac, instruction[1:])
		case strings.Contains(item, "to-soil"):
			almanac = addStuffs(almanac, instruction[1:], 1)
		case strings.Contains(item, "to-fertilizer"):
			almanac = addStuffs(almanac, instruction[1:], 2)
		case strings.Contains(item, "to-water"):
			almanac = addStuffs(almanac, instruction[1:], 3)
		case strings.Contains(item, "to-light"):
			almanac = addStuffs(almanac, instruction[1:], 4)
		case strings.Contains(item, "to-temperature"):
			almanac = addStuffs(almanac, instruction[1:], 5)
		case strings.Contains(item, "to-humidity"):
			almanac = addStuffs(almanac, instruction[1:], 6)
		case strings.Contains(item, "to-location"):
			almanac = addStuffs(almanac, instruction[1:], 7)
		}
	}
	lowest := math.Inf(1)
	for _, list := range almanac {
		lowest = math.Min(lowest, float64(list[len(list)-1]))
		// fmt.Println(id, list)
	}
	fmt.Printf("Part 1: %d\n", int(lowest))
}
func addSeeds(almanac *(map[int][8]int), seeds []string) {
	for _, s := range strings.Split(strings.Join(seeds, " "), " ") {
		n, err := strconv.Atoi(s)
		if err == nil {
			value := (*almanac)[n]
			value[0] = n
			(*almanac)[n] = value
		}
	}
}
func addStuffs(almanac map[int][8]int, instruction []string, index int) (res map[int][8]int) {
	res = almanac
	for _, s := range strings.Split(strings.Join(instruction, " "), "\r\n") {
		if len(s) == 0 {
			continue
		}
		row := toInt(strings.Split(s, " "))
		if canSkip(&res, row, index) {
			continue
		}
		for id, list := range res {
			if isInRange(list[index-1], row[1], row[1]+row[2]) {
				if list[index] != 0 {
					continue
				}
				list[index] = list[index-1] + (row[0] - row[1])
				res[id] = list
			}
		}
	}
	fillMap(&res, index)
	return res
}
func canSkip(almanac *(map[int][8]int), row []int, index int) bool {
	res := true
	for _, x := range *almanac {
		if isInRange(x[index-1], row[1], row[1]+row[2]) {
			res = false
		}
	}
	return res
}
func fillMap(res *map[int][8]int, index int) {
	for id, list := range *res {
		if list[index] == 0 {
			list[index] = list[index-1]
			(*res)[id] = list
		}
	}
}
func isInRange(number, min, max int) bool {
	return number >= min && number <= max
}
func toInt(s []string) (n []int) {
	for _, x := range s {
		y, _ := strconv.Atoi(x)
		n = append(n, y)
	}
	return
}
