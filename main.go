package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"regexp"
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
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", EncodeHandler(fs))
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

func EncodeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		re, _ := regexp.Compile("[.](gz)$")
		if re.MatchString(r.URL.Path) {
			w.Header().Add("Content-Encoding", "gzip")
		}
		next.ServeHTTP(w, r)
	})
}
