package main

import "fmt"

func main()  {
	// r := [...]int{99:-1}
	//var arr [3]int = [3]int{1,2,3}
	//arr1 := [3]int{}
	//a := []int{1,2}
	//b := [2]int{2,1}
	//c := [2]int{1,2}
	//d := [...]int{2,1}
	//fmt.Println(r)
	//fmt.Println(arr)
	//fmt.Println(arr1)
	//fmt.Println(a == b)
	//fmt.Println(a == c)
	//fmt.Println(b == d)
//	e := [3]int{1,2}
//	fmt.Println(a == e)

	a := [3]int{1,2,3}
	fmt.Println(cap(a))
	fmt.Println(len(a))
	arrayTest(a)
	fmt.Println(a)

}

func arrayTest(array [3]int)  {
	for i:= len(array)-1; i>=0; i-- {
		array[i] = 0
	}
}
