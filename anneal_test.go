package goanneal

import (
	"math/rand"
	"testing"
)

var distanceMatrix map[string]distanceMap

type TravelState struct {
	state         []string
	previousState []string
	bestState     []string
}

// Save an exact copy of the current state to the previous state
func (ts *TravelState) Backup() {
	copy(ts.previousState, ts.state)
}

// Restore an the current state to the previous state
func (ts *TravelState) Restore() {
	copy(ts.state, ts.previousState)
}

// Save an exact copy of the current state to the best state
func (ts *TravelState) RecordBest() {
	copy(ts.bestState, ts.state)
}

// Restore an the current state to the best state
func (ts *TravelState) RememberBest() {
	copy(ts.state, ts.bestState)
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

	tsp.Anneal()
	ts := tsp.State.(*TravelState)
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
	initialState := &TravelState{
		state: problem.CitiesKeys(),
		previousState: make([]string, len(problem.CitiesKeys())),
		bestState:     make([]string, len(problem.CitiesKeys()))}
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
