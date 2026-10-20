package main

import (
	"bufio"
	"flag"
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
	serverAddr := flag.String("ip", "127.0.0.1:8030", "server ip:port")
	flag.Parse()

	stdin := bufio.NewReader(os.Stdin)
	conn, err := net.Dial("tcp", *serverAddr)
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
