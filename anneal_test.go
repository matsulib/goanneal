package goanneal

import (
	"math/rand"
	"testing"
)

var distanceMatrix map[string]distanceMap

type TravelState struct {
	state []string
}

// Returns an address of an exact copy of the receiver's state
func (ts *TravelState) Copy() interface{} {
	copiedState := make([]string, len(ts.state))
	copy(copiedState, ts.state)
	return &TravelState{state: copiedState}
}

// Swaps two cities in the route.
func (ts *TravelState) Move() {
	a := rand.Intn(len(ts.state))
	b := rand.Intn(len(ts.state))
	ts.state[a], ts.state[b] = ts.state[b], ts.state[a]
}

// Calculates the length of the route.
func (ts *TravelState) Energy() float64 {
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

func TestAnneal(t *testing.T) {
	problem := newAmerica()
	distanceMatrix = problem.DistanceMatrix
	// initial state, a randomly-ordered itinerary
	initialState := &TravelState{state: problem.CitiesKeys()}
	shuffle(initialState.state)

	tsp := NewAnnealer(initialState)
	tsp.Steps = 50000

	state, _ := tsp.Anneal()
	ts := state.(*TravelState)
	actual := len(ts.state)
	expected := 20
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestAuto(t *testing.T) {
	problem := newAmerica()
	distanceMatrix = problem.DistanceMatrix
	// initial state, a randomly-ordered itinerary
	initialState := &TravelState{state: problem.CitiesKeys()}
	shuffle(initialState.state)

	tsp := NewAnnealer(initialState)

	autoSchedule := tsp.Auto(0.05, 2000)
	tsp.SetSchedule(autoSchedule)

	if tsp.Tmax != autoSchedule["tmax"] {
		t.Errorf("got %v\nwant %v", tsp.Tmax, autoSchedule["tmax"])
	}

	if tsp.Tmin != autoSchedule["tmin"] {
		t.Errorf("got %v\nwant %v", tsp.Tmin, autoSchedule["tmin"])
	}

	if tsp.Steps != int(autoSchedule["steps"]) {
		t.Errorf("got %v\nwant %v", tsp.Steps, autoSchedule["steps"])
	}

	if tsp.Updates != int(autoSchedule["updates"]) {
		t.Errorf("got %v\nwant %v", tsp.Updates, autoSchedule["updates"])
	}
}
