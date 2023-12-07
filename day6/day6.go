package day6

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day6() {
	file, err := os.Open("day6/day6.txt")
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
	data := (strings.Split(string(content), "\r\n"))
	var input [][]int
	var input2 []int
	re := regexp.MustCompile(`\d+`)
	for _, v := range data {
		line := toInt(re.FindAllString(v, -1))
		input = append(input, line)
		n, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(v, ":")[1], " ", ""))
		input2 = append(input2, n)
	}
	part1(input)
	part2(input2)
}
func part2(input []int) {
	fmt.Println("Part 2: ", len(findCombo(input[0], input[1])))
}
func part1(input [][]int) {
	var combo [][][]int
	for i := 0; i < len(input[0]); i++ {
		combo = append(combo, findCombo(input[0][i], input[1][i]))
	}
	fmt.Println("Part 1: ", count(combo))
}
func findCombo(n, z int) (combo [][]int) {
	for x := 0; x < n; x++ {
		y := n - x
		if x*y > z {
			combo = append(combo, []int{x, y})
		}
	}
	return
}
func toInt(s []string) (out []int) {
	for _, x := range s {
		n, _ := strconv.Atoi(x)
		out = append(out, n)
	}
	return
}
func count(input [][][]int) int {
	count := 1
	for _, x := range input {
		count *= len(x)
	}
	return count
}
