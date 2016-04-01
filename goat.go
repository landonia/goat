// Copyright 2016 Landon Wainwright.

// Package goat provides a go asynchronous ticker wrapper
package goat

import (
	"fmt"
	"sync"
	"time"
)

// Goat is the type for the interval executioner
type Goat struct {
	duration time.Duration
	handler  func()
	exit     chan bool
	r        bool
	mut      sync.Mutex
}

// New will create a new goat that will call the function at every
// duration interval.
func New(duration time.Duration, handler func()) *Goat {
	return &Goat{
		duration: duration,
		handler:  handler,
		exit:     make(chan bool),
	}
}

// Start will start the ticker execution routine if it is not currently executing
func (goat *Goat) Start() error {
	goat.mut.Lock()
	defer goat.mut.Unlock()
	if goat.r {
		return fmt.Errorf("Already executing - Stop first to start a new goat")
	}
	goat.r = true

	// Launch the routine for handling the ticker execution
	go func() {
		ticker := time.NewTicker(goat.duration)
		for {
			select {
			case <-goat.exit:
				ticker.Stop()
				return
			case <-ticker.C:

				// Call the handler
				goat.handler()
			}
		}
	}()
	return nil
}

// Stop will stop the execution routine and will block until the shutdown is
// received
func (goat *Goat) Stop() error {
	goat.mut.Lock()
	defer goat.mut.Unlock()
	if !goat.r {
		return fmt.Errorf("Not executing - Start first to stop a goat")
	}
	goat.exit <- true
	goat.r = false
	return nil
}
