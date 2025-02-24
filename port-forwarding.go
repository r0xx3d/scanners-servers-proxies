package main


import (
	"io"
	"net"
	"log"
)


func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "somerandomwebsite.com")

