package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"bufio"
	"strconv"
)

type ConnData struct {
	CONN_HOST string
	CONN_PORT string
	POST_HEADER string
}

func main() {

	header, port := askPort()
	connData := ConnData{"localhost", port, header}

	tcpAddr, err := net.ResolveTCPAddr("tcp", connData.CONN_HOST+":"+connData.CONN_PORT)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to server through tcp.
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	connData.connect(conn)
	go printOutput(conn)
	connData.writeInput(conn)
}

func askPort() (string, string) {
	port := 0
	fmt.Print("Port : ")
	fmt.Scanf("%d", &port)
	return "POST  HTTP/1.1\r\nHost: 127.0.0.1:" + strconv.Itoa(port) + "\r\nContent-Type: text/plain\r\nUser-Agent: chatting\r\n\r\n", strconv.Itoa(port)
}

func (c * ConnData)post(head string, body string) string{
	if head == "" {
		return c.POST_HEADER + body
	} else {
		return head + body
	}
}

func (c* ConnData)connect(conn *net.TCPConn) {
	conn.Write([]byte(c.post("", "CONNECT")))
}

func (c* ConnData)writeInput(conn *net.TCPConn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		println("input data =>")
		text, err := reader.ReadString('\n')
		re, err := conn.Write([]byte(c.post("", text)))
		if err != nil {
			fmt.Printf("%d :", re)
			log.Println(err)
		}
	}
}

func printOutput(conn *net.TCPConn) {
	for {
		msg := make([]byte, 256)
		num, err := conn.Read(msg)
		if err == io.EOF {
			conn.Close()
			fmt.Println("Connection Closed. Bye bye.")
			os.Exit(0)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(msg[:num]))
	}
}