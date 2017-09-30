package goanneal

import (
	"fmt"
	"math"
	"time"
)

func divMod(x, y int) (int, int) {
	return x / y, x % y
}

// Returns x rounded to n significant figures.
func roundFigure(x float64, n int) float64 {
	return roundPlus(x, n-int(math.Ceil(math.Log10(math.Abs(x)))))
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func roundPlus(f float64, digits int) float64 {
	shift := math.Pow(10, float64(digits))
	return round(f*shift) / shift
}

// Returns time in seconds as a string formatted HHHH:MM:SS.
func timeString(seconds float64) string {
	s := int(round(seconds))
	h, s := divMod(s, 3600)
	m, s := divMod(s, 60)
	return fmt.Sprintf("%4d:%02d:%02d", h, m, s)
}

func now() float64 {
	return float64(time.Now().UnixNano()) / float64(math.Pow(10, 9))
}
