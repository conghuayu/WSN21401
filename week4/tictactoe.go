package main

import (
	// 	"fmt"
	"strconv"
	"strings"
)

func ShowBoard(board [3][3]int) string {
	//0:. 1:o 2:x
	str := ""
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch board[i][j] {
			case 0:
				str += "."
			case 1:
				str += "o"
			case 2:
				str += "x"
			default:
				str = "Error!"
				return str
			}
		}
		str += "\n"
	}
	return str
}

func InputOX(position string) [2]int {
	// 	fmt.Printf("Player %d: Input (x,y) ", player)

	tmp := strings.Replace(position, " ", "", -1)
	positions := strings.Split(tmp, ",")

	i, _ := strconv.Atoi(positions[0])
	j, _ := strconv.Atoi(positions[1])

	positions_num := [2]int{i, j}
	return positions_num
}

func PutOX(board [3][3]int, player int, position [2]int) ([3][3]int, string) {
	i := position[0]
	j := position[1]

	if board[i][j] != 0 {
		return board, "Can't put in this position!"
	}

	board[i][j] = player
	return board, "Done!"
}
