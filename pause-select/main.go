package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

const _lockdownFile = "/tmp/lockdown"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// What we want here is to loop forever doing periodic work.
// But we want a way of stopping the main working loop waiting on
// a condition of our choosing.
func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// this is the channel we will use to block on a lockdown condition
	lockdownC := make(chan struct{})
	// run a routine that its only purpose is to send a signal to
	// the channel based on some condition of our choosing
	go lockdownSignal(lockdownC)

	for {
		// Do work.
		fmt.Println("Hello im a thing")

		select {
		case <-time.After(time.Second * 3):
			// Block on the lockdown condition
			//
			// Does this mean the context wont be called?
			//
			// This also means we add the time inside the signal loop
			// 	that might be longer than we want...
			<-lockdownC
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func lockdownSignal(lockdownC chan struct{}) {
	for {
		// In our case we just do a boolean on some condition
		if !isLockdownCondition() {
			// if we are not under lockdown send signal
			lockdownC <- struct{}{}
		}

		select {
		case <-time.After(time.Second):
		}
	}
}

func isLockdownCondition() bool {
	if _, err := os.Stat(_lockdownFile); os.IsNotExist(err) {
		return false
	} else if err != nil {
		// demo code so we just bail
		log.Fatal(err)
	}
	return true
}
