package day4

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day4() {
	file, err := os.Open("day4/day4.txt")
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
	re := regexp.MustCompile(`\d+`)
	cards := make(map[int]Card)
	for _, line := range data {
		nums := toInt(re.FindAllString(line, -1))
		// cards[nums[0]] = Card{1, nums[1:6], nums[6:]} //example
		cards[nums[0]] = Card{1, nums[1:11], nums[11:]} //part 1
	}
	part1(cards)
	part2(cards)
}

func part1(cards map[int]Card) {
	var sum int
	for _, card := range cards {
		var score float64
		for _, n := range card.Have {
			if contains(n, card.Win) {
				score++
			}
		}
		if score > 0 {
			sum += int(math.Pow(2, score-1))
		}
	}
	fmt.Println("Part 1:", sum)
}

func part2(cards map[int]Card) {
	for id := 1; id < len(cards); id++ {
		card := cards[id]
		var match float64
		for _, n := range card.Have {
			if contains(n, card.Win) {
				match++
			}
		}
		if match > 0 {
			for i := 0; i < int(math.Min(match, float64(len(cards)))); i++ {
				next := cards[id+i+1]
				next.Count = next.Count + 1*card.Count
				cards[id+i+1] = next
			}
		}
	}
	var sum int
	for _, card := range cards {
		sum += card.Count
	}
	fmt.Println("Part 2:", sum)
}

func toInt(strings []string) (numbers []int) {
	for _, s := range strings {
		num, _ := strconv.Atoi(s)
		numbers = append(numbers, num)
	}
	return
}

func contains(n int, win []int) bool {
	for _, x := range win {
		if x == n {
			return true
		}
	}
	return false
}

type Card struct {
	Count int
	Win   []int
	Have  []int
}
