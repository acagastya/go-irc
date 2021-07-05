package main

import (
	"log"
)

func main() {
	conn := GetTLSConn()
	defer conn.Close()

	log.Println("client: connected to: ", conn.RemoteAddr())

	ircClient := GetIRCClient(conn)
	err := ircClient.Run() // Run IRC client
	if err != nil {
		log.Fatal(err)
	}
}
