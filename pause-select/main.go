package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/atomic"
)

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

	lockdown := atomic.NewBool(false)
	lockdownC := make(chan struct{})

	go lockdownSignal(lockdown, lockdownC)

	for {
		fmt.Println("Hello im a thing")

		select {
		case <-time.After(time.Second * 3):
			// does this mean the context wont be called?
			<-lockdownC
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func lockdownSignal(lockdown *atomic.Bool, lockdownC chan struct{}) {
	for {
		checkLockdownCondition(lockdown)

		if !lockdown.Load() {
			lockdownC <- struct{}{}
		}

		select {
		case <-time.After(time.Second):
		}
	}
}

func checkLockdownCondition(lockdown *atomic.Bool) {
	if _, err := os.Stat("/tmp/lockdown"); os.IsNotExist(err) {
		lockdown.Store(false)
	} else {
		lockdown.Store(true)
	}
}
