package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	flag "github.com/spf13/pflag"
)

var timeout time.Duration

func init() {
	flag.Usage = func() {
		fmt.Printf("usage : %s host port\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.DurationVarP(&timeout, "timeout", "t", 10*time.Second, "timeout for connecting to the server")
	flag.Parse()
	if len(flag.Args()) < 2 {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	host := flag.Args()[0]
	port := flag.Args()[1]
	log.Printf("host=%s, port=%s, timeout=%v", host, port, timeout)

	dialer := &net.Dialer{}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)

	conn, err := dialer.DialContext(ctx, "tcp", host+":"+port) //conn
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		log.Fatalf("Get SIGINT signal (Ctrl+C)")
	}()

	chanExit := make(chan int, 1)
	//горутина берёт из STDIN и пишет в сокет
	go writeRoutine(chanExit, conn)

	//горутина получает из сокета и выводит в STDOUT
	go readRoutine(chanExit, conn)

	<-chanExit

	log.Print("closing the socket")
	cancel()
	conn.Close()
	log.Println("closed")
}

func readRoutine(chanExit chan int, conn net.Conn) {
	scanner := bufio.NewScanner(conn)
OUTER:
	for {
		select {
		case <-chanExit:
			break OUTER
		default:
			if !scanner.Scan() {
				log.Println("server close connection")
				break OUTER
			}
			text := scanner.Text()
			log.Printf("From server: %s", text)
		}
	}
	//log.Printf("Finished readRoutine")
}

func writeRoutine(chanExit chan int, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
OUTER:
	for {
		if scanner.Scan() {
			text := scanner.Text()
			if text == "exit" {
				conn.Write([]byte("exit")) // send exit to server
				break OUTER
			} else if text != "" {
				log.Printf("To server: %v\n", text)
				_, err := conn.Write([]byte(fmt.Sprintf("%s\n", text)))
				if err != nil {
					log.Printf("cannot write to connection: %v", err)
					break OUTER
				}
			}
		} else { // EOF
			conn.Write([]byte("exit")) // send exit to server
			break OUTER
		}
	}
	log.Printf("close connection\n")
	chanExit <- 1
	//log.Printf("Finished writeRoutine")
}

//go run . --timeout 15s 127.0.0.1 8080
