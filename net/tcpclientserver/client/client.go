package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	message       = "Ping"
	StopCharacter = "\r\n\r\n"
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
		_, err := conn.Write([]byte(message + strconv.Itoa(i)))
		if err != nil {
			fmt.Println("err: ", err)
			break
		}
		fmt.Println("Send: ", message)
		runtime.Gosched()
	}
}

func Reader(conn net.Conn) {

	for {
		time.Sleep(time.Second)
		chunk = make([]byte, 1024)
		bytesWasRead, err := conn.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err", err)
			conn.Close()
			return
		}
		buffer = append(buffer, chunk[:bytesWasRead]...)
		if bytesWasRead < 1024 {
			if bytesWasRead > 0 {
				fmt.Println("Read: ", string(buffer[:bytesWasRead]))
			}
			buffer = make([]byte, 0)
		}
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
