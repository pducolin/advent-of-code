package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputData string

const (
	ROCK    = "R"
	PAPER   = "P"
	SCISSOR = "S"
)

var player1ToRPS = map[string]string{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSOR,
}

var player2ToRPS = map[string]string{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSOR,
}

var shapeScore = map[string]int{
	ROCK:    1,
	PAPER:   2,
	SCISSOR: 3,
}

var toWin = map[string]string{
	ROCK:    PAPER,
	PAPER:   SCISSOR,
	SCISSOR: ROCK,
}

var toLose = map[string]string{
	ROCK:    SCISSOR,
	PAPER:   ROCK,
	SCISSOR: PAPER,
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	fmt.Println("Running part", part)

	if part == 1 {
		fmt.Println(part1(inputData))
	} else {
		fmt.Println(part2(inputData))
	}
}

func part1(data string) string {
	totalScore := 0
	for _, line := range strings.Split(data, "\n") {
		players := strings.Split(line, " ")
		player1RPS := player1ToRPS[players[0]]
		player2RPS := player2ToRPS[players[1]]
		totalScore += evaluateRoundScore(player1RPS, player2RPS)
	}
	return fmt.Sprint(totalScore)
}

func part2(data string) string {
	totalScore := 0
	for _, line := range strings.Split(data, "\n") {
		players := strings.Split(line, " ")
		player1RPS := player1ToRPS[players[0]]
		player2RPS := player1RPS
		if players[1] == "X" {
			// need to lose
			player2RPS = toLose[player1RPS]
		} else if players[1] == "Y" {
			// need to draw, do nothing
		} else {
			// need to win
			player2RPS = toWin[player1RPS]
		}
		totalScore += evaluateRoundScore(player1RPS, player2RPS)
	}
	return fmt.Sprint(totalScore)
}

func evaluateRoundScore(player1RPS string, player2RPS string) int {
	score := shapeScore[player2RPS]

	if player1RPS == player2RPS {
		return score + 3
	}

	if (player1RPS == ROCK && player2RPS == SCISSOR) || (player1RPS == SCISSOR && player2RPS == PAPER) || (player1RPS == PAPER && player2RPS == ROCK) {
		return score
	}

	return score + 6
}
