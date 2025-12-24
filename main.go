package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp", ":3000");
	if(err != nil) {
		log.Print(err);
		return;
	}
	defer conn.Close();
	log.Println("server is running on the port : 3000");

	buff := make([] byte, 1024);

	for {
		n, address, err := conn.ReadFrom(buff);
		if (err != nil) {
			continue;
		}
		log.Printf("Received %d bytes from %s\n", n, address.String());
		msg := buff[:n];
		log.Println(string(msg));
		conn.WriteTo([]byte("pong"), address);
	}
}

// [IP HEADER][UDP HEADER][PAYLOAD="hello"]