package main

import (
	"fmt"
	"math"
)

func IntToRoman(num int) string {
	if num > 4999 {
		return "Overflow!"
	}
	intToRomanMap := [4]map[int]string{
		map[int]string{1: "I", 5: "V"},
		map[int]string{1: "X", 5: "L"},
		map[int]string{1: "C", 5: "D"},
		map[int]string{1: "M"},
	}

	for i := 0; i < 4; i++ {
		for j := 2; j <= 9; j++ {
			tmp := ""
			one := intToRomanMap[i][1]
			five := intToRomanMap[i][5]
			ten := ""
			if i < 3 {
				ten = intToRomanMap[i+1][1]
			}

			switch j {
			//2=1+1 3=1+1+1 4=1+5   6=5+1 7=5+1+1 8=5+1+1+1 9=1+10
			case 0:
				tmp = ""
			case 2:
				tmp = one + one
			case 3:
				tmp = one + one + one
			case 4:
				tmp = one + five
			case 6:
				tmp = five + one
			case 7:
				tmp = five + one + one
			case 8:
				tmp = five + one + one + one
			case 9:
				tmp = one + ten
			default:
				continue
			}

			intToRomanMap[i][j] = tmp

			if i == 3 && j == 4 {
				intToRomanMap[i][j] = one + one + one + one
				break
			}

		}
	}

	roman := ""

	for i := 3; i >= 0; i-- {
		tens := int(math.Pow(10, float64(i)))
		digit := int(num / tens)

		roman += intToRomanMap[i][digit]

		num %= tens
	}

	return roman

}

func main() {
	fmt.Println(IntToRoman(5000))
}
