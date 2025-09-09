package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//resolve address for UDP connection
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:42068")
	if err != nil {
		log.Fatal("error", err)
	}

	//prepare a UDP connection
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal("ERROR", err)
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	//bufio.Reader that reads from os.stdin
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(">")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error", err)
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatal("ERROR", err)
		}
	}
}
