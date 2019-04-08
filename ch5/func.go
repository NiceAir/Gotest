package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main()  {
	f := squares()
	ff := squares

	fmt.Println(ff()())
	fmt.Println(ff()())
	fmt.Println(ff()())
	fmt.Println(ff()())
	fmt.Println("----------------")
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())


}