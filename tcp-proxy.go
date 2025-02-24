package main

import (
	"log"
	"io"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("Host unreachable...")
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp",":8888")
	if err != nil {
		log.Fatalln("Unable to bind port...")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection...")
		}
		go handle(conn)
	}
}
