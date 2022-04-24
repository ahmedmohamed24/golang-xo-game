package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	xoBoard "github.com/ahmedmohamed24/golang-xo-game/board"
	"github.com/alash3al/go-color"
)

var board = [3][3]byte{
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
}

type player struct {
	name  string
	mark  byte
	score int
}

func main() {
	xoBoard.BoardDisplay(&board)
	xoBoard.ClearTerminal()
	var name string
	color.Green("%s\n", "First player name:")
	takeUserName(&name)
	firstPlayer := player{name: name, score: 0, mark: 'x'}
	color.Blue("%s\n", "Second player name:")
	takeUserName(&name)
	secondPlayer := player{name: name, score: 0, mark: 'y'}
	playRound(&firstPlayer, &secondPlayer)
}
func (p *player) play() {
	color.Magenta("%vScore:%v\n", p.name, p.score)
	xoBoard.BoardDisplay(&board)
	var row, column int
	for {
		color.Yellow("%v\n", "Enter the row number 0,1,2")
		_, err := fmt.Scanf("%d", &row)
		if err == nil && row <= 2 {
			break
		}
		color.Red("Invalid row value!")
	}

	for {
		color.Yellow("%v\n", "Enter the column number 0,1,2")
		_, err := fmt.Scanf("%d", &column)
		if err == nil && column <= 2 {
			break
		}
		color.Red("Invalid column value!")
	}
	if board[column][row] == byte(0) {
		board[column][row] = p.mark
	} else {
		xoBoard.ClearTerminal()
		color.Red("%v\n", "Please select an empty cell!")
		p.play()
	}

}
func (p *player) checkWinner() bool {

	//check similarity in rows
	for i := 0; i < 3; i++ {
		if (board[i][0] == board[i][1]) && (board[i][1] == board[i][2]) && (board[i][1] == byte(p.mark)) {
			p.score += 1
			return true
		}
	}
	//check similarity in columns
	for i := 0; i < 3; i++ {
		if (board[0][i] == board[1][i]) && (board[1][i] == board[2][i]) && (board[0][i] == byte(p.mark)) {
			p.score += 1
			return true
		}
	}
	//check similarity in diagonals
	if (board[0][0] == board[1][1]) && (board[1][1] == board[2][2]) && (board[0][0] == byte(p.mark)) {
		p.score += 1
		return true
	}
	if (board[0][2] == board[1][1]) && (board[1][1] == board[2][0]) && (board[0][2] == byte(p.mark)) {
		p.score += 1
		return true
	}
	return false
}

func playRound(firstPlayer *player, secondPlayer *player) {
	xoBoard.ClearBoard(&board)
	for i := 0; ; i++ {
		xoBoard.ClearTerminal()
		if i%2 == 0 {
			firstPlayer.play()
			winner := firstPlayer.checkWinner()
			if winner {
				xoBoard.BoardDisplay(&board)
				color.Cyan("%s\n", "Winner")
				time.Sleep(time.Second * 4)
				xoBoard.ClearBoard(&board)
				playRound(firstPlayer, secondPlayer)

			}
		} else {
			secondPlayer.play()
			winner := secondPlayer.checkWinner()
			if winner {
				xoBoard.BoardDisplay(&board)
				color.Cyan("%s\n", "Winner")
				time.Sleep(time.Second * 4)
				xoBoard.ClearBoard(&board)
				playRound(firstPlayer, secondPlayer)
			}

		}
		if checkFinished() {
			xoBoard.BoardDisplay(&board)
			color.Cyan("%s\n", "Finished")
			time.Sleep(time.Second * 4)
			xoBoard.ClearBoard(&board)
			playRound(firstPlayer, secondPlayer)

		}
	}
}
func checkFinished() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == byte(0) {
				return false
			}

		}
	}
	return true

}

func takeUserName(name *string) {
	for i := 0; ; i++ {
		terminalReader := bufio.NewReader(os.Stdout)
		val, err := terminalReader.ReadString('\n')
		if err == nil {
			*name = string(val)
			return
		}
		color.Red("%s %s", err.Error(), "Please enter a valid username!")
	}
}
