package board

import "fmt"

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}
func ClearBoard(board *[3][3]byte) {
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			board[i][k] = byte(0)
		}
	}
}
