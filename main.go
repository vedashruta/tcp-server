package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8989")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	ch := make(chan struct{}, 1)
	defer close(ch)
	for {
		connection, err := listener.Accept()
		if err != nil {
			return
		}
		go handler(connection)
	}
}

func handler(conn net.Conn) {
	var payload []byte
	defer conn.Close()
	for {
		var temp []byte
		_, err := conn.Read(temp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		payload = append(payload, temp...)
	}
	fmt.Println(string(payload))
}
