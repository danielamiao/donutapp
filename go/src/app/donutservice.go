package main

import (
	"time"

	"context"
)

const (
	fryDuration = time.Millisecond * 550
	payDuration = time.Millisecond * 250
	topDuration = time.Millisecond * 150
)

type DonutService struct {
	payer  *Payer
	fryer  *Fryer
	topper *Topper
}

func newDonutService() *DonutService {
	return &DonutService{
		payer:  NewPayer(payDuration),
		fryer:  newFryer(fryDuration),
		topper: newTopper(topDuration),
	}
}

func (ds *DonutService) makeDonut(ctx context.Context, flavor string) error {
	ds.payer.BuyDonut(ctx, flavor)
	ds.fryer.FryDonut(ctx, flavor)
	ds.topper.SprinkleTopping(ctx, flavor)
	return nil
}
