package main

import (
	"fmt"
	"net"
	"os"
)



func main() {
	addr := os.Getenv("STATSD_SINK_ADDR")
	if addr == "" {
		fmt.Println("Unable to read STATSD_SINK_ADDR from environment")
		os.Exit(1)
	}

	listenAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	ln, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 1024)

	for {
		n, _, err := ln.ReadFromUDP(buf)
		fmt.Print(string(buf[0:n]))

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
