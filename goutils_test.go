package goutils

import (
	"os"
	"strings"
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

func TestStrSliceContains(t *testing.T) {
	a := []string{"bar", "foo"}
	b := "foo test"
	_, ok := Str.SliceContains(a, b)
	if !ok {
		t.Errorf("IntInSlice() = %v, want %v", ok, true)
	}
}

func TestSliceContains(t *testing.T) {
	a := []string{"bar", "foo"}
	b := "foo"
	_, ok := Slice.Contains(a, b)
	if !ok {
		t.Errorf("IntInSlice() = %v, want %v", ok, true)
	}
}

func TestSliceTrimSpace(t *testing.T) {
	arr := []string{"bar ", " foo", " tst "}
	Slice.TrimSpace(arr)
	for _, a := range arr {
		if len(a) != 3 {
			t.Errorf("TrimSpace() = %v, want %v", a, strings.TrimSpace(a))
		}
	}
}
