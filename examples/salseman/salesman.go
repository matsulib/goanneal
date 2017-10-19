// +build ignore

package main

import (
	"fmt"
	"math/rand"

	"github.com/matsulib/goanneal"
)

var distanceMatrix map[string]distanceMap

type travelState struct {
	state         []string
	previousState []string
	bestState     []string
}

// Save an exact copy of the current state to the previous state
func (ts *travelState) Backup() {
	copy(ts.previousState, ts.state)
}

// Restore an the current state to the previous state
func (ts *travelState) Restore() {
	copy(ts.state, ts.previousState)
}

// Save an exact copy of the current state to the best state
func (ts *travelState) RecordBest() {
	copy(ts.bestState, ts.state)
}

// Restore an the current state to the best state
func (ts *travelState) RememberBest() {
	copy(ts.state, ts.bestState)
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
	initialState := &travelState{state: problem.CitiesKeys()}
	shuffle(initialState.state)

	tsp := goanneal.NewAnnealer(initialState)
	tsp.Steps = 100000

	tsp.Anneal()
	ts := tsp.State.(*travelState)
	for ts.state[0] != "New York City" {
		ts.state = append(ts.state[1:], ts.state[:1]...)
	}

	fmt.Println(int(ts.Energy()), "mile route:")
	for i := 0; i < len(ts.state); i++ {
		fmt.Println("\t", ts.state[i])
	}
}
