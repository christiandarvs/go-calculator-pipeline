package main

import "testing"

func TestAddNum(t *testing.T) {
	got := AddNum(2, 3)
	want := 5

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSubtractNum(t *testing.T) {
	got := SubtractNum(2, 3)
	want := -1

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMultiplyNum(t *testing.T) {
	got := MultiplyNum(4, 2)
	want := 8

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDivideNum(t *testing.T) {
	got := DivideNum(10, 2)
	want := 5

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
