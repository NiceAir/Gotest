package main

import (
	"fmt"
	"log"
	"os"
)

// pc[i] is the population count of i.
var pc [256]byte
var cwd string
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}


func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}


func main()  {
	//var j = 0
	//for i := range pc {
	//	fmt.Printf("%d %T   ", pc[i], pc[i]);
	//	j++
	//	if j >= 8 {
	//		fmt.Printf("\n");
	//		j = 0
	//	}
	//}
	fmt.Println(1i*1i);
}
