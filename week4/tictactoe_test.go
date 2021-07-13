package main

import (
	"testing"
)

func TestShowBoard(t *testing.T) {
	input := [3][3]int{
		{1, 2, 1},
		{0, 2, 1},
		{2, 0, 1},
	}
	expected := "oxo\n.xo\nx.o\n"
	result := ShowBoard(input)
	if result != expected {
		t.Errorf("The expected board should be:\n%s, but is:\n%s", expected, result)
	}
}

func TestInputOX(t *testing.T) {
	input := "0,1"
	expected := [2]int{0, 1}
	result := InputOX(input)
	if result != expected {
<<<<<<< HEAD
		t.Errorf("The expected position should be:\nexpected, but is:\nresult")
=======
		t.Errorf("The expected position should be:\n%v, but is:\n%v", expected, result)
>>>>>>> dev
	}
}

func TestPutOX(t *testing.T) {
	input_0 := [3][3]int{
		{1, 0, 1},
		{0, 2, 0},
		{2, 0, 1},
	}
	input_1 := 2
	input_2 := [2]int{0, 1}

	expected_0 := [3][3]int{
		{1, 2, 1},
		{0, 2, 0},
		{2, 0, 1},
	}
	expected_1 := "Done!"
	result_0, result_1 := PutOX(input_0, input_1, input_2)

	if result_0 != expected_0 {
		t.Error("The expected board should be:\n", expected_0)
		t.Error(", but is:\n", result_0)
	}
	if result_1 != expected_1 {
		t.Errorf("The expected message should be \"%s\",\nbut is \"%s\"", expected_1, result_1)
	}
}

// func TestST1(t *testing.T) {
// 	expected :=
// 	result :=
// 	if result != expected {
// 		t.Error()
// 	}
// }

<<<<<<< HEAD
=======
func TestJudgeWin(t *testing.T) {
	input := [3][3]int{
		{1, 2, 1},
		{0, 2, 1},
		{2, 0, 1},
	}
	expected := 1
	result := JudgeWin(input)
	if result != expected {
		t.Errorf("The expected situation should be player %d win , but winner is %d", expected, result)
	}
}

>>>>>>> dev
// func TestST2(t *testing.T) {
// 	expected :=
// 	result :=
// 	if result != expected {
// 		t.Error()
// 	}
// }

// func TestUS1(t *testing.T) {
// 	expected :=
// 	result :=
// 	if result != expected {
// 		t.Error()
// 	}
// }
