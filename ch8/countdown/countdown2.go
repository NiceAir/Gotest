package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	abort := make(chan struct {})
	normal := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")

	tick := time.Tick(1 * time.Second)
	go func() {
		for i := 10; i>0; i-- {
			fmt.Println(i)
			<- tick
		}
		normal <- struct{}{}
	}()

	select {
	case <-normal:
		close(normal)
		fmt.Println("发射!")

	case <-abort:
		fmt.Println("发射中断")
	//	return

	//default:
		//fmt.Println("---------------------")
	}
	fmt.Println("退出select");
	time.Sleep(5 * time.Second)
	//launch()
}

