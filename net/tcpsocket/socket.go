package main

import (
	"bufio"
	"fmt"
	"net"
)

func HandleServercConnection(conn net.Conn) {
	name := conn.RemoteAddr().String()
	fmt.Printf("%+v connected\n", name)
	conn.Write([]byte("Hello" + name + "\n\r"))
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "Exit" {
			conn.Write([]byte("good bye\n\r"))
			break
		} else if text != "" {
			conn.Write([]byte("You entered text " + text + "\r\n"))
		}
	}
}

func Server() {
	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go HandleServercConnection(conn)
	}
}

func main() {
	Server()
}
