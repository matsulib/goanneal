package goanneal

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

type State interface {
	Copy() interface{} // Returns an address of an exact copy of the current state
	Move()             // Move to a different state
	Energy() float64   // Return the energy of the current state
}

// Performs simulated annealing by calling functions to calculate
// energy and make moves on a state.  The temperature schedule for
// annealing may be provided manually or estimated automatically.
type Annealer struct {
	// parameters
	Tmax    float64 // max temperature
	Tmin    float64 // minimum temperature
	Steps   int
	Updates int

	// user settings
	CopyStrategy string
	UserExit     bool

	// placeholders
	State      interface{}
	bestState  interface{}
	bestEnergy float64
	start      float64
}

func NewAnnealer(initialState interface{}) *Annealer {
	a := new(Annealer)
	a.State = initialState
	a.Tmax = 25000.0
	a.Tmin = 2.5
	a.Steps = 50000
	a.Updates = 100
	return a
}

// Takes the output from `auto` and sets the attributes
func (self *Annealer) SetSchedule(schedule map[string]float64) {
	self.Tmax = schedule["tmax"]
	self.Tmin = schedule["tmin"]
	self.Steps = int(schedule["steps"])
	self.Updates = int(schedule["updates"])
}

// Wrapper for internal update.
// If you override the self.update method,
// you can chose to call the self.default_update method
// from your own Annealer.
func (self *Annealer) update(args ...interface{}) {
	self.defaultUpdate(args[0].(int),
		args[1].(float64),
		args[2].(float64),
		args[3].(float64),
		args[4].(float64))
}

// Default update, outputs to stderr.
// Prints the current temperature, energy, acceptance rate, improvement rate, elapsed time, and remaining time.
// The acceptance rate indicates the percentage of moves since the last update
// that were accepted by the Metropolis algorithm.
// It includes moves that decreased the energy, moves that left the energy unchanged,
// and moves that increased the energy yet were reached by thermal excitation.
// The improvement rate indicates the percentage of moves since the last update that strictly decreased the energy.
// At high temperatures it will include both moves that improved the overall state and
// moves that simply undid previously accepted moves that increased the energy by thermal excititation.
// At low temperatures it will tend toward zero as the moves that can decrease the energy are exhausted and
// moves that would increase the energy are no longer thermally accessible.
func (self *Annealer) defaultUpdate(step int, T float64, E float64, acceptance float64, improvement float64) {
	elapsed := now() - self.start
	if step == 0 {
		fmt.Fprintln(os.Stderr, " Temperature        Energy    Accept   Improve     Elapsed   Remaining")
		fmt.Fprintf(os.Stderr, "\r%12.5f  %12.2f                      %s            ", T, E, timeString(elapsed))
	} else {
		remain := float64(self.Steps-step) * (elapsed / float64(step))
		fmt.Fprintf(os.Stderr, "\r%12.5f  %12.2f  %7.2f%%  %7.2f%%  %s  %s",
			T, E, 100.0*acceptance, 100.0*improvement, timeString(elapsed), timeString(remain))
	}
}

