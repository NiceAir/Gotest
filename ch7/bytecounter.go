package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error){
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c ByteCounter)String() string {
	return fmt.Sprintf("It is %d\n", c)
}

func main()  {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hahaa, %s", name)
	fmt.Println(c)

}