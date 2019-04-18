package main

import (
	"Gotest/package/memo/memo"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main()  {
	m := memo.New(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs()  {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s,  %s   %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func incomingURLs() []string {
	urls := []string{
		"https://github.com/astaxie/",
		"https://www.baidu.com/",
		"https://books.studygolang.com/gopl-zh/",
		"https://github.com/astaxie/",
		"https://www.baidu.com/",
		"http://www.huya.com/",
		"http://www.huya.com/",
		"https://books.studygolang.com/gopl-zh/",
		}
	return urls
}