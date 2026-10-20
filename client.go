package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println(msg)
	}
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	conn, err := net.Dial("tcp", "127.0.0.1:8030")
	if err != nil {
		// probably no server running
		fmt.Println(err.Error())
		return
	}
	for {
		fmt.Println("Enter text:")
		text, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Fprintln(conn, text)
		read(&conn)
	}
}
