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

type RoundOutcome int

const (
	Win RoundOutcome = iota
	Draw
	Loss
)

var inputToGameMoveMap = map[string]GameMove{
	"A": Rock,
	"B": Paper,
	"C": Scissor,
	"X": Rock,
	"Y": Paper,
	"Z": Scissor,
}

var inputToRoundOutcomeMap = map[string]RoundOutcome{
	"X": Loss,
	"Y": Draw,
	"Z": Win,
}

var GameMoveScoreMap = map[GameMove]int{
	Rock:    1,
	Paper:   2,
	Scissor: 3,
}

const (
	DrawOutcomeScore = 3
	WinOutcomeScore  = 6
	LossOutcomeScore = 0
)

func getPartTwoPlayerTwoGameMove(playerOneGameMove GameMove, playerTwoRoundOutcome RoundOutcome) GameMove {
	if playerTwoRoundOutcome == Draw {
		return playerOneGameMove
	}

	if playerTwoRoundOutcome == Win {
		if playerOneGameMove == Rock {
			return Paper
		}
		if playerOneGameMove == Paper {
			return Scissor
		}
		if playerOneGameMove == Scissor {
			return Rock
		}
	}

	if playerOneGameMove == Rock {
		return Scissor
	}
	if playerOneGameMove == Paper {
		return Rock
	}
	return Paper
}

func getGameMovesForBothPlayers(reader io.Reader) (playerOneGameMoves []GameMove, playerTwoPartOneGameMoves []GameMove, playerTwoPartTwoGameMoves []GameMove) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	playerOneGameMoves = make([]GameMove, 0)
	playerTwoPartOneGameMoves = make([]GameMove, 0)
	playerTwoPartTwoGameMoves = make([]GameMove, 0)

	for scanner.Scan() {
		gameInputs := strings.Fields(scanner.Text())

		playerOneGameMove := inputToGameMoveMap[gameInputs[0]]
		playerTwoPartOneGameMove := inputToGameMoveMap[gameInputs[1]]
		playerTwoRoundOutcome := inputToRoundOutcomeMap[gameInputs[1]]

		playerTwoPartTwoGameMove := getPartTwoPlayerTwoGameMove(playerOneGameMove, playerTwoRoundOutcome)

		playerOneGameMoves = append(playerOneGameMoves, playerOneGameMove)
		playerTwoPartOneGameMoves = append(playerTwoPartOneGameMoves, playerTwoPartOneGameMove)
		playerTwoPartTwoGameMoves = append(playerTwoPartTwoGameMoves, playerTwoPartTwoGameMove)
	}

	return
}

func getOutcomeOfRound(playerOneMove GameMove, playerTwoMove GameMove) int {
	if playerOneMove == playerTwoMove {
		return DrawOutcomeScore
	}

	if playerOneMove == Scissor && playerTwoMove == Paper ||
		playerOneMove == Paper && playerTwoMove == Rock ||
		playerOneMove == Rock && playerTwoMove == Scissor {
		return LossOutcomeScore
	}

	return WinOutcomeScore
}

func main() {
	file, err := os.Open("/home/ec2-user/go/src/github.com/iamwillzhu/adventofcode2022day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scorePartOne := 0
	scorePartTwo := 0

	playerOneGameMoves, playerTwoPartOneGameMoves, playerTwoPartTwoGameMoves := getGameMovesForBothPlayers(bufio.NewReader(file))

	for index := range playerOneGameMoves {
		playerOneGameMove := playerOneGameMoves[index]
		playerTwoPartOneGameMove := playerTwoPartOneGameMoves[index]
		playerTwoPartTwoGameMove := playerTwoPartTwoGameMoves[index]

		scorePartOne += GameMoveScoreMap[playerTwoPartOneGameMove] + getOutcomeOfRound(playerOneGameMove, playerTwoPartOneGameMove)
		scorePartTwo += GameMoveScoreMap[playerTwoPartTwoGameMove] + getOutcomeOfRound(playerOneGameMove, playerTwoPartTwoGameMove)
	}

	fmt.Printf("The score of the strategy guide for part 1 is %d\n", scorePartOne)
	fmt.Printf("The score of the strategy guide for part 2 is %d\n", scorePartTwo)
}
