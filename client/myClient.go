package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	msgP := flag.String("msg", "Default message", "The message you want to send")
	conn, err := net.Dial("tcp", "127.0.0.1:8030")
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(conn, *msgP)
}
