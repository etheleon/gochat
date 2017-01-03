package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

// RunHost this makes the program run as host
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	fmt.Println("this is the ip and port:" + ipAndPort)
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error cannot listen: " + listenErr.Error())
	}

	conn, connErr := listener.Accept()
	if connErr != nil {
		log.Fatal("Error cannot accept: " + connErr.Error())
	}
	fmt.Println("New connection accepted")
	for {
		handleHost(conn)
	}
}

func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error cannot read: " + readErr.Error())
	}
	fmt.Println("Message received: ", message)

	fmt.Print("Send message: ")
	replyReader := bufio.NewReader(os.Stdin)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error", replyErr.Error())
	}
	fmt.Fprint(conn, replyMessage)
}

// RunGuest
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		fmt.Println("Error", dialErr.Error())
	}
	for {
		handleGuest(conn)
	}
}

func handleGuest(conn net.Conn) {
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}
	fmt.Fprint(conn, message)

	replyReader := bufio.NewReader(conn)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}
	fmt.Println("Message received:", replyMessage)
}
