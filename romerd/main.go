package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	Port = flag.Int("port", 9876, "Server port")
)

func init() {
	flag.Parse()
}

func handleEcho(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)
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

func handleEat(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)
	log.Printf("Accept from %s, local %s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	for {
		_, err := conn.Read(buf)
		if err != nil {
			log.Printf("Read from %s fail: %s\n", conn.RemoteAddr().String(), err)
			return
		}
	}
}

func handlePut(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 8192)
	log.Printf("Accept from %s, local %s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	for {
		_, err := conn.Write(buf)
		if err != nil {
			log.Printf("write to %s fail: %s\n", conn.RemoteAddr().String(), err)
			return
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", *Port))
	if err != nil {
		log.Fatal("Listen fail: ", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept fail: ", err)
			continue
		}
		go handlePut(conn)
	}
}
