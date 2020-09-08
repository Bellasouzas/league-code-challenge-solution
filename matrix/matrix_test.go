package matrix

import (
	"testing"
)

func TestReadCsv(t *testing.T) {

}

func TestMatrixInt(t *testing.T) {

}

func TestMatrixIsEmpty(t *testing.T) {

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
