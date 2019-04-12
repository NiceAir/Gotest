package main

import (
	"fmt"
	"time"
)

func counter(out chan <- int)  {
	for x := 0; x < 5; x++ {
		out <- x
		time.Sleep(500 * time.Millisecond)
	}
	close(out)
}

func squarer(out chan <- int, in <- chan int)  {
	for v := range in  {
		out <- v * v
	}
	close(out)
}

func printer(in <- chan int)  {
	for v := range in {
		fmt.Println(v)
	}
}

func main()  {
	naturals := make(chan int)
	squares := make(chan int)
	fmt.Println(cap(naturals))
	time.Sleep(1*time.Second)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}