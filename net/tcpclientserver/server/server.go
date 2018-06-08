package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"strconv"
	"time"
)

const (
	Message = "Pong"
)

func socketServer(port int) {
	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handler(conn)
		Writer(conn)
	}

}

func Writer(conn net.Conn) {
	time.Sleep(time.Second * 5)
	var w = bufio.NewWriter(conn)
	var i int
	for {
		i++
		myPong := Message + ":" + strconv.Itoa(i)
		_, err := w.Write([]byte(myPong))
		if err != nil {
			break
		}
		w.Flush()
		fmt.Println("send:", myPong)
		time.Sleep(time.Second)
		runtime.Gosched()
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)
iLOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])
		switch err {
		case io.EOF:
			break iLOOP
		case nil:
			log.Println("Receive:", data)
		default:
			fmt.Println(err)
			return
		}
	}
	w.Write([]byte(Message))
	w.Flush()
	fmt.Println(Message)
}

func main() {
	socketServer(3333)
}
