package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	a := 5
	b := 5
	want := 10
	result := sum(a, b)

	if want != result {
		t.Fail()
	}
}

func TestSumZeroValues(t *testing.T) {
	a := 0
	b := 0
	want := 0
	result := sum(a, b)

	if want != result {
		t.Fail()
	}
}

func TestSumNegativeValues(t *testing.T) {
	a := -5
	b := -5
	want := -10
	result := sum(a, b)

	if want != result {
		t.Fail()
	}
}
