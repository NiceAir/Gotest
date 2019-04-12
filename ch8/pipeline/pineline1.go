package main

import (
	"fmt"
	"time"
)

func main()  {
	naturals := make(chan int)
	squares := make(chan int)

	//go func() {
	//	for x := 0; ; x++ {
	//		naturals <- x
	//		time.Sleep(1*time.Second)
	//	}
	//}()
	go h1(naturals)

	//go func() {
	//	for ; ;  {
	//		x := <- naturals
	//		squares <- x * x
	//	}
	//}()
	go h2(squares, naturals)

	//for ; ;  {
	//	fmt.Println(<-squares)
	//}
	h3(squares)

}

func h1(naturals chan int)  {
	for x := 0; x < 5; x++ {
		naturals <- x
		time.Sleep(1*time.Second)
	}
	close(naturals)
	fmt.Println("关闭naturals\n")
}

func h2(squares chan int, naturals chan int)  {
	//for ; ;  {
	//	x , ok:= <- naturals
	//	if !ok {
	//		close(squares)
	//		fmt.Println("关闭squares\n")
	//		break
	//	}else {
	//		squares <- x * x
	//	}
	//
	//}
	for x := range naturals {
		squares <- x * x
	}
	close(squares)
	fmt.Println("关闭squares\n")
}

func h3(squares chan int)  {
	for ; ;  {
		x, ok := <- squares
		if !ok {
			break
		}else {
			fmt.Println(x)
		}
	}
}