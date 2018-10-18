package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	message = "Ping"
)

func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn.Write([]byte(message))
		log.Printf("Send: %s", message)
		buff := make([]byte, 1024)
		n, _ := conn.Read(buff)
		log.Printf("Receive: %s", buff[:n])
		time.Sleep(time.Second * 1)
		fmt.Scanln()
	}
}

func main() {
	var (
		ip   = "127.0.0.1"
		port = 3333
	)
	SocketClient(ip, port)
}
