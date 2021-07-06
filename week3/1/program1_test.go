package main

import (
	//"strconv"
	"testing"
)

func TestProgram01_1(t *testing.T) {
	input := 1990
	expected := "MCMXC"
	result := IntToRoman(input)
	if result != expected {
		t.Error("Test01 is failed, get roman number = " + result)
	}
}

func TestProgram01_2(t *testing.T) {
	input := 2008
	expected := "MMVIII"
	result := IntToRoman(input)
	if result != expected {
		t.Error("Test02 is failed, get roman number = " + result)
	}
}

func TestProgram01_3(t *testing.T) {
	input := 99
	expected := "XCIX"
	result := IntToRoman(input)
	if result != expected {
		t.Error("Test03 is failed, get roman number = " + result)
	}
}

func TestProgram01_4(t *testing.T) {
	input := 47
	expected := "XLVII"
	result := IntToRoman(input)
	if result != expected {
		t.Error("Test04 is failed, get roman number = " + result)
	}
}
