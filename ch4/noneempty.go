package main

import "fmt"

func noneempty(strings []string) []string  {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return  strings
}

func noneempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out,s)
		}
	}
	fmt.Printf("%p ", out)
	fmt.Println(out)

	fmt.Printf("%p ", strings)
	fmt.Println(strings)
	return out
}

func main()  {
	str := []string{"one", "two", "", "threee"}
	noneempty2(str)
}
