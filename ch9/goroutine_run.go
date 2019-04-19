package main

import "fmt"

func main()  {
	//for i:=0; ; i++ {
	//	go func() {
	//		ch1 := make(chan string)
	//		ch1 <- "ping-pong"
	//	}()
	//
	//	go func() {
	//		ch2 := make(chan string)
	//		ch2 <- "ping-pong"
	//	}()
	//	fmt.Println(i)
	//}

	for ; ;  {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
