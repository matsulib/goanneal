// +build ignore

package main

import (
	"fmt"
	"math/rand"

	"github.com/matsulib/goanneal"
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

func main() {
	problem := NewAmerica()
	distanceMatrix = problem.DistanceMatrix
	// initial state, a randomly-ordered itinerary
	initialState := &TravelState{state: problem.CitiesKeys()}
	shuffle(initialState.state)

	tsp := goanneal.NewAnnealer(initialState)
	//tsp.Steps = 5000000 * 4
	autoSchedule := tsp.Auto(1, 20000)
	tsp.SetSchedule(autoSchedule)
	fmt.Println(tsp.Tmax, tsp.Tmin, tsp.Steps, tsp.Updates)

	state, energy := tsp.Anneal()
	ts := state.(*TravelState)
	for ts.state[0] != "Grand Canyon, AZ" {
		ts.state = append(ts.state[1:], ts.state[:1]...)
	}

	fmt.Println(int(energy), "mile route:")
	for i := 0; i < len(ts.state); i++ {
		fmt.Println("\t", ts.state[i])
	}
}
