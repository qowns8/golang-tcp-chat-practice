package util

import (
	"net"
	"fmt"
)

type Client struct{
	connection net.Conn
}

type Server struct {
	clients []Client
}

func (s * Server) GetData(client Client) (*HttpData, []byte) {

	buffer := make([]byte, 1024)

	client.connection.Read(buffer)

	str := ""
	for index, item := range buffer {
		if buffer[index] != 0 {
			str += string(item)
		}
	}

	return s.Parse([]byte(str)), buffer
}

func (s * Server) Parse(bytes []byte) *HttpData {

	result := &HttpData{"", nil, ""}
	str := string(bytes)
	result.method = result.methodParse(str)
	result.header = result.headerParse(bytes)
	if result.method != "GET" {
		result.body = result.bodyParse(str)
	} else {
		result.body = ""
	}

	return result
}

func (s * Server) ConnectHost(client Client) {
	Datas, bytes := s.GetData(client)

	println(client.connection, bytes)

	if (Datas.method == "")  {
		return
	}

	if Datas.body == "CONNECT" {
		client.connection.Write([]byte("HTTP/1.1 200 Connection established\n"))
		s.clients = append(s.clients, client)
		go s.LisenHandling(client)
	}

	return
}

func (s * Server) LisenHandling(client Client) {
	for {
		data, bytes := s.GetData(client)
		if bytes != nil {
			for i := 0; i < len(s.clients); i++ {
				s.clients[i].connection.Write([]byte(data.body))
			}
		}
	}
}

func (s * Server) AcceptClient(server net.Listener) chan Client  {

	channel := make(chan Client)
	go func() {
		for {
			cli, err := server.Accept()
			if cli == nil {
				fmt.Println("acceptClient: Couldn't accept : ", err.Error())
				continue
			}
			channel <- Client{cli}
		}
	}()
	return channel
}