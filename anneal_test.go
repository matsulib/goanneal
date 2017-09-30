package goanneal

import (
	"math/rand"
	"testing"
)

var distanceMatrix map[string]DMap

type TravelState struct {
	state []string
}

// Returns an address of an exact copy of the receiver's state
func (self *TravelState) Copy() interface{} {
	copiedState := make([]string, len(self.state))
	copy(copiedState, self.state)
	return &TravelState{state: copiedState}
}

// Swaps two cities in the route.
func (self *TravelState) Move() {
	a := rand.Intn(len(self.state))
	b := rand.Intn(len(self.state))
	self.state[a], self.state[b] = self.state[b], self.state[a]
}

// Calculates the length of the route.
func (self *TravelState) Energy() float64 {
	e := 0.0
	for i := 0; i < len(self.state); i++ {
		if i == 0 {
			e += distanceMatrix[self.state[len(self.state)-1]][self.state[0]]
		} else {
			e += distanceMatrix[self.state[i-1]][self.state[i]]
		}
	}
	return e
}

func TestAnneal(t *testing.T) {
	problem := NewAmerica()
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
	problem := NewAmerica()
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
