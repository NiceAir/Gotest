package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	db := database{"shoes":51, "socks":5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollors float32

func (d dollors) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollors

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}