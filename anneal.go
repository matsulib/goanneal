package goanneal

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

// State is an interface of a state of a problem.
// These three methods will handle the state.
type State interface {
	Copy() interface{} // Returns an address of an exact copy of the current state
	Move()             // Move to a different state
	Energy() float64   // Return the energy of the current state
}

// Annealer performs simulated annealing by calling functions to calculate
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
	State      State
	bestState  State
	bestEnergy float64
	startTime  float64
}

// NewAnnealer initializes an Annealer struct
func NewAnnealer(initialState State) *Annealer {
	a := new(Annealer)
	a.State = initialState
	a.Tmax = 25000.0
	a.Tmin = 2.5
	a.Steps = 50000
	a.Updates = 100
	return a
}

// SetSchedule takes the output from `auto` and sets the attributes
func (a *Annealer) SetSchedule(schedule map[string]float64) {
	a.Tmax = schedule["tmax"]
	a.Tmin = schedule["tmin"]
	a.Steps = int(schedule["steps"])
	a.Updates = int(schedule["updates"])
}

// Outputs to stderr.
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
func (a *Annealer) update(step int, T float64, E float64, acceptance float64, improvement float64) {
	elapsed := now() - a.startTime
	if step == 0 {
		fmt.Fprintln(os.Stderr, " Temperature        Energy    Accept   Improve     Elapsed   Remaining")
		fmt.Fprintf(os.Stderr, "\r%12.5f  %12.2f                      %s            ", T, E, timeString(elapsed))
	} else {
		remain := float64(a.Steps-step) * (elapsed / float64(step))
		fmt.Fprintf(os.Stderr, "\r%12.5f  %12.2f  %7.2f%%  %7.2f%%  %s  %s",
			T, E, 100.0*acceptance, 100.0*improvement, timeString(elapsed), timeString(remain))
	}
}

// Anneal minimizes the energy of a system by simulated annealing.
// Parameters
// state : an initial arrangement of the system
// Returns
// (state, energy): the best state and energy found.
func (a *Annealer) Anneal() (interface{}, float64) {
	step := 0
	a.startTime = now()

	// Precompute factor for exponential cooling from Tmax to Tmin
	if a.Tmin <= 0.0 {
		panic("Exponential cooling requires a minimum temperature greater than zero.")
	}
	Tfactor := -math.Log(a.Tmax / a.Tmin)

	// Note initial state
	T := a.Tmax
	E := a.State.Energy()
	prevState := a.State.Copy().(State)
	prevEnergy := E
	a.bestState = a.State.Copy().(State)
	a.bestEnergy = E
	traials, accepts, imporoves := 0, 0.0, 0.0
	var updateWavelength float64
	if a.Updates > 0 {
		updateWavelength = float64(a.Steps) / float64(a.Updates)
		a.update(step, T, E, 0.0, 0.0)
	}

	// Attempt moves to new states
	for step < a.Steps && !a.UserExit {
		step++
		T = a.Tmax * math.Exp(Tfactor*float64(step)/float64(a.Steps))
		a.State.Move()
		E := a.State.Energy()
		dE := E - prevEnergy
		traials++
		if dE > 0.0 && math.Exp(-dE/T) < rand.Float64() {
			// Restore previous state
			a.State = prevState.Copy().(State)
			E = prevEnergy
		} else {
			// Accept new state and compare to best state
			accepts += 1.0
			if dE < 0.0 {
				imporoves += 1.0
			}
			prevState = a.State.Copy().(State)
			prevEnergy = E
			if E < a.bestEnergy {
				a.bestState = a.State.Copy().(State)
				a.bestEnergy = E
			}
		}
		if a.Updates > 1 {
			if (step / int(updateWavelength)) > ((step - 1) / int(updateWavelength)) {
				a.update(step, T, E, accepts/float64(traials), imporoves/float64(traials))
				traials, accepts, imporoves = 0, 0.0, 0.0
			}
		}
	}
	fmt.Fprintln(os.Stderr, "")

	// Return best state and energy
	return a.bestState, a.bestEnergy
}

// Auto explores the annealing landscape and estimates optimal temperature settings.
// Returns a dictionary suitable for the `set_schedule` method.
func (a *Annealer) Auto(minutes float64, steps int) map[string]float64 {

	// Anneals a system at constant temperature and returns the state,
	// energy, rate of acceptance, and rate of improvement.
	run := func(T float64, steps int) (float64, float64, float64) {
		E := a.State.Energy()
		prevState := a.State.Copy().(State)
		prevEnergy := E
		accepts, improves := 0.0, 0.0
		for i := 0; i < steps; i++ {
			a.State.Move()
			E = a.State.Energy()
			dE := E - prevEnergy
			if dE > 0.0 && math.Exp(-dE/T) < rand.Float64() {
				a.State = prevState.Copy().(State)
				E = prevEnergy
			} else {
				accepts += 1.0
				if dE < 0.0 {
					improves += 1.0
				}
				prevState = a.State.Copy().(State)
				prevEnergy = E
			}
		}
		return E, float64(accepts) / float64(steps), float64(improves) / float64(steps)
	}

	step := 0
	a.startTime = now()

	// Attempting automatic simulated anneal...
	// Find an initial guess for temperature
	T := 0.0
	E := a.State.Energy()
	a.update(step, T, E, 0.0, 0.0)
	for T == 0.0 {
		step++
		a.State.Move()
		T = math.Abs(a.State.Energy() - E)
	}

	// Search for Tmax - a temperature that gives 98% acceptance
	E, acceptance, improvement := run(T, steps)

	step += steps
	for acceptance > 0.98 {
		T = roundFigure(T/1.5, 2)
		E, acceptance, improvement = run(T, steps)
		step += steps
		a.update(step, T, E, acceptance, improvement)
	}
	for acceptance < 0.98 {
		T = roundFigure(T*1.5, 2)
		E, acceptance, improvement = run(T, steps)
		step += steps
		a.update(step, T, E, acceptance, improvement)
	}
	Tmax := T

	// Search for Tmin - a temperature that gives 0.02% improvement
	for improvement > 0.02 {
		T = roundFigure(T/1.5, 2)
		E, acceptance, improvement = run(T, steps)
		step += steps
		a.update(step, T, E, acceptance, improvement)
	}
	Tmin := T

	// Calculate anneal duration
	elapsed := now() - a.startTime
	duration := roundFigure(60.0*minutes*float64(step)/elapsed, 2)

	// Don't perform anneal, just return params
	return map[string]float64{"tmax": Tmax, "tmin": Tmin, "steps": duration, "updates": float64(a.Updates)}
}
