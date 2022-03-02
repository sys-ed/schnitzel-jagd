package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.ListenPacket("udp", "127.0.0.161:6644")
	if err != nil {
		log.Fatalf("can't create socket: %v", err)
	}

	remoteAddr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.12"),
		Port: 6655,
	}

	for {
		conn.WriteTo([]byte("hello"), &remoteAddr)
		time.Sleep(time.Second * 5)
	}
}
