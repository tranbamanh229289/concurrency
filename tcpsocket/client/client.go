package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, error := net.Dial("tcp", "localhost:333")
	if error != nil {
		log.Fatal(error)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		conn.Write([]byte(input))
		go onMessage(conn)
	}
	conn.Close()
}

func onMessage (conn net.Conn){
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(msg)
	}
}