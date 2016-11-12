package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
	"net"
)

var (
	addr = kingpin.Flag("address", "listener host:port").Required().String()
)

func main() {
	kingpin.Parse()

	listenAddr, err := net.ResolveUDPAddr("udp", *addr)
	if err != nil {
		fmt.Println(err.Error())
	}

	ln, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		fmt.Println(err.Error())
	}

	buf := make([]byte, 1024)

	for {
		n, _, err := ln.ReadFromUDP(buf)
		fmt.Println(string(buf[0:n]))

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
