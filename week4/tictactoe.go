package main

import (
	"fmt"
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

func PutOX(board [3][3]int, player int, positions_num [2]int) ([3][3]int, string) {
	i := positions_num[0]
	j := positions_num[1]

	if board[i][j] != 0 {
		return board, "Can't put on this position!"
	}

	board[i][j] = player
	return board, "Done!"
}

func tri_equals(i int, j int, k int) bool {
	if i == j && i == k && j == k {
		return true
	} else {
		return false
	}
}

func JudgeWin(board [3][3]int) int {
	//row
	for i := 0; i < 3; i++ {
		if tri_equals(board[i][0], board[i][1], board[i][2]) {
			return board[i][0]
		}
		if tri_equals(board[0][i], board[1][i], board[2][i]) {
			return board[0][i]
		}
	}
	if tri_equals(board[0][0], board[1][1], board[2][2]) || tri_equals(board[2][0], board[1][1], board[0][2]) {
		return board[1][1]
	}
	return 0
}

func main() {
	board := [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	for marks := 1; marks <= 9; {
		var player int
		var position string
		var message string
		if marks%2 == 1 {
			player = 1
		} else {
			player = 2
		}
		fmt.Printf("Player %d: Input (x,y) ", player)
		fmt.Scan(&position)

		positions_num := InputOX(position)
		board, message = PutOX(board, player, positions_num)

		if message != "Done!" {
			fmt.Println(message)
		} else {
			fmt.Print(ShowBoard(board))

			if JudgeWin(board) != 0 {
				fmt.Printf("Player %d won!", player)
				return
			}

			marks++
		}
	}
}
