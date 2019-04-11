package main

import (
	"Gotest/package/tempconv"
	"flag"
	"fmt"
	"io"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main()  {
	flag.Parse()
	fmt.Println(temp)

	var x io.Writer
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	f(x)
}

func f(out io.Writer)  {
	if out != nil {
		fmt.Printf("%T\n", out)
		fmt.Println(out)
		out.Write([]byte("done!\n"))
	} else {
		fmt.Println(2)
	}
}
