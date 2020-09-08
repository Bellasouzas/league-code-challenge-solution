package matrix

import (
	"github.com/bellasouzas/league-code-challenge-solution/controllers"
	"testing"
)

func TestReadCsv(t *testing.T) {

}

func TestMatrixInt(t *testing.T) {
	records := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	got := [][]int{1,2,3\n4,5,6,7,8,9}
	want := MatrixInt(records)
	if got != want {
		t.Errorf("The result is not correct, got %d, expected %d", got, want)
	}
}

func TestMatrixIsEmpty(t *testing.T) {
	squareEmpty := [][]int {{0},{0}}
	got := true
	want := MatrixIsEmpty(squareEmpty)
	if got != want {
		t.Errorf("The result is not correct, got %v, expected %v", got, want)
	}

}

func TestMatrixSquare(t *testing.T) {
	squareMatrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	got := true
	want := MatrixSquare(squareMatrix)
	if got != want {
		t.Errorf("The result is not correct, got %v, expected %v", got, want)
	}
}
