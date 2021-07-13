package main

import (
	"strconv"
	"testing"
)

func TestCalRPN01(t *testing.T) { //Plus
	input := []string{"1", "2", "+"}
	expected := 3
	result := calRPN(input)
	if result != expected {
		t.Error("Test01 is failed, get calculator result = " + strconv.Itoa(result))
	}
}

func TestCalRPN02(t *testing.T) { //Minus
	input := []string{"2", "1", "-"}
	expected := 1
	result := calRPN(input)
	if result != expected {
		t.Error("Test01 is failed, get calculator result = " + strconv.Itoa(result))
	}
}

func TestCalRPN03(t *testing.T) { //Multiplication
	input := []string{"2", "3", "*"}
	expected := 6
	result := calRPN(input)
	if result != expected {
		t.Error("Test01 is failed, get calculator result = " + strconv.Itoa(result))
	}
}

func TestCalRPN04(t *testing.T) { //Division
	input := []string{"1", "/", "3", "1", "/", "6", "+"}
	expected := 1 / 2
	result := calRPN(input)
	if result != expected {
		t.Error("Test01 is failed, get calculator result = " + strconv.Itoa(result))
	}
}
