package goutils

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestIntInSlice(t *testing.T) {
	a := 1
	arr := []int{2, 1}
	ok := Int.InSlice(a, arr)
	if !ok {
		t.Errorf("IntInSlice() = %v, want %v", ok, true)
	}
}

func TestStringInSlice(t *testing.T) {
	a := "foo"
	arr := []string{"bar", "foo"}
	ok := Str.InSlice(a, arr)
	if !ok {
		t.Errorf("IntInSlice() = %v, want %v", ok, true)
	}
}