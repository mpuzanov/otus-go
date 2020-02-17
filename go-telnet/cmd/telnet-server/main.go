package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	flag "github.com/spf13/pflag"
)

// go build ./cmd/telnet-server
// ./telnet-server -p 8080
// ./telnet-server.exe -p 8080

func main() {
	var port string
	flag.StringVarP(&port, "port", "p", "5000", "listen default port")
	flag.Parse()

	listen, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	defer listen.Close()
	log.Printf("Listen: %s\n", listen.Addr())

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalf("Cannot accept: %v", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	//каждое соединение на 60 сек для проверки
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(60)*time.Second)

	log.Printf("connected %s\n", conn.RemoteAddr())
	conn.Write([]byte(fmt.Sprintf("Welcome to %s, friend from %s\n", conn.LocalAddr(), conn.RemoteAddr())))

	scanner := bufio.NewScanner(conn)
OUTER:
	for {
		select {
		case <-ctx.Done():
			log.Printf("Closing ctx")
			break OUTER
		default:
			if scanner.Scan() {
				text := scanner.Text()
				log.Printf("received from %s: %s", conn.RemoteAddr(), text)
				if text == "quit" || text == "exit" {
					break OUTER
				} else if text != "" {
					conn.Write([]byte(fmt.Sprintf("I have received '%s'\n", text)))
				}
			}
			if err := scanner.Err(); err != nil {
				log.Printf("Error happened on connection with %s: %v", conn.RemoteAddr(), err)
			}
		}
	}
	cancel()
	log.Printf("Closing connection with %s", conn.RemoteAddr())
}
