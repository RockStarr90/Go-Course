package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	Listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := Listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleInput(conn)
		handleConn(conn)

	}
}

func handleConn(c net.Conn) {
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			c.Close()
			return
		}
		time.Sleep(3 * time.Second)
	}
}

func handleInput(c net.Conn) {
	reader := bufio.NewReader(c)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			c.Close()
			return
		}
		if s == "exit\r\n" {
			c.Close()
			fmt.Printf(s)
			os.Exit(0)
		}
	}
}
