package main

import "sync"

var (
	mu sync.Mutex
	blance int
)

func Deposit(amount int)  {
	mu.Lock()
	blance += amount
	mu.Unlock()
}

func Blance() int {
	mu.Lock()
	b := blance
	mu.Unlock()
	mu.Unlock()
	return b
}

