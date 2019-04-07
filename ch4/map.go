package main

import (
	"fmt"
)

func main()  {
	ages := make(map[string]int)
	//ages := map[string]int{}
	ages["haha"] = 21

	ages["charlie"] = 34
	ages["hehehe"] = 2
	ages["qwe"] = 4
	ages["gfsd"] = 5

	//delete(ages, "hahaa")

	//fmt.Println(ages)
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	
}