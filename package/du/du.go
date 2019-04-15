package du

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WalkDir(dir string, fileSizes chan <- int64)  {
	for _, entry := range Dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			WalkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}
var sema = make(chan struct{}, 20)  //并发量限制在20

func Dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {<- sema}()

	entries, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func PrintDiskUsage(nfiles, nbytes int64)  {
	fmt.Printf("%d files    %.3f GB\n", nfiles, float64(nbytes)/1e9)
}