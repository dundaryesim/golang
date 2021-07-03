// 1.Extend listing 16.8 to display all the chess pieces at their starting positions using
// the characters kqrbnp for black pieces along the top and uppercase KQRBNP for white pieces on the bottom.

// 2.Write a function that nicely displays the board.

// 3.Instead of strings, use [8][8]rune to represent the board. Recall that rune literals are
// surrounded with single quotes and can be printed with the %c format verb.

package main

import "fmt"

func main() {
	var board [8][8]rune
	// black pieces
	board[0][0] = 'r'
	board[0][1] = 'n'
	board[0][2] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'
	board[0][5] = 'b'
	board[0][6] = 'n'
	board[0][7] = 'r'

	// white pieces
	board[7][0] = 'R'
	board[7][1] = 'N'
	board[7][2] = 'B'
	board[7][3] = 'Q'
	board[7][4] = 'K'
	board[7][5] = 'B'
	board[7][6] = 'N'
	board[7][7] = 'R'

	//pawns
	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}
	fmt.Printf("%c", board)
}
