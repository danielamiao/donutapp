package main

import (
	"fmt"
	"math"
	"time"

	"math/rand"

	"sync"

	"context"
)

const (
	maxQueueDuration = float64(8 * time.Second)
)

func SleepGaussian(d time.Duration, queueLength float64) {
	cappedDuration := float64(d)
	if queueLength > 4 {
		cappedDuration = math.Min(float64(time.Millisecond*50), maxQueueDuration/(queueLength-4))
	}
	// noise := (float64(cappedDuration) / 3) * rand.NormFloat64()
	time.Sleep(time.Duration(cappedDuration))
}

func main() {
	ds := newDonutService()

	// Make fake queries in the background.
	backgroundProcess(10, ds, runFakeUser)
}

func backgroundProcess(max int, ds *DonutService, f func(flavor string, ds *DonutService)) {
	var wg sync.WaitGroup
	for i := 0; i < max; i++ {
		var flavor string
		x := rand.Float32()
		if x < 0.1 {
			flavor = "sprinkles"
		} else if x < 0.3 {
			flavor = "chocolate"
		} else {
			flavor = "cinnamon"
		}
		go f(flavor, ds)
		wg.Add(1)
	}
	wg.Wait()
}

func runFakeUser(flavor string, ds *DonutService) {
	for {
		SleepGaussian(2500*time.Millisecond, 1)
		err := ds.makeDonut(context.Background(), flavor)
		if err != nil {
			fmt.Println(err)
		}
	}
}
