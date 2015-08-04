package main

import (
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 64)
	log.Printf("Accept from %s, local %s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Read from %s fail: %s\n", conn.RemoteAddr().String(), err)
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			log.Printf("write to %s fail: %s\n", conn.RemoteAddr().String(), err)
			return
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":9876")
	if err != nil {
		log.Fatal("Listen fail: ", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept fail: ", err)
			continue
		}
		go handleConn(conn)
	}
}
