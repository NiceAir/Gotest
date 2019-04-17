package main

import (
	"fmt"
	"time"
)

func main()  {
	begin := time.Now()
	time.Sleep(3*time.Second)
	end := time.Now()

	res := end.Sub(begin).Seconds()
	fmt.Println(begin)
	fmt.Println(end)
	fmt.Printf("%T\n", res)
	fmt.Println(res)



}
