package routes

import (
	_ "github.com/stretchr/testify"
	"testing"
)

records := [][]string {
	{"1","2","3"},
	{"4","5","6"},
	{"7","8","9"},
}

//todo escver Echo test
func TestHandleEcho(t *testing.T) {
	expected := Echo(records)

}

//TODO Escrever Invert Test
func TestHandleInvert(t *testing.T) {

}

//TODO Escrever HandleFlatten Test
func TestHandleFlatten(t *testing.T) {

}

//TODO Escrever Sum test
func TestHandleSum(t *testing.T) {

}

//todo escrever HandleMultiply test
func TestHandleMultiply(t *testing.T) {

}
