package day7

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
}

func Day7() { //copied cuz my head herts https://github.com/mnml/aoc/blob/main/2023/07/1.go
	input, _ := os.ReadFile("day7/day7.txt")

	hands := []Hand{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		h := Hand{}
		fmt.Sscanf(s, "%s %d", &h.Cards, &h.Bid)
		hands = append(hands, h)
	}

	winnings := func(jokers bool) (w int) {
		slices.SortFunc(hands, func(a, b Hand) int {
			return cmp(a.Cards, b.Cards, jokers)
		})
		for i, h := range hands {
			w += (i + 1) * h.Bid
		}
		return
	}

	fmt.Println(winnings(false))
	fmt.Println(winnings(true))
}

func cmp(a, b string, jokers bool) int {
	j, r := "J", "TAJBQCKDAE"
	if jokers {
		j, r = "23456789TQKA", "TAJ0QCKDAE"
	}

	typ := func(cards string) string {
		k := 0
		for _, j := range strings.Split(j, "") {
			n, t := strings.ReplaceAll(cards, "J", j), 0
			for _, s := range n {
				t += strings.Count(n, string(s))
			}
			k = slices.Max([]int{k, t})
		}
		return map[int]string{5: "0", 7: "1", 9: "2", 11: "3", 13: "4", 17: "5", 25: "6"}[k]
	}

	return strings.Compare(
		typ(a)+strings.NewReplacer(strings.Split(r, "")...).Replace(a),
		typ(b)+strings.NewReplacer(strings.Split(r, "")...).Replace(b),
	)
}

// func Day7() {

// 	file, err := os.Open("day7/day7.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()
// 	content, err := io.ReadAll(file)
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}
// 	data := (strings.Split(string(content), "\r\n"))
// 	deck := make(map[string]Hand)
// 	for _, v := range data {
// 		line := strings.Split(v, " ")
// 		deck[line[0]] = Hand{toInt(line[1]), 0}
// 	}
// 	part1(deck)
// }
// func part1(deck map[string]Hand) {
// 	for face, hand := range deck {
// 		faceSplit := strings.Split(face, "")
// 		sort.Slice(faceSplit, func(i, j int) bool {
// 			return faceSplit[i] < faceSplit[j]
// 		})
// 		hand.Strength = checkHand(faceSplit)
// 		deck[face] = hand
// 	}
// 	assignStrengths(&deck)
// 	var winnings int
// 	for _, hand := range deck {
// 		winnings += hand.Bid * hand.Strength
// 	}
// 	fmt.Println(winnings)
// }
// func assignStrengths(inputMap *(map[string]Hand)) {
// 	tempMap := make(map[int][]string)
// 	for k, v := range *inputMap {
// 		tempArray := tempMap[v.Strength]
// 		tempArray = append(tempArray, k)
// 		tempMap[v.Strength] = tempArray
// 	}
// 	strength := 1
// 	for i := 1; i <= 7; i++ {
// 		if list, ok := tempMap[i]; ok {
// 			if len(list) == 1 {
// 				temp := (*inputMap)[list[0]]
// 				temp.Strength = strength
// 				(*inputMap)[list[0]] = temp
// 				strength++
// 			} else {
// 				sort.Sort(sort.Reverse(sort.StringSlice(list)))
// 				for _, key := range list {
// 					temp := (*inputMap)[key]
// 					temp.Strength = strength
// 					(*inputMap)[key] = temp
// 					strength++
// 				}
// 			}
// 		}
// 	}
// }
// func checkHand(faces []string) int {
// 	handType := 1
// 	if fiveKind(faces) {
// 		handType = 7
// 	} else if fourKind(faces) {
// 		handType = 6
// 	} else if fullHouse(faces) {
// 		handType = 5
// 	} else if threeKind(faces) {
// 		handType = 4
// 	} else if twoPair(faces) {
// 		handType = 3
// 	} else if onePair(faces) {
// 		handType = 2
// 	}
// 	return handType
// }
// func toInt(s string) (n int) {
// 	n, _ = strconv.Atoi(s)
// 	return
// }

// //	func max(x, y int) int {
// //		if x > y {
// //			return x
// //		}
// //		return y
// //	}
// func fiveKind(faces []string) bool {
// 	return strings.Count(strings.Join(faces, ""), faces[0]) == 5
// }
// func fourKind(faces []string) bool {
// 	return strings.Count(strings.Join(faces, ""), faces[0]) == 4
// }
// func fullHouse(faces []string) bool {
// 	if (strings.Count(strings.Join(faces, ""), faces[0]) == 3) && strings.Count(strings.Join(faces, ""), faces[4]) == 2 || (strings.Count(strings.Join(faces, ""), faces[0]) == 2) && strings.Count(strings.Join(faces, ""), faces[4]) == 3 {
// 		return true
// 	}
// 	return false
// }
// func threeKind(faces []string) bool {
// 	if (strings.Count(strings.Join(faces, ""), faces[0]) == 3) && strings.Count(strings.Join(faces, ""), faces[4]) != 2 || (strings.Count(strings.Join(faces, ""), faces[0]) != 2) && strings.Count(strings.Join(faces, ""), faces[4]) == 3 {
// 		return true
// 	}
// 	return false
// }
// func twoPair(faces []string) bool {
// 	if (strings.Count(strings.Join(faces, ""), faces[1]) == 2) && strings.Count(strings.Join(faces, ""), faces[3]) == 2 {
// 		return true
// 	}
// 	return false
// }
// func onePair(faces []string) bool {
// 	for _, e := range faces {
// 		if strings.Count(strings.Join(faces, ""), e) == 2 {
// 			return true
// 		}
// 	}
// 	return false
// }

// type Hand struct {
// 	Bid, Strength int
// }
