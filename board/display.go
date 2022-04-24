package board

import "fmt"

func BoardDisplay(board *[3][3]byte) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch board[j][i] {
			case 'x':
				fmt.Printf("%v\t", string(byte(board[j][i])))
			case 'y':
				fmt.Printf("%v\t", string(byte(board[j][i])))
			default:
				fmt.Printf("%x\t", board[j][i])
			}
		}
		fmt.Println()
	}

}
