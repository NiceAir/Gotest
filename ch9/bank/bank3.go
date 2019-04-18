package main

import "sync"

var (
	mu sync.Mutex
	blance int
)

func Deposit(amount int)  {
	mu.Lock()
	defer mu.Unlock()
	blance += amount

}

func Blance() int {
	mu.Lock()
	defer mu.Unlock()
	b := blance
	return b
}

