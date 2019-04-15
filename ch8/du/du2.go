package main

import (
	"flag"
	"time"
	"Gotest/package/du"
)



var verbose = flag.Bool("v", false, "show verbose progress message")

func main()  {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0{
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			du.WalkDir(root, fileSizes)
		}
		close(fileSizes)
	}()


	var nfiles, nbytes int64
	//for size := range fileSizes {
	//	nfiles++
	//	nbytes += size
	//}
	var tick <- chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for ; ;  {
		select {
		case size, ook := <-fileSizes:
			if !ook {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			du.PrintDiskUsage(nfiles, nbytes)
		}

	}
	du.PrintDiskUsage(nfiles, nbytes)
}
