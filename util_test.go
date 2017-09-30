package goanneal

import (
	"regexp"
	"strconv"
	"testing"
)

func TestDivMod(t *testing.T) {
	a, b := divMod(10, 3)
	expected_a, expected_b := 3, 1
	if a != expected_a || b != expected_b {
		t.Errorf("got %v, %v\nwant %v, %v", a, b, expected_a, expected_b)
	}
}

func TestRoundFigure(t *testing.T) {
	a := roundFigure(1111.1111, 3)
	expected_a := 1110.0
	if a != expected_a {
		t.Errorf("got %v\nwant %v", a, expected_a)
	}

	b := roundFigure(1.1111, 3)
	expected_b := 1.11
	if b != expected_b {
		t.Errorf("got %v\nwant %v", b, expected_b)
	}
}

func TestTimeString(t *testing.T) {
	seconds := 2*3600 + 3*60 + 4.5
	a := timeString(seconds)
	expected_a := "   2:03:05"
	if a != expected_a {
		t.Errorf("got %v\nwant %v", a, expected_a)
	}
}

func TestNow(t *testing.T) {
	a := strconv.FormatFloat(now(), 'f', 6, 64)
	expected_a := "1506180653.476423"
	if !check_regexp(`^\d{10}\.\d{5,}`, a) {
		t.Errorf("got %v\nwant a string like %v", a, expected_a)
	}
}

func check_regexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
