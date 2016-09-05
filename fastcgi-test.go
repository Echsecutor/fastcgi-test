package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"strconv"
)

func main() {
	var port int
	var folder string
	flag.IntVar(&port, "port", 2020, "port to listen on for fastcgi connections")
	flag.StringVar(&folder, "folder", "C:\\", "folder to browse (just to return something)")
	flag.Parse()

	log.Println("Running fastcgi-test on localhost:", port, " serving a browser in", folder)

	l, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	h := http.FileServer(http.Dir(folder))
	log.Fatal(fcgi.Serve(l, h))
}
