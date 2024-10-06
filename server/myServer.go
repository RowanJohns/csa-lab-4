package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err0 := net.Listen("tcp", ":8030")
	if err0 != nil {
		panic(err0)
	}
	conn, err1 := ln.Accept()
	if err1 != nil {
		panic(err1)
	}
	reader := bufio.NewReader(conn)
	msg, err2 := reader.ReadString('\n')
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(msg)
}
