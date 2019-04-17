package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

var alivetime int = 10     //单位秒

func main()  {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for ; ;  {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string

var (
	entering = make(chan client)  //加入聊天室
	leaving = make(chan client)   //离开聊天室
	messages = make(chan string)  //公共消息
)

func broadcaster()  {
	clients := make(map[client]bool)

	for ; ;  {
		select {
		case msg := <- messages:     //进行广播
			for cli := range clients {
				select {
				case cli <- msg:
				default:
				}
			}
		case cli := <- entering:      //加入
			clients[cli] = true
		case cli := <- leaving:       //离开
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn)  {
	ch := make(chan string)

	go clientwriter(conn, ch)  //获取客户端地址，并将客户端信息发送至conn

	who := conn.RemoteAddr().String()  //获取客户端ip和端口号
	ch <- who
	messages <- who + " 加入聊天室"
	entering <- ch


	active_begin := make(chan time.Time)
	timeout := make(chan struct{}, 1)    //notice: 这个channel的容量必须不能是0，否则会死锁
	go func() {
		var end = time.Time{}
		begin := time.Now()
		for ; ;  {
			select {
			case begin = <- active_begin:
			default:
			}
			end = time.Now()
			if end.Sub(begin).Seconds() > float64(alivetime) {
				// 超时
				timeout <- struct{}{}
				conn.Close()
				break
			}
			time.Sleep(500 * time.Millisecond)
			fmt.Print(who + " ")
			fmt.Println(end.Sub(begin))
		}
		fmt.Print(who + " ")
		fmt.Println(end.Sub(begin))
	}()

	input := bufio.NewScanner(conn)  //读取conn中的信息，并发送至公共消息channel：message

	for input.Scan()  {
		mes := who + ": " + input.Text()
		messages <- mes
		fmt.Printf("%d %d %s say: %s\n", cap(active_begin), len(active_begin),who, mes)
		active_begin <- time.Now()

	}


	select {
	case <-timeout:
		messages <- who + "超时"
		break
	default:
	}

	fmt.Print(who + " ")
	fmt.Println("close\n\n")
	leaving <- ch
	messages <- who + "离开了聊天室"
	conn.Close()
}

func clientwriter(conn net.Conn, ch <- chan string)  {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}