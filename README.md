# goanneal

[Simulated annealing optimization](http://en.wikipedia.org/wiki/Simulated_annealing) in go.

This package is based on a python module, [simanneal](https://github.com/perrygeo/simanneal).

## Usage

1. Create a problem which you want to solve by implementing `State` interface.

    ```go
    type State interface {
        Copy() interface{} // Returns an address of an exact copy of the current state
        Move()             // Move to a different state
        Energy() float64   // Return the energy of the current state
    }
   ```
   In general, Egergy() means a objective function to minimize. 
   
2. Get a new `Annealer` object by calling `NewAnnealer(State)`.
3. `Annealer`'s fields, which are annealing parameters, can be changed.
4. Execute annealing by calling `Annealer.Anneal()`.

Then you'll get an approximate solution.


## Example: Travelling Salesman Problem (TSP)

The quintessential discrete optimization problem is the [travelling salesman problem](http://en.wikipedia.org/wiki/Travelling_salesman_problem). 

Please refer to [examples/salesman](https://github.com/matsulib/goanneal/tree/master/examples/salseman).

### 0. Install it first
```
go get github.com/matsulib/goanneal
```

### 1. Implement `State` interface as `TravelState`
```go
import (
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
```

### 2. Get a new `Annealer` object 
```go
    initialState := &TravelState{state: problem.CitiesKeys()}

    tsp := goanneal.NewAnnealer(initialState)
```

### 3. Change a parameter of the number of iterations.
```go
    tsp.Steps = 100000
```

### 4. Execute annealing
```go
    state, energy := tsp.Anneal()
```

### 5. Output
```go
    ts := state.(*TravelState)
    for ts.state[0] != "New York City" {
        ts.state = append(ts.state[1:], ts.state[:1]...)
    }

    fmt.Println(int(energy), "mile route:")
    for i := 0; i < len(ts.state); i++ {
        fmt.Println("\t", ts.state[i])
    }
```

## Annealing parameters

Getting the annealing algorithm to work effectively and quickly is a matter of tuning parameters. The defaults are:

    Tmax = 25000.0  # Max (starting) temperature
    Tmin = 2.5      # Min (ending) temperature
    Steps = 50000   # Number of iterations
    Updates = 100   # Number of updates (by default an update prints to stdout)

These can vary greatly depending on your objective function and solution space.

See: [simanneal](https://github.com/perrygeo/simanneal)