package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

var program string

var redirectStderr = flag.Bool("redirectStderr", false, "redirects stderr over the network")
var port = flag.String("address", ":8080", "the address to listen on")

func main() {
	flag.Parse()

	program = flag.Arg(0)

	ln, err := net.Listen("tcp", *port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println(program)
	cmd := exec.Command(program)
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	cmd.Stdin = reader
	cmd.Stdout = writer
	if *redirectStderr {
		cmd.Stderr = writer
	} else {
		cmd.Stderr = os.Stderr
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	conn.Close()
}
