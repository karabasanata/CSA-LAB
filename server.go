package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			// probably the client has disconnected
			fmt.Println(err.Error())
			return
		}
		fmt.Println(msg)
		fmt.Fprintln(*conn, "OK")
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8030")
	if err != nil {
		// probably another server listening on this port already
		fmt.Println(err.Error())
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go handleConnection(&conn)
	}

}
