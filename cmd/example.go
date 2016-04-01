package main

import (
	"log"
	"time"

	"github.com/landonia/goat"
)

func main() {

	// Create a new goat
	g := goat.New(time.Second*5,
		func() {
			log.Println("Executed")
		},
	)

	// Start the execution
	err := g.Start()
	if err != nil {
		// It means that it has already been started
		log.Println("Already started")
	}

	<-time.After(time.Minute)

	// Ready to shutdown
	err = g.Stop()
	if err != nil {
		// It means that it was not running
		log.Println("Not started")
	}

	<-time.After(time.Second * 10)
}
