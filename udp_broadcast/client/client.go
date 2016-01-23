package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Starting client...")
	c, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	dst, err := net.ResolveUDPAddr("udp", "255.255.255.255:8032")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := c.WriteTo([]byte("hello"), dst); err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 512)
	n, peer, err := c.ReadFrom(data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Peer: ", peer)
	log.Println("Data: ", string(data[:n]))
}
