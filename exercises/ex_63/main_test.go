package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println("Testing Add function.")
	total := Add(3, 2)
	if total != 5 {
		t.Errorf("Expected 5, actual %d", total)
	}
}

func TestSub(t *testing.T) {
	fmt.Println("Testing substract function.")
	total := substract(3, 2)
	if total != 1 {
		t.Errorf("Expected 1, actual %d", total)
	} else {
		t.Log("Success!")
	}
}

func TestdoMathAdd(t *testing.T) {
	fmt.Println("Testing doMath Add function.")
	total := doMath(3, 2, Add)
	if total != Add(3, 2) {
		t.Errorf("Expected 5, actual %d", total)
	} else {
		t.Log("Success!")
	}
}

func TestdoMathSub(t *testing.T) {
	fmt.Println("Testing Add function.")
	total := doMath(3, 2, substract)
	if total != substract(3, 2) {
		t.Errorf("Expected 1, actual %d", total)
	} else {
		t.Log("Success!")
	}
}
