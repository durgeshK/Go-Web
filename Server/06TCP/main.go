package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(conn)
	}

}

func serve(conn net.Conn) {
	ns := bufio.NewScanner(conn)
	for ns.Scan() {
		ln := ns.Text()
		if ln == "" {
			break
		}

		fmt.Println(ns.Text())
	}

	body := "Hello World, I am learning Golang"
	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintln(conn, "HTTP-version SP status-code SP reason-phrase CRLF")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
	//fmt.Println("Code got here.")
	conn.Close()
}
