package day1

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day1() {
	file, err := os.Open("day1/day1.txt")
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
	part1(data)
	part2(data)
}

func part1(data []string) {
	var cal []int
	re := regexp.MustCompile(`\d`)
	for _, s := range data {
		nums := re.FindAllString(s, -1)
		if nums != nil {
			first := nums[0]
			last := nums[len(nums)-1]
			num, _ := strconv.Atoi(first + last)
			cal = append(cal, num)
		}
	}

	sum := 0
	for _, n := range cal {
		sum += n
	}
	fmt.Println("Part 1: ", sum)
}

func part2(data []string) {
	var cal []int
	numMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	test := strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four",
		"f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e") //:<
	re := regexp.MustCompile(`(?:one|two|three|four|five|six|seven|eight|nine|\d)`)

	for _, s := range data {
		nums := re.FindAllStringSubmatch(test.Replace(s), -1)
		first, err := strconv.Atoi(nums[0][0])
		if err != nil {
			first = numMap[nums[0][0]]
		}
		last, err := strconv.Atoi(nums[len(nums)-1][0])
		if err != nil {
			last = numMap[(nums[len(nums)-1])[0]]
		}
		num := first*10 + last
		cal = append(cal, num)
	}

	sum2 := 0
	for _, n := range cal {
		sum2 += n
	}
	fmt.Println("Part 2: ", sum2)
}
