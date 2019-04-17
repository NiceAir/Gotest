package main

import "fmt"

type Cake struct {
	state string
}

func baker(cooked chan <- *Cake)  {
	for ; ;  {
		cake := new(Cake)
		fmt.Printf("%T", cake)
		return
		cake.state = "cooked"
		cooked <- cake
	}
}

func icer(iced chan <- *Cake, cooked <- chan *Cake)  {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake
	}
}


func main()  {

}