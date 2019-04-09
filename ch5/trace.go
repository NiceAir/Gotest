package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation()  {
	defer trace("bigSlowOperation")()
//	extra parentheses
	fmt.Println(1111)
	time.Sleep(10 * time.Second)
//	operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	fmt.Println(2222)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main()  {
	bigSlowOperation()
}