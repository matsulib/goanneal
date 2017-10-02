package goanneal

import (
	"regexp"
	"strconv"
	"testing"
)

func TestDivMod(t *testing.T) {
	a, b := divMod(10, 3)
	expectedA, expectedB := 3, 1
	if a != expectedA || b != expectedB {
		t.Errorf("got %v, %v\nwant %v, %v", a, b, expectedA, expectedB)
	}
}

func TestRoundFigure(t *testing.T) {
	a := roundFigure(1111.1111, 3)
	expectedA := 1110.0
	if a != expectedA {
		t.Errorf("got %v\nwant %v", a, expectedA)
	}

	b := roundFigure(1.1111, 3)
	expectedB := 1.11
	if b != expectedB {
		t.Errorf("got %v\nwant %v", b, expectedB)
	}
}

func TestTimeString(t *testing.T) {
	seconds := 2*3600 + 3*60 + 4.5
	a := timeString(seconds)
	expectedA := "   2:03:05"
	if a != expectedA {
		t.Errorf("got %v\nwant %v", a, expectedA)
	}
}

func TestNow(t *testing.T) {
	a := strconv.FormatFloat(now(), 'f', 6, 64)
	expectedA := "1506180653.476423"
	if !checkRegexp(`^\d{10}\.\d{5,}`, a) {
		t.Errorf("got %v\nwant a string like %v", a, expectedA)
	}
}

func checkRegexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
