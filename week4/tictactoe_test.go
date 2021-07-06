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
		t.Error("The expected board should be:\nexpected, but is:\nresult")
	}
}

func TestInputOX(t *testing.T) {
	input_0 := 1
	input_1 := "0,1"
	expected := [2]int{0, 1}
	result := InputOX(input_0, input_1)
	if result != expected {
		t.Error("The expected position should be:\nexpected, but is:\nresult")
	}
}

// func TestST1(t *testing.T) {
// 	expected :=
// 	result :=
// 	if result != expected {
// 		t.Error()
// 	}
// }

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
