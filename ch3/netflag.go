package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoolback
	FlagPointPoint
	FlagMulticast
)

func IsUp(v Flags) bool {
	return v & FlagUp == FlagUp
}

func TurnDown(v *Flags)  {
	*v &^= FlagUp
}
func SetBroadcast(v *Flags)  {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v & (FlagBroadcast|FlagMulticast) != 0
}
func main()  {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
	fmt.Println(13 &^ 4)
	fmt.Println(13 &^ 6)
	fmt.Printf("%T ", 212.0-22)
	fmt.Println(212.0-22)
	fmt.Printf("%T ", 22.0/5)
	fmt.Println(22.0/5)
	fmt.Printf("%T ", 22/5.0)
	fmt.Println(22/5.0)
	fmt.Printf("%T ", 22.0/5.0)
	fmt.Println(22.0/5.0)
	// 1101  13
	// 0110   4
	// 1001   9
}