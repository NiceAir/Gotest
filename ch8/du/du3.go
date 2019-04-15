package main

import (
	"Gotest/package/du"
	"flag"
	"path/filepath"
	"sync"
	"time"
)

func main()  {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSize := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSize)
	}

	go func() {
		n.Wait()
		close(fileSize)
	}()

	var nfiles, nbytes int64
	tick := time.Tick(500 * time.Millisecond)
loop:
	for ; ;  {
		select {
		case size, ok := <- fileSize:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <- tick:
			du.PrintDiskUsage(nfiles, nbytes)
		}
	}
	du.PrintDiskUsage(nfiles, nbytes)

}

func walkDir(dir string, n *sync.WaitGroup, fileSize chan <- int64)  {
	defer n.Done()

	for _, entry := range du.Dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}