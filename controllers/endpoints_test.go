package controllers_test

import (
	"github.com/bellasouzas/league-code-challenge-solution/controllers"
	"testing"
)

// This test verify if a square matrix as a string
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

//func TestInvert(t *testing.T) {
//	records := [][]string {
//		{"1","2","3"},
//		{"4","5","6"},
//		{"7","8","9"},
//	}
//	got := "1,2,3\n4,5,6\n7,8,9"
//	want :=  controllers.Invert(records)
//	if got != want {
//		t.Errorf("The result is not correct, got %v, expected %v", got, want)
//	}
//}
