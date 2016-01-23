package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Starting server...")
	c, err := net.ListenPacket("udp", ":8032")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	b := make([]byte, 512)
	n, peer, err := c.ReadFrom(b)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n, "bytes read from", peer)
	if _, err := c.WriteTo(b[:n], peer); err != nil {
		log.Fatal(err)
	}

	log.Println()
}
