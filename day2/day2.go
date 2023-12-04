package day2

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day2() {
	file, err := os.Open("day2/day2.txt")
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
	re := regexp.MustCompile(`(\d+)\s+(red|green|blue)`)
	var games []Game
	for i, line := range data {
		var r, g, b float64
		cubes := re.FindAllString(line, -1)
		for _, cube := range cubes {
			temp := strings.Split(cube, " ")
			color := temp[1]
			count, err := strconv.ParseFloat(temp[0], 64)
			if err != nil {
				fmt.Println("Error converting to float64: ", err)
				return
			}
			switch {
			case color == "red":
				r = math.Max(r, count)
			case color == "green":
				g = math.Max(g, count)
			case color == "blue":
				b = math.Max(b, count)
			}
		}
		games = append(games, Game{i + 1, r, g, b})
	}
	var possible int
	var redCheck float64 = 12
	var greenCheck float64 = 13
	var blueCheck float64 = 14
	for _, game := range games {
		if game.Red <= redCheck && game.Blue <= blueCheck && game.Green <= greenCheck {
			possible += game.NumG
		}
	}
	fmt.Println("Part 1: ", possible)
}

func part2(data []string) {
	re := regexp.MustCompile(`(\d+)\s+(red|green|blue)`)
	var games2 []Game
	for i, line := range data {
		var r, g, b float64
		cubes := re.FindAllString(line, -1)
		for _, cube := range cubes {
			temp := strings.Split(cube, " ")
			color := temp[1]
			count, err := strconv.ParseFloat(temp[0], 64)
			if err != nil {
				fmt.Println("Error converting to float64: ", err)
				return
			}
			switch {
			case color == "red":
				r = math.Max(r, count)
			case color == "green":
				g = math.Max(g, count)
			case color == "blue":
				b = math.Max(b, count)
			}
		}
		games2 = append(games2, Game{i + 1, r, g, b})
	}
	var power int
	for _, game := range games2 {
		power += int(game.Red) * int(game.Blue) * int(game.Green)
	}
	fmt.Println("Part 2: ", power)
}

type Game struct {
	NumG  int
	Red   float64
	Green float64
	Blue  float64
}
