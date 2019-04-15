package main

import (
	"Gotest/package/du"
	"flag"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var	done = make(chan struct{})

func main()  {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSize := make(chan int64)    //传送每个文件的大小
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSize)
	}
	go func() {  //监听键盘，当有事件触发时通过关闭done pipeline 来退出所有walkDir 的goroutine
		os.Stdout.Read(make([]byte, 1))
		close(done)
	}()

	go func() {         //等待全部walkDir的goroutine退出后关闭pipeline
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

func cancelled() bool {
	select {
	case <- done:
		return true
	default:
		return false
	}
}