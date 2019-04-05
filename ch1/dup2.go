package main

import (
	"fmt"
	"os"
	"bufio"
)

func main()  {

	counts_map := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts_map)
	}else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 %v\n", err)
			}

			countLines(f, counts_map)

			f.Close()
		}
	}

	for line, n := range counts_map {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line);
		}
	}
}

func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	//	fmt.Println(counts)
	}
}