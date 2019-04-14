package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)



	for countdowm := 10; countdowm > 0; countdowm-- {
		fmt.Println(countdowm)
		<- tick
	}
	//launch()
}
