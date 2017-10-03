// +build ignore

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	landmarks      map[string]tuple
	distanceMatrix map[string]distanceMap
}

func newAmerica() *america {
	a := new(america)

	// set landmarks
	bl, err := ioutil.ReadFile("./examples/salseman50/landmarks.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(bl, &a.landmarks); err != nil {
		log.Fatal(err)
	}

	// set distance_matrix
	bd, err := ioutil.ReadFile("./examples/salseman50/distance_matrix.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(bd, &a.distanceMatrix); err != nil {
		log.Fatal(err)
	}

	for ci := range a.landmarks {
		for cj := range a.landmarks {
			if val, ok := a.distanceMatrix[ci][cj]; ok {
				// pass
			} else {
				if ci == cj {
					a.distanceMatrix[ci][cj] = 0
				}
				val = a.distanceMatrix[cj][ci] / 1000.0
				a.distanceMatrix[ci][cj] = val
				a.distanceMatrix[cj][ci] = val
			}
		}
	}

	return a
}

func (a *america) landmarksKeys() []string {
	ks := []string{}
	for k := range a.landmarks {
		ks = append(ks, k)
	}
	return ks
}
