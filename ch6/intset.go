package main

import (
	"Gotest/package/bitMap"
	"fmt"
)

func main()  {
	var x, y bitMap.IntSet

	var _ fmt.Stringer = &x
	x.Add(1)
	x.Add(144)
	x.Add(9)
	//	fmt.Println(x.String())

	y.Add(9)
	y.Add(12)
	//	fmt.Println(y.String())

	x.UnionWith(&y)
	//	fmt.Println(x.String())
	//	fmt.Println(x.Has(2), x.Has(9))

	fmt.Println(&x)
	fmt.Println(x.String())
	//fmt.Println(x)
	fmt.Println(x.Len())
	fmt.Println(y.Len())

	x.Remove(9)
	fmt.Println(&x)

	p := x.Copy()
	fmt.Println(p)

	x.Clear()
	fmt.Println(&x)
	fmt.Println(p)

	var initial [64]byte
	var buf []byte
	fmt.Printf("%T %p\n", buf, buf)
	fmt.Println(buf)
	for i := 0; i<len(initial); i++ {
		initial[i] = uint8(i)
	}
	buf = initial[:0]
	fmt.Printf("%T %p\n", buf, buf)
	fmt.Println(buf)

}
