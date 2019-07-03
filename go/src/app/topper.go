package main

import (
	"time"

	"context"

	"fmt"
)

type Topper struct {
	duration   time.Duration
	topperLock *SmartLock
}

func newTopper(duration time.Duration) *Topper {
	return &Topper{
		duration:   duration,
		topperLock: NewSmartLock("topper-lock"),
	}
}

func (t *Topper) SprinkleTopping(ctx context.Context, flavor string) error {
	t.topperLock.Lock()
	fmt.Printf("Sprinkling topping for %s donut...\n", flavor)
	SleepGaussian(t.duration, t.topperLock.QueueLength())
	if flavor == "chocolate" {
		for i := 0; i < 4; i++ {
			SleepGaussian(t.duration, t.topperLock.QueueLength())
		}
	}
	t.topperLock.Unlock()

	return nil
}
