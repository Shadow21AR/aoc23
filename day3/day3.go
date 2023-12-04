package day3

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	file, err := os.Open("day3/day3.txt")
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
	height := len(data) - 1
	re := regexp.MustCompile(`(\d+|(\.)+|(@|#|\$|%|&|\*|-|\+|=|/)+)`)
	var numbers []Numbers
	symbol := make(map[Loc]string)
	for i, line := range data {
		lineChar := re.FindAllString(line, -1)
		var x int
		for _, v := range lineChar {
			numLen := len(v)
			num, err := strconv.Atoi(v)
			if err != nil {
				if !strings.Contains(v, ".") {
					symbol[Loc{x, height - i}] = v
				}
			} else {
				var neighbour []Loc
				for n := 0; n <= numLen+1; n++ {
					nX := x - 1 + n
					neighbour = append(neighbour, Loc{nX, height - i}, Loc{nX, height - i - 1}, Loc{nX, height - i + 1})
				}
				numbers = append(numbers, Numbers{num, Loc{x, height - i}, Loc{x + numLen - 1, height - i}, numLen, neighbour})
			}
			x += numLen
		}
	}
	var sum int
	toInclude := make(map[Loc]int)
	for _, number := range numbers {
		for _, neighbour := range number.neighbours {
			if _, exists := symbol[Loc{neighbour.X, neighbour.Y}]; exists {
				toInclude[number.SLoc] = number.Value
				break
			}
		}
	}
	for _, n := range toInclude {
		sum += n
	}
	fmt.Println("Part1 :", sum)
}

func part2(data []string) {
	height := len(data) - 1
	re := regexp.MustCompile(`(\d+|(\.)+|(@|#|\$|%|&|\*|-|\+|=|/)+)`)
	numbers := make(map[Loc]Loc2)
	gears := make(map[Loc]string)
	for i, line := range data {
		lineChar := re.FindAllString(line, -1)
		var x int
		for _, v := range lineChar {
			numLen := len(v)
			num, err := strconv.Atoi(v)
			if err != nil {
				if strings.Contains(v, "*") {
					gears[Loc{x, height - i}] = v
				}
			} else {
				for l := 0; l <= numLen-1; l++ {
					numbers[Loc{x + l, height - i}] = Loc2{x, height - i, num}
				}
			}
			x += numLen
		}
	}
	gearNeighbour := make(map[Loc]map[Loc]int)
	for pos := range gears {
		temp := make(map[Loc]int)
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if val, ok := numbers[Loc{pos.X + x, pos.Y + y}]; ok {
					temp[Loc{val.X, val.Y}] = val.Value
				}
			}
		}
		gearNeighbour[pos] = temp
	}
	// fmt.Println(gearNeighbour)
	var sum int
	for _, num := range gearNeighbour {
		// fmt.Println(loc, num, len(num))
		if len(num) == 2 {
			temp := 1
			for _, x := range num {
				temp *= x
			}
			sum += temp
		}
	}
	fmt.Println("Part2 :", sum)
}

type Numbers struct {
	Value      int
	SLoc, ELoc Loc
	Length     int
	neighbours []Loc
}

type Loc struct {
	X, Y int
}

type Loc2 struct {
	X, Y, Value int
}
