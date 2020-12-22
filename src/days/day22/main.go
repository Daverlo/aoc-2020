package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//return map[ingredient] -> (map [alergen] -> foodIndex)
func parseInput(path string) ([]int, []int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	rd := bufio.NewReader(f)

	var player1Deck []int
	var player2Deck []int
	player1Cards := true
	for {
		var line string
		line, err = rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, nil, err
			}
		}
		line = strings.TrimSuffix(line, "\n")

		if line == "" {
			player1Cards = false
		}

		v, err := strconv.Atoi(line)
		if err == nil {
			if player1Cards {
				player1Deck = append(player1Deck, v)
			} else {
				player2Deck = append(player2Deck, v)
			}
		}
	}

	return player1Deck, player2Deck, nil
}

func part1(player1Deck []int, player2Deck []int) int {
	for len(player1Deck) > 0 && len(player2Deck) > 0 {
		// fmt.Println(player1Deck, player2Deck)
		p1 := player1Deck[0]
		p2 := player2Deck[0]
		player1Deck = player1Deck[1:]
		player2Deck = player2Deck[1:]

		if p1 > p2 {
			player1Deck = append(player1Deck, p1, p2)
		} else {
			player2Deck = append(player2Deck, p2, p1)
		}
	}
	// fmt.Println(player1Deck, player2Deck)

	winnerDeck := player1Deck
	if len(player1Deck) == 0 {
		winnerDeck = player2Deck
	}

	res := 0
	for i, v := range winnerDeck {
		// fmt.Println(len(winnerDeck)-i+1, v)
		res += (len(winnerDeck) - i) * v
	}

	return res
}

func part2(player1Deck []int, player2Deck []int) int {
	res := 0
	return res
}

func main() {
	args := os.Args[1:]
	player1Deck, player2Deck, err := parseInput(args[0])

	if err != nil {
		panic(err)
	}

	output := part1(player1Deck, player2Deck)
	fmt.Println(output)

	output = part2(player1Deck, player2Deck)
	fmt.Println(output)
}
