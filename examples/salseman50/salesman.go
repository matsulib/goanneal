// +build ignore

package main

import (
	"fmt"
	"math/rand"

	"github.com/matsulib/goanneal"
)

var distanceMatrix map[string]distanceMap

type travelState struct {
	state []string
}

// Returns an address of an exact copy of the receiver's state
func (ts *travelState) Copy() interface{} {
	copiedState := make([]string, len(ts.state))
	copy(copiedState, ts.state)
	return &travelState{state: copiedState}
}

// Swaps two cities in the route.
func (ts *travelState) Move() {
	a := rand.Intn(len(ts.state))
	b := rand.Intn(len(ts.state))
	ts.state[a], ts.state[b] = ts.state[b], ts.state[a]
}

// Calculates the length of the route.
func (ts *travelState) Energy() float64 {
	e := 0.0
	for i := 0; i < len(ts.state); i++ {
		if i == 0 {
			e += distanceMatrix[ts.state[len(ts.state)-1]][ts.state[0]]
		} else {
			e += distanceMatrix[ts.state[i-1]][ts.state[i]]
		}
	}
	return e
}

func main() {
	problem := newAmerica()
	distanceMatrix = problem.distanceMatrix
	// initial state, a randomly-ordered itinerary
	initialState := &travelState{state: problem.landmarksKeys()}
	shuffle(initialState.state)

	tsp := goanneal.NewAnnealer(initialState)
	tsp.Steps = 10000000
	tsp.Tmin = 1.0

	state, energy := tsp.Anneal()
	ts := state.(*travelState)
	for ts.state[0] != "Grand Canyon National Park, Arizona, USA" {
		ts.state = append(ts.state[1:], ts.state[:1]...)
	}

	fmt.Println(int(energy), "km route:")
	for i := 0; i < len(ts.state); i++ {
		fmt.Println("\t", ts.state[i])
	}
	//22493.016
}
