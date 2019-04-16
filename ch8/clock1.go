package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main()  {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for ; ; {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn)  {
	defer c.Close()
	for i := 0; ; i++ {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		fmt.Println(i)
		if err != nil {
			return
		}
		time.Sleep(1*time.Second)
	}
}