// Minimizes the energy of a system by simulated annealing.
// Parameters
// state : an initial arrangement of the system
// Returns
// (state, energy): the best state and energy found.
func (self *Annealer) Anneal() (interface{}, float64) {
	step := 0
	self.start = now()

	// Precompute factor for exponential cooling from Tmax to Tmin
	if self.Tmin <= 0.0 {
		panic("Exponential cooling requires a minimum temperature greater than zero.")
	}
	Tfactor := -math.Log(self.Tmax / self.Tmin)

	// Note initial state
	currentState := self.State.(State)
	T := self.Tmax
	E := currentState.Energy()
	prevState := currentState.Copy()
	prevEnergy := E
	self.bestState = currentState.Copy()
	self.bestEnergy = E
	traials, accepts, imporoves := 0, 0.0, 0.0
	var updateWavelength float64
	if self.Updates > 0 {
		updateWavelength = float64(self.Steps) / float64(self.Updates)
		self.update(step, T, E, 0.0, 0.0)
	}

	// Attempt moves to new states
	for step < self.Steps && !self.UserExit {
		//time.Sleep(100 * time.Millisecond)
		step++
		T = self.Tmax * math.Exp(Tfactor*float64(step)/float64(self.Steps))
		currentState.Move()
		E := currentState.Energy()
		dE := E - prevEnergy
		traials++
		if dE > 0.0 && math.Exp(-dE/T) < rand.Float64() {
			// Restore previous state
			currentState = prevState.(State).Copy().(State)
			E = prevEnergy
		} else {
			// Accept new state and compare to best state
			accepts += 1.0
			if dE < 0.0 {
				imporoves += 1.0
			}
			prevState = currentState.Copy()
			prevEnergy = E
			if E < self.bestEnergy {
				self.bestState = currentState.Copy()
				self.bestEnergy = E
			}
		}
		if self.Updates > 1 {
			if (step / int(updateWavelength)) > ((step - 1) / int(updateWavelength)) {
				self.update(step, T, E, accepts/float64(traials), imporoves/float64(traials))
				traials, accepts, imporoves = 0, 0.0, 0.0
			}
		}
	}
	fmt.Fprintln(os.Stderr, "")

	// Return best state and energy
	return self.bestState, self.bestEnergy
}

// Explores the annealing landscape and estimates optimal temperature settings.
// Returns a dictionary suitable for the `set_schedule` method.
func (self *Annealer) Auto(minutes float64, steps int) map[string]float64 {

	// Anneals a system at constant temperature and returns the state,
	// energy, rate of acceptance, and rate of improvement.
	run := func(T float64, steps int) (float64, float64, float64) {
		currentState := self.State.(State)
		E := currentState.Energy()
		prevState := currentState.Copy()
		prevEnergy := E
		accepts, improves := 0.0, 0.0
		for i := 0; i < steps; i++ {
			currentState.Move()
			E = currentState.Energy()
			dE := E - prevEnergy
			if dE > 0.0 && math.Exp(-dE/T) < rand.Float64() {
				currentState = prevState.(State).Copy().(State)
				E = prevEnergy
			} else {
				accepts += 1.0
				if dE < 0.0 {
					improves += 1.0
				}
				prevState = currentState.Copy()
				prevEnergy = E
			}
		}
		return E, float64(accepts) / float64(steps), float64(improves) / float64(steps)
	}

	step := 0
	self.start = now()

	// Attempting automatic simulated anneal...
	// Find an initial guess for temperature
	currentState := self.State.(State)
	T := 0.0
	E := currentState.Energy()
	self.update(step, T, E, 0.0, 0.0)
	for T == 0.0 {
		step++
		currentState.Move()
		T = math.Abs(currentState.Energy() - E)
	}

	// Search for Tmax - a temperature that gives 98% acceptance
	E, acceptance, improvement := run(T, steps)

	step += steps
	for acceptance > 0.98 {
		T = roundFigure(T/1.5, 2)
		E, acceptance, improvement = run(T, steps)
		step += steps
		self.update(step, T, E, acceptance, improvement)
	}
	for acceptance < 0.98 {
		T = roundFigure(T*1.5, 2)
		E, acceptance, improvement = run(T, steps)
		step += steps
		self.update(step, T, E, acceptance, improvement)
	}
	Tmax := T

	// Search for Tmin - a temperature that gives 0.02% improvement
	for improvement > 0.02 {
		T = roundFigure(T/1.5, 2)
		E, acceptance, improvement = run(T, steps)
		step += steps
		self.update(step, T, E, acceptance, improvement)
	}
	Tmin := T

	// Calculate anneal duration
	elapsed := now() - self.start
	duration := roundFigure(60.0*minutes*float64(step)/elapsed, 2)

	// Don't perform anneal, just return params
	return map[string]float64{"tmax": Tmax, "tmin": Tmin, "steps": duration, "updates": float64(self.Updates)}
}
