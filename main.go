package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

var (
	localServerHost  string
	remoteServerHost string

	listeningHost  string
	listeningPort  int
	pointingToHost string
	pointingToPort int
	help           bool
)

func main() {
	flag.StringVar(&listeningHost, "lhost", "", "Listening Host IP Address")
	flag.IntVar(&listeningPort, "lport", 0, "Listening port")
	flag.StringVar(&pointingToHost, "thost", "", "Pointing to target Host IP Address")
	flag.IntVar(&pointingToPort, "tport", 0, "Pointing to the target port")
	flag.BoolVar(&help, "help", false, "Print default help")

	flag.Parse()

	if help {
		exitMain()
	}

	// make the flags required
	if len(listeningHost) == 0 || listeningPort == 0 || len(pointingToHost) == 0 || pointingToPort == 0 {
		exitMain()
	}

	if net.ParseIP(listeningHost) == nil || net.ParseIP(pointingToHost) == nil {
		fmt.Printf("IP Address - Invalid\n")
		exitMain()
	}

	localServerHost = listeningHost + ":" + strconv.Itoa(listeningPort)
	remoteServerHost = pointingToHost + ":" + strconv.Itoa(pointingToPort)
	fmt.Println("Listening on: ", localServerHost, " target host: ", remoteServerHost)

	ln, err := net.Listen("tcp", localServerHost)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Reverse proxy server up and listening on ", localServerHost)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func exitMain() {
	flag.PrintDefaults()
	os.Exit(1)
}

func forward(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}

func handleConnection(c net.Conn) {

	remote, err := net.Dial("tcp", remoteServerHost)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection from:", c.RemoteAddr(), "to ->", remoteServerHost)

	// go routines to initiate bi-directional communication for local server with a remote server
	go forward(c, remote)
	go forward(remote, c)
}
