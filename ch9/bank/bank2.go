package main

var (
	sema = make(chan struct{}, 1)
	blance int
)

func Deposit(amount int)  {
	sema <- struct{}{}
	blance += amount
	<- sema
}

func Blance() int {
	sema <- struct{}{}
	b := blance
	<- sema
	return b
}
