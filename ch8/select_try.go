package main

import (
	"fmt"
	"os"
)

func main()  {
	chnnel_a := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		chnnel_a <- struct{}{}
	}()
loop:    //notice: 这个标签为必须的，否则就会陷入for死循环!
	for i := 0; ; i++ {
		select {
		case a := <- chnnel_a:
			fmt.Fprintln(os.Stdout, a)
			break loop
		default:
		}
		fmt.Fprintf(os.Stdout, "第%d次循环\n", i)
	}
	fmt.Println("end")
}
