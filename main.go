package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type GameMove int

const (
	Rock GameMove = iota
	Paper
	Scissor
)

var inputToGameMoveMap = map[string]GameMove{
	"A": Rock,
	"B": Paper,
	"C": Scissor,
	"X": Rock,
	"Y": Paper,
	"Z": Scissor,
}

var GameMoveScoreMap = map[GameMove]int{
	Rock:    1,
	Paper:   2,
	Scissor: 3,
}

const (
	Draw = 3
	Win  = 6
	Loss = 0
)

func getGameMovesForBothPlayers(reader io.Reader) ([]GameMove, []GameMove) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	playerOneMoves := make([]GameMove, 0)
	playerTwoMoves := make([]GameMove, 0)

	for scanner.Scan() {
		gameMoves := strings.Fields(scanner.Text())

		playerOneMoves = append(playerOneMoves, inputToGameMoveMap[gameMoves[0]])
		playerTwoMoves = append(playerTwoMoves, inputToGameMoveMap[gameMoves[1]])
	}

	return playerOneMoves, playerTwoMoves
}

func getOutcomeOfRound(playerOneMove GameMove, playerTwoMove GameMove) int {
	if playerOneMove == playerTwoMove {
		return Draw
	}

	if playerOneMove == Scissor && playerTwoMove == Paper ||
		playerOneMove == Paper && playerTwoMove == Rock ||
		playerOneMove == Rock && playerTwoMove == Scissor {
		return Loss
	}

	return Win
}

func main() {
	file, err := os.Open("/home/ec2-user/go/src/github.com/iamwillzhu/adventofcode2022day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	score := 0

	playerOneMoves, playerTwoMoves := getGameMovesForBothPlayers(bufio.NewReader(file))

	for index := range playerOneMoves {
		playerOneMove := playerOneMoves[index]
		playerTwoMove := playerTwoMoves[index]

		score += GameMoveScoreMap[playerTwoMove] + getOutcomeOfRound(playerOneMove, playerTwoMove)
	}

	fmt.Printf("The score of the strategy guide is %d\n", score)

}
