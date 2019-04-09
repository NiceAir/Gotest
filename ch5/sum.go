package main

import (
	"fmt"
	"os"
)

func sum(vals ... int) int {
	var total int = 0
	for _, value := range vals {
		total += value
	}
	return total
}

func main()  {
	var a = [3]int{1,2,3}
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1,2,3))
	fmt.Println(sum(1,2,3,4))
	fmt.Println(sum(a[:]...))

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}

func errorf(linenum int, format string, args ... interface{})  {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args ...)
	fmt.Fprintln(os.Stderr)
}