package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	Message       = "Pong"
	Message2      = "Pong222"
	StopCharacter = "\r\n\r\n"
)

func SocketServer(port int) {

	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		// os.Exit(1)
	}
	log.Printf("Begin listen port: %d", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("reset connection........", err)
			continue
		}
		go handler(conn)
	}

}

func handler(conn net.Conn) {
	fmt.Println("New connection........")
	defer conn.Close()

	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

	// ILOOP:
	for {
		log.Println("Start read....")
		n, err := r.Read(buf)
		data := string(buf[:n])
		log.Println("Read from buffer complite......")

		switch err {
		case io.EOF:
			log.Println("io.EOF:", err)
			return
		case nil:
			log.Println("receive:", data)
		default:
			log.Printf("Receive data failed:%s\n", err)
			return
		}

		w.Write([]byte(Message))
		w.Flush()
		log.Printf("Send: %s", Message)

		w.Write([]byte(Message2))
		w.Flush()
		log.Printf("Send: %s", Message2)

	}

}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return
}

func main() {

	port := 3333

	SocketServer(port)

}
