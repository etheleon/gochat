package main

import (
	"flag"
	"fmt"
	"gochat/lib"
	"os"
)

func main() {
	var isHost bool
	flag.BoolVar(&isHost, "listen", false, "Listen on specific ip address")
	flag.Parse()
	if isHost {
		//go run main.go -listen <ip>
		fmt.Println(os.Args[2])
		connIP := os.Args[2]
		lib.RunHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		lib.RunGuest(connIP)
	}

}
