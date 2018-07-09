// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)

	tcpconn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatalf("able to work only with tcp connection")
	}

	err = tcpconn.CloseWrite()
	if err != nil {
		log.Fatalf("close write tcp connection has failed: %v", err)
	}

	<-done // wait for background goroutine to finish

	// where close read here or in goroutine? -> error: close read tcp connection has failed: close tcp 127.0.0.1:48776->127.0.0.1:8000: shutdown: transport endpoint is not connected
	err = tcpconn.CloseRead()
	if err != nil {
		log.Fatalf("close read tcp connection has failed: %v", err)
	}
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
