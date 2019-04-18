package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu sync.RWMutex
	blance int
)

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if blance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func Deposit(amount int)  {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return blance
}

func deposit(amount int)  {
	blance += amount
}
