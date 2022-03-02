package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listenAddr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.12"),
		Port: 6655,
	}
	conn, err := net.ListenUDP("udp", &listenAddr)
	if err != nil {
		// Linux's loopback allows to bind to addresses other than
		// 127.0.0.1. We are using this here to reduce the likelihood
		// of failing to bind (and who knows, maybe to show off).
		log.Fatalf("Can't bind and listen. Are you running on Linux? Err: %v", err)
	}

	log.Printf("Ready to receive on %v\n", conn.LocalAddr())
	buf := make([]byte, 4096)
	for {
		b, err := conn.Read(buf)
		if err != nil {
			log.Fatalf("Unexpected receive error: %v", err)
		}

		fmt.Printf("Got: %v\n", string(buf[0:b]))
	}
}
