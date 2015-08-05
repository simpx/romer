package main

import (
	"io"
	"log"
	"net"
)

func main() {
	remoteAddr := "127.0.0.1:9876"
	conn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Fatalf("Dial to %s fail: %s\n", remoteAddr, err)
	}
	defer conn.Close()
	buf := make([]byte, 128)
	_, err = conn.Write(buf)
	if err != nil {
		log.Printf("write to %s fail: %s\n", conn.RemoteAddr().String(), err)
		return
	}
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		log.Printf("read from %s fail: %s\n", conn.RemoteAddr().String(), err)
		return
	}
}
