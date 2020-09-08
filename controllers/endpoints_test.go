package controllers_test

import (
	"github.com/bellasouzas/league-code-challenge-solution/controllers"
	"testing"
)

func TestEcho(t *testing.T) {
	records := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	got := "1,2,3\n4,5,6\n7,8,9\n"
	want := controllers.Echo(records)
	if got != want {
		t.Errorf("The result is not correct, got %v, expected %v", got, want)
	}
}

func TestInvert(t *testing.T) {
	records := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	got := "1,4,7\n2,5,8\n3,6,9\n"
	want := controllers.Invert(records)
	if got != want {
		t.Errorf("The result is not correct, got %v, expected %v", got, want)
	}
}

func TestFlatten(t *testing.T) {
	records := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	got := "1,2,3,4,5,6,7,8,9"
	want := controllers.Flatten(records)
	if got != want {
		t.Errorf("The result is not correct, got %v, expected %v", got, want)
	}
}

func TestSum(t *testing.T) {
	records := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	got := 45
	want := controllers.Sum(records)
	if got != want {
		t.Errorf("The result is not correct, got %d, expected %d", got, want)
	}
}
