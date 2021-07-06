package main

import (
	"strconv"
	"testing"
)

func TestProgram02_1(t *testing.T) {
	input := "MCMXC"
	expected := 1990
	result := romanToInt(input)
	if result != expected {
		t.Error("Test01 is failed, get int number = " + strconv.Itoa(result))
	}
}

func TestProgram02_2(t *testing.T) {
	input := "MMVIII"
	expected := 2008
	result := romanToInt(input)
	if result != expected {
		t.Error("Test02 is failed, get int number = " + strconv.Itoa(result))
	}
}

func TestProgram02_3(t *testing.T) {
	input := "XCIX"
	expected := 99
	result := romanToInt(input)
	if result != expected {
		t.Error("Test03 is failed, get int number = " + strconv.Itoa(result))
	}
}

func TestProgram02_4(t *testing.T) {
	input := "XLVII"
	expected := 47
	result := romanToInt(input)
	if result != expected {
		t.Error("Test04 is failed, get int number = " + strconv.Itoa(result))
	}
}
