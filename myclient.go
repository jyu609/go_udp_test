package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args


	ser_add := arguments[1] //server ip:port
	loc_add := arguments[2] //local ip:port


	ServerAddr,err := net.ResolveUDPAddr("udp",ser_add)
	LocalAddr, err := net.ResolveUDPAddr("udp", loc_add)

	c, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err = c.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024) //receive from server
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("From server: %s\n", string(buffer[0:n]))

	}
}