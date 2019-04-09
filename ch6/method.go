package main

import (
	"fmt"
	"Gotest/package/geometry"
)

func main()  {
	perim := geometry.Path{
		{1,1},
		{5,1},
		{5,4},
		{1,1},
	}

	fmt.Println(perim.Distance())

	p := geometry.Point{1,2}
	pp := &p
	pp.ScaleBy(2)
//	geometry.Point{3,4}.ScaleBy(2) // cannot call pointer method on geometry.Point literal
	fmt.Println(p)

	q := geometry.Point{3,4}
	pp.Distance(q)


}
