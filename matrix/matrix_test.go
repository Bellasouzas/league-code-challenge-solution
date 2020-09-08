package matrix

import (
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
	got := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	want := MatrixInt(records)
	if got == nil {
		t.Errorf("The result is not correct, got %v, expected %v", got, want)
	}
}

func TestMatrixIsEmpty(t *testing.T) {
	matrixEmpty := [][]string{}
	got := false
	want := MatrixIsNotEmpty(matrixEmpty)
	if got != want {
		t.Errorf("The result is not correct, got %t, expected %t", got, want)
	}

}

func TestMatrixSquare(t *testing.T) {
	matrixNotSquare := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}
	got := true
	want := MatrixIsNotEmpty(matrixNotSquare)
	if got != want {
		t.Errorf("The result is not correct, got %v, expected %v", got, want)
	}
}
