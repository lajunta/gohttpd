package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	host string
	port string
	dir  string
)

func flagParse() {
	flag.StringVar(&port, "port", "7070", "server running port")
	flag.StringVar(&host, "host", "127.0.0.1", "server host address")
	flag.StringVar(&dir, "dir", ".", "file server path")
	flag.Parse()
}

func init() {
	flagParse()
}
func main() {

	http.Handle("/", http.FileServer(http.Dir(dir)))

	log.Printf("Serving %s on HTTP port: %s\n", dir, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
