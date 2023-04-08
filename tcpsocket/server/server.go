package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
	conns []net.Conn
	connCh = make(chan net.Conn)
	closeCh = make(chan net.Conn)
	msgCh = make(chan string)
)

func main() {
	listener, error := net.Listen("tcp", ":333")
	if error != nil {
		log.Fatal(error)
	}
	defer listener.Close()
	go func() {
		for {
			conn, error := listener.Accept()
			if error != nil {
				log.Fatal("connection is error")
			}
			conns = append(conns, conn)
			connCh <- conn
		}
	}()
	
	for {
		select {
		case conn := <- connCh:
			go onMessage(conn)
		case msg := <- msgCh:
			fmt.Print(msg)
		case conn := <- closeCh:
			fmt.Println("close connection")
			disConnect(conn)
		}
	}
	
}

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		
		msgCh <- msg
		publishMessage(conn, msg)
	}
	closeCh <- conn
}

func publishMessage(conn net.Conn, message string) {
	for _, connItem := range conns {
		if connItem != conn {
			connItem.Write([]byte(message))
		}
	}
}
func disConnect(conn net.Conn) {
	var i int;
	for i := range conns {
		if conns[i] == conn {
			break;
		}
	}
	conns = append(conns[:i], conns[i+1:]...)
}