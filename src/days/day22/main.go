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

type Round struct {
	Player1Value int
	Player2Value int
}

func part2(player1Deck []int, player2Deck []int) int {
	_, res := PlayGame(player1Deck, player2Deck)
	return res
}

var gameMem = make(map[Round][]int)

func PlayGame(player1Deck []int, player2Deck []int) (int, int) {
	v1 := Value(player1Deck)
	v2 := Value(player2Deck)
	r := Round{Player1Value: v1, Player2Value: v2}
	if v, ok := gameMem[r]; ok {
		return v[0], v[1]
	}

	newPlayer1Deck := make([]int, len(player1Deck))
	copy(newPlayer1Deck, player1Deck)
	newPlayer2Deck := make([]int, len(player2Deck))
	copy(newPlayer2Deck, player2Deck)
	mem := make(map[Round]bool)

	var winner int
	var gameEnded bool
	for {
		winner, gameEnded, newPlayer1Deck, newPlayer2Deck = playRound(newPlayer1Deck, newPlayer2Deck, mem)
		if gameEnded {
			winnerDeck := newPlayer1Deck
			if winner == 2 {
				winnerDeck = newPlayer2Deck
			}
			v := Value(winnerDeck)
			gameMem[r] = []int{winner, v}
			return winner, v
		}
	}
}

func playRound(player1Deck []int, player2Deck []int, mem map[Round]bool) (int, bool, []int, []int) {
	if len(player1Deck) == 0 || len(player2Deck) == 0 {
		winner := 1
		if len(player1Deck) == 0 {
			winner = 2
		}
		return winner, true, player1Deck, player2Deck
	}
	// Condition 1: If the round has already been played
	// Player1 wins
	v1 := Value(player1Deck)
	v2 := Value(player2Deck)
	r := Round{Player1Value: v1, Player2Value: v2}

	if _, ok := mem[r]; ok {
		return 1, true, player1Deck, player2Deck
	}
	mem[r] = true

	p1 := player1Deck[0]
	p2 := player2Deck[0]
	player1Deck = player1Deck[1:]
	player2Deck = player2Deck[1:]

	// Condition 2: The winner is determined by playing a new GAME
	if len(player1Deck) >= p1 && len(player2Deck) >= p2 {
		winner, _ := PlayGame(player1Deck[:p1], player2Deck[:p2])
		if winner == 1 {
			player1Deck = append(player1Deck, p1, p2)
			return 1, false, player1Deck, player2Deck
		}
		player2Deck = append(player2Deck, p2, p1)
		return 2, false, player1Deck, player2Deck
	}

	// Condition 3: Highest card wins
	if p1 > p2 {
		player1Deck = append(player1Deck, p1, p2)
		return 1, false, player1Deck, player2Deck
	}
	player2Deck = append(player2Deck, p2, p1)
	return 2, false, player1Deck, player2Deck
}

func Value(deck []int) int {
	res := 0
	for i, v := range deck {
		res += (len(deck) - i) * v
	}
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
