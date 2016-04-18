# goat

A simple wrapper for creating a ticker execution of a function.

## Overview

Goat provides the core `for/select` wrapper that is used to safely execute within
a go routine using a ticker to signal execution of a function.
The function provided will be executed at the ticker duration.
A goat can be shutdown cleanly. A Goat can be reused - if it has been shutdown
it can be started again without creating a new Goat.

## Maturity

Ready to go and is used in production

## Installation

simply run `go get github.com/landonia/goat`

## Use as Library

```go
	package main

	import (
  		"github.com/landonia/goat"
  	)

  	func main() {
  		// Create a new goat
  		g := goat.New(time.Minute,
         		func() {
           			fmt.Println("Executed")
         		},
       		)

       		// Start the execution
       		err := g.Start()
       		if err != nil {
       			// It means that it has already been started
         		fmt.Println("Already started")
       		}

       		// Do work
       		...

       		// Ready to shutdown
       		err = g.Stop()
       		if err != nil {
        		// It means that it was not running
         		fmt.Println("Not started")
       		}
       	}
```
## Example

simply run `go run github/landonia/goat/cmd/example.go`

## About

goat was written by [Landon Wainwright](http://www.landotube.com) | [GitHub](https://github.com/landonia).

Follow me on [Twitter @landoman](http://www.twitter.com/landoman)!
