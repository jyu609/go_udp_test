package main

import (
	"fmt"
	"net"
	"os"
)



func main() {
	arguments := os.Args
	portnum := ":" + arguments[1]  //port


	ServerAddr, err := net.ResolveUDPAddr("udp", portnum)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)// buffer from client

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("Received from client ", string(buffer[0:n-1]))

		data := []byte("this info is from server")//data send to client
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}