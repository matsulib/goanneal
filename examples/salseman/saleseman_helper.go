// +build ignore

package main

import (
	"math"
	"math/rand"
)

// shuffle a Slice
func shuffle(data []string) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

// Convert degrees to radians
func radians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

// Calculates distance between two latitude-longitude coordinates.
func distance(a [2]float64, b [2]float64) float64 {
	R := 3963.0 // radius of Earth (miles)
	lat1, lon1 := radians(a[0]), radians(a[1])
	lat2, lon2 := radians(b[0]), radians(b[1])
	return math.Acos(math.Sin(lat1)*math.Sin(lat2)+
		math.Cos(lat1)*math.Cos(lat2)*math.Cos(lon1-lon2)) * R
}

type tuple [2]float64
type distanceMap map[string]float64

type america struct {
	cities         map[string]tuple
	distanceMatrix map[string]distanceMap
}

func newAmerica() *america {
	a := new(america)
	a.cities = map[string]tuple{
		"New York City": {40.72, 74.00},
		"Los Angeles":   {34.05, 118.25},
		"Chicago":       {41.88, 87.63},
		"Houston":       {29.77, 95.38},
		"Phoenix":       {33.45, 112.07},
		"Philadelphia":  {39.95, 75.17},
		"San Antonio":   {29.53, 98.47},
		"Dallas":        {32.78, 96.80},
		"San Diego":     {32.78, 117.15},
		"San Jose":      {37.30, 121.87},
		"Detroit":       {42.33, 83.05},
		"San Francisco": {37.78, 122.42},
		"Jacksonville":  {30.32, 81.70},
		"Indianapolis":  {39.78, 86.15},
		"Austin":        {30.27, 97.77},
		"Columbus":      {39.98, 82.98},
		"Fort Worth":    {32.75, 97.33},
		"Charlotte":     {35.23, 80.85},
		"Memphis":       {35.12, 89.97},
		"Baltimore":     {39.28, 76.62}}

	// create a distance matrix
	a.distanceMatrix = map[string]distanceMap{}
	for ka, va := range a.cities {
		a.distanceMatrix[ka] = distanceMap{}
		for kb, vb := range a.cities {
			if kb == ka {
				a.distanceMatrix[ka][kb] = 0.0
			} else {
				a.distanceMatrix[ka][kb] = distance(va, vb)
			}
		}
	}
	return a
}

func (a *america) CitiesKeys() []string {
	ks := []string{}
	for k := range a.cities {
		ks = append(ks, k)
	}
	return ks
}
