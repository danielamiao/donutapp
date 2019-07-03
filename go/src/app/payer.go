package main

import (
	"time"

	"context"

	"fmt"
)

type Payer struct {
	duration time.Duration
}

func NewPayer(duration time.Duration) *Payer {
	return &Payer{
		duration: duration,
	}
}

func (m *Payer) BuyDonut(ctx context.Context, flavor string) {
	fmt.Printf("Paying for %s donut...\n", flavor)
	SleepGaussian(m.duration, 1)
}
