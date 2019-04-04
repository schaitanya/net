package net

import (
	"fmt"
	"net"
)

// Connect ...
func Connect(mx string, c chan string) {
	addr := fmt.Sprintf("%s:25", mx)

	fmt.Println("Connecting to: ", addr)

	conn, err := net.DialTCP(addr, nil, nil)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	go initWriter(conn, c)
	go initReader(conn, c)
}

func initWriter(conn *net.TCPConn, c chan string) {
	for {
		msg := <-c
		conn.Write([]byte(msg))
	}
}

func initReader(conn *net.TCPConn, c chan string) {
	for {
		var b []byte
		conn.Read(b)
		c <- string(b)
	}
}
