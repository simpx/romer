package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var (
	Host    = flag.String("host", "127.0.0.1", "Server host")
	Port    = flag.Int("port", 9876, "Server port")
	Seconds = flag.Int("sec", 10, "Seconds")
)

func init() {
	flag.Parse()
}

func doForSeconds(fun func(), s time.Duration) float64 {
	var stop bool = false
	f := func() {
		stop = true
	}
	time.AfterFunc(s*time.Second, f)
	start := time.Now()
	for stop == false {
		fun()
	}
	end := time.Now()
	return end.Sub(start).Seconds()
}

func main() {
	remoteAddr := fmt.Sprintf("%s:%d", *Host, *Port)
	conn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Fatalf("Dial to %s fail: %s\n", remoteAddr, err)
	}
	defer conn.Close()
	size := 4096
	buf := make([]byte, size)
	n := 0.0

	f := func() {
		/*
			_, err = conn.Write(buf)
			if err != nil {
				log.Printf("write to %s fail: %s\n", conn.RemoteAddr().String(), err)
				return
			}
		*/
		_, err = io.ReadFull(conn, buf)
		if err != nil {
			log.Printf("read from %s fail: %s\n", conn.RemoteAddr().String(), err)
			return
		}
		n += float64(size)
	}
	seconds := doForSeconds(f, 10)

	log.Printf("%f bytes, %f MiB\n", n, n/1024/1024)
	log.Printf("%f seconds\n", seconds)
	log.Printf("%f MiB/s\n", n/1024/1024/seconds)
	log.Printf("%f Mb\n", n/1000/1000/seconds*8)
}
