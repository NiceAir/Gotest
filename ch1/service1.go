package main

import (
	"fmt"
	"log"
	"net/http"
)


func main()  {
	http.HandleFunc("/", handler0)
	http.HandleFunc("/1", handler1)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler0(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handler1(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "你好")

}