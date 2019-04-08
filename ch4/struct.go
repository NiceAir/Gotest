package main

import "fmt"

type point struct {
	X, Y	int
}

type Circle struct {
	point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main()  {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	var x = Wheel{
		Circle:Circle{
			point:point{X:8, Y:8},
		//	Radius:5,
		},
		Spokes:20,
	}

	fmt.Printf("%#v\n", w)
	fmt.Printf("%v\n", x)
}