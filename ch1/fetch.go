// Fetch prints the content found at a URL.
package main

import (
	"fmt"
//	"io/ioutil"
	"os"
	"net/http"
	"io"
	"strings"
)

func main()  {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		response, err := http.Get(url)
		//fmt.Println(response.Status)
		//os.Exit(1)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//b, err := ioutil.ReadAll(response.Body)
		b, err := io.Copy(os.Stdout,response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Printf("%s\n", b)

	}
}