package main

import (
	"flag"
	"log"
	"net"
	"net/http"
)

var (
	port string
	dir  string
)

func init() {
	flag.StringVar(&port, "port", "7070", "server running port")
	flag.StringVar(&dir, "dir", ".", "file server path")
	flag.Parse()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	lip := GetOutboundIP()
	log.Printf("Serving %s on HTTP %s:%s\n", dir, lip, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
