package main

import (
	"time"

	"fmt"

	"context"
)

type Fryer struct {
	fryerLock *SmartLock
	duration  time.Duration
}

func newFryer(duration time.Duration) *Fryer {
	return &Fryer{
		duration:  duration,
		fryerLock: NewSmartLock("fryer-lock"),
	}
}

func (f *Fryer) FryDonut(ctx context.Context, flavor string) {
	f.fryerLock.Lock()
	defer f.fryerLock.Unlock()

	fmt.Printf("Frying %s donut...\n", flavor)
	SleepGaussian(f.duration, f.fryerLock.QueueLength())
}
