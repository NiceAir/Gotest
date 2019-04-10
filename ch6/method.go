package main

import (
	"Gotest/package/geometry"
	"fmt"
	"image/color"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{mapping:make(map[string]string)}

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
	pp.Distance(p)
	pp.ScaleBy(2)
	p.ScaleBy(2)
//	geometry.Point{3,4}.ScaleBy(2) // cannot call pointer method on geometry.Point literal
	fmt.Println(p)

	q := geometry.Point{3,4}
	pp.Distance(q)

	red := color.RGBA{255, 0,0, 255}
	blue := color.RGBA{0,0,255,255}
	var ppp = geometry.ColoredPoint{&geometry.Point{1,1}, red}
	var qqq = geometry.ColoredPoint{&geometry.Point{5,4}, blue}

	fmt.Println(ppp.Distance(*qqq.Point))
	ppp.ScaleBy(2)
	qqq.ScaleBy(2)
	fmt.Println(ppp.Distance(*qqq.Point))



}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
