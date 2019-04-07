package main

import "fmt"

func main() {
	array1 := [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	//	Q1 := array1[4:7]
	summer := array1[0:12]
	//	fmt.Println(summer)
	//	fmt.Printf("summer: %p\n", summer)
	fmt.Println(array1)
	//	fmt.Printf("array: %p\n", array1)
	//	fmt.Println(cap(summer))
	//	newSummer := summer[:cap(summer)]
	//	fmt.Println(newSummer)
	//	a := summer[1:2]
	//	fmt.Println(a)
	//	summer[0] = 100
	fmt.Println(summer)
	//	fmt.Printf("summer: %p\n", summer)
	//	fmt.Println(array1)
	//	fmt.Printf("array: %p\n", array1)

	Zero(array1)
	fmt.Println(array1)
	//	fmt.Println(summer)

	fmt.Printf("%T\n", summer)
	//Zero(summer)
	//fmt.Println(array1)
	//fmt.Println(summer)

	b := [2]int{1,1}
	bb := b[1:]
	fmt.Println(bb)
	if bb == nil {
		fmt.Println(1)
	}

	var runes []rune
	for _, r := range "Hello 世界"  {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

func Zero(s [12]int)  {
	for i := len(s)-1; i>=0; i-- {
		s[i] = 0
	}
}
