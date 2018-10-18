package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	message       = "Ping"
	StopCharacter = "\r\n\r\n"
	lenBuf        = 1024
)

var (
	buffer []byte
	chunk  []byte
)

func Writer(conn net.Conn) {
	var i int
	for {
		i++
		time.Sleep(time.Second * 3)
		fmt.Scanln()
		_, err := conn.Write([]byte(message + strconv.Itoa(i)))
		if err != nil {
			fmt.Println("err: ", err)
			break
		}
		fmt.Println("Send: ", message)
		time.Sleep(time.Second)
	}
}

func Reader(conn net.Conn) {
	for {
		time.Sleep(time.Second)
		chunk = make([]byte, lenBuf)
		bytesWasRead, err := conn.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err", err)
			conn.Close()
			return
		}
		fmt.Println("Read: ", string(chunk[:bytesWasRead]))
	}
}

func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	go Writer(conn)
	go Reader(conn)
	fmt.Scanln()
}

func main() {
	var (
		ip   = "127.0.0.1"
		port = 3333
	)
	SocketClient(ip, port)
}
