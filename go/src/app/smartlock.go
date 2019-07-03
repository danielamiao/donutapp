package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	seededRNG = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type SmartLock struct {
	realLock    sync.Mutex
	name        string
	queueLength int64
	acquired    time.Time
}

func NewSmartLock(name string) *SmartLock {
	return &SmartLock{
		realLock: sync.Mutex{},
		name:     name,
	}
}

func (sl *SmartLock) Lock() {
	waiters := atomic.AddInt64(&sl.queueLength, 1)
	before := time.Now()
	sl.realLock.Lock()
	lockDuration := time.Now().Sub(before)

	if lockDuration.Seconds() > 0.01 {
		fmt.Printf("%d waiters for %s\n", waiters, sl.name)
		// trace something?
	}
	atomic.AddInt64(&sl.queueLength, -1)
	sl.acquired = time.Now()
}

func (sl *SmartLock) Unlock() {
	released := time.Now()

	heldTime := released.Sub(sl.acquired)
	if heldTime.Seconds() > 0.01 {
		// trace something?
	}
	sl.realLock.Unlock()
}

func (sl *SmartLock) QueueLength() float64 {
	return float64(sl.queueLength)
}
