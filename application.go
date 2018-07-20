package main

import (
	"net"
	"./util"
	"strconv"
)

func main() {

	port := 5001

	server, err := net.Listen("tcp", ":" + strconv.Itoa(port))

	if server == nil {
		panic("init: port listening error : " + err.Error())
	}

	defer server.Close()
	proxServer := &util.Server{}
	connections := proxServer.AcceptClient(server)
	for {
		go proxServer.ConnectHost(<-connections)
	}

}
