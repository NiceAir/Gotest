package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main()  {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)    //io.Copy本身会一直循环，当任意一端关闭后才退出循环
		log.Println("服务器断开...")
		done <- struct{}{}
	}()
	go mustCopy(conn, os.Stdin)
	<- done
	conn.Close()
	//os.Stdin.Close()

}

func mustCopy(dst io.Writer, src io.Reader)  {
	if res, err := io.Copy(dst, src); err != nil {
		log.Fatal(err, res)
	}
